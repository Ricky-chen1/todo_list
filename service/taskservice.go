package service

import (
	"time"
	"westonline/models"
	"westonline/utilities/serializer"

	"gorm.io/gorm"
)

type CommonTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0表示未完成，1表示完成
}

type UpdateTaskService struct {
	Status int `json:"status" form:"status"`
}

type ListAllTaskService struct {
	PageSize int  `json:"page_size" form:"page_size" default:"0"`
	PageNum  int  `json:"page_num" form:"page_num"`
	AllList  bool `json:"all_list" form:"all_list" default:"true"`
	Status   int  `json:"status" form:"status"`
}

// 输入关键字查询事项
type ReadTaskService struct {
	PageSize int    `json:"page_size" form:"page_size"`
	PageNum  int    `json:"page_num" form:"page_num"`
	Keyword  string `json:"keyword" form:"keyword"`
}

type DeleteAllTaskService struct {
	Status    int  `json:"status" form:"status"`
	AllDelete bool `json:"all_delete" form:"all_delete" default:"true"`
}

type DeleteTaskService struct {
}

func (service *CommonTaskService) Create(uid uint) serializer.Response {
	//查找对应用户
	user := models.User{}
	var count int64
	models.DB.First(&user, uid).Count(&count)
	if count == 0 {
		return serializer.Response{
			Status: 400,
			Msg:    "用户不存在",
		}
	}
	task := models.Task{
		User:    user,
		Uid:     uid,
		Title:   service.Title,
		Content: service.Content,
		Status:  0,
		AddTime: int(time.Now().Unix()),
	}
	//创建一条备忘录
	if err := models.DB.Create(&task).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建备忘录成功",
	}
}

// 更新一条待办备忘录
func (service *UpdateTaskService) Update(uid uint, tid string) serializer.Response {
	task := models.Task{}
	if err := models.DB.Where("id=?", tid).First(&task).Error; err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "未查找到事项",
		}
	}
	if task.Uid != uid {
		return serializer.Response{
			Status: 403,
			Msg:    "该用户无权访问",
		}
	}
	if service.Status == task.Status {
		return serializer.Response{
			Status: 400,
			Msg:    "多余的更改操作",
		}
	}
	if err := models.DB.Model(&task).Update("status", service.Status).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库更新失败",
			Err:    err,
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "更改成功",
		Data: serializer.StatusData{
			StatusMsg: serializer.StatusMap[service.Status],
		},
	}
}

// 更新所有事项(需要更改)
func (service *UpdateTaskService) UpdateAll(uid uint) serializer.Response {
	var tasks []models.Task
	models.DB.Model(&models.Task{}).Where("status=? AND uid=?", service.Status, uid).Find(&tasks)
	//状态取反
	res := service.Status * -1
	err := models.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&models.Task{}).
		Preload("User").Where("uid=?", uid).Update("status", res).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "更改失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "更改成功",
		Data: serializer.StatusData{
			StatusMsg: serializer.StatusMap[res],
		},
	}
}

// 查询所有事项
func (service *ListAllTaskService) List(uid uint) serializer.Response {
	var tasks []models.Task
	var total int64
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	if !service.AllList {
		models.DB.Model(&models.Task{}).Preload("User").Where("uid=?", uid).Where("status=?", service.Status).
			Count(&total).Limit(service.PageSize).
			Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
		return serializer.BuildListResponse(tasks, uint(total))
	}
	models.DB.Model(&models.Task{}).Preload("User").Where("uid=?", uid).Count(&total).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(tasks, uint(total))
}

// 根据关键字查询备忘录(带分页)
func (service *ReadTaskService) Read(uid uint) serializer.Response {
	var total int64
	var tasks []models.Task
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	models.DB.Model(&models.Task{}).Preload("User").Where("uid=?", uid).
		Where("content LIKE ?", "%"+service.Keyword+"%").Count(&total).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(tasks, uint(total))
}

// 删除一条待办/已完成事项
func (service *DeleteTaskService) Delete(uid uint, tid string) serializer.Response {
	task := models.Task{}
	models.DB.Model(&models.Task{}).First(&task, tid)
	if uid != task.Uid {
		return serializer.Response{
			Status: 400,
			Msg:    "该用户无权限访问他人备忘录",
		}
	}
	if err := models.DB.Delete(&task).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

// 删除所有备忘录
func (service *DeleteAllTaskService) DeleteAll(uid uint) serializer.Response {
	if !service.AllDelete {
		err := models.DB.Where("status=? And uid=?", service.Status, uid).Delete(&models.Task{}).Error
		if err != nil {
			return serializer.Response{
				Status: 500,
				Msg:    "相应事务删除失败",
			}
		}
		return serializer.Response{
			Status: 200,
			Msg:    "已删除所有待办或已完成事务",
		}
	}
	if err := models.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Preload("User").
		Where("uid=?", uid).Delete(&models.Task{}).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "备忘录删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "成功删除所有备忘录",
	}
}
