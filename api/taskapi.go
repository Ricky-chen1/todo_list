package api

import (
	"westonline/pkg/utils"
	"westonline/service"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	//获取传来的数据
	var taskCreate *service.CommonTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"), "golang")
	if err := c.ShouldBind(&taskCreate); err != nil {
		c.JSON(400, err)
	} else {
		res := taskCreate.Create(claim.Id)
		c.JSON(200, res)
	}
}

func UpdateTask(c *gin.Context) {
	var taskUpdate *service.UpdateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"), "golang")
	if err := c.ShouldBind(&taskUpdate); err != nil {
		c.JSON(400, err)
	} else {
		res := taskUpdate.Update(claim.Id, c.Param("id"))
		c.JSON(200, res)
	}
}

func UpdateAllTask(c *gin.Context) {
	var AlltaskUpdate *service.UpdateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"), "golang")
	//传入需要更改待办事项还是已完成事项
	if err := c.ShouldBind(&AlltaskUpdate); err != nil {
		c.JSON(400, err)
	} else {
		res := AlltaskUpdate.UpdateAll(claim.Id)
		c.JSON(200, res)
	}
}

func ReadTask(c *gin.Context) {
	var taskRead *service.ReadTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"), "golang")
	if err := c.ShouldBind(&taskRead); err != nil {
		c.JSON(400, err)
	} else {
		res := taskRead.Read(claim.Id)
		c.JSON(200, res)
	}
}

func ListAllTask(c *gin.Context) {
	var taskList *service.ListAllTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"), "golang")
	if err := c.ShouldBind(&taskList); err != nil {
		c.JSON(400, err)
	} else {
		res := taskList.List(claim.Id)
		c.JSON(200, res)
	}
}

func DeleteTask(c *gin.Context) {
	var taskDelete *service.DeleteTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"), "golang")
	if err := c.ShouldBind(&taskDelete); err != nil {
		c.JSON(400, err)
	} else {
		res := taskDelete.Delete(claim.Id, c.Param("id"))
		c.JSON(200, res)
	}
}

func DeleteAllTask(c *gin.Context) {
	var AlltaskDelete *service.DeleteAllTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"), "golang")
	if err := c.ShouldBind(&AlltaskDelete); err != nil {
		c.JSON(400, err)
	} else {
		res := AlltaskDelete.DeleteAll(claim.Id)
		c.JSON(200, res)
	}
}
