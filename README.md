# todo_list 备忘录

**此项目使用Gin+Gorm ，基于RESTful API实现的一个备忘录**

# 项目主要功能介绍  
* 1.用户注册登录 ( jwt-go鉴权 )  
* 2.新增/删除/修改/查询 备忘录  
* 3.分页功能  

# 项目结构  
```
westonline/  
├── api  
├── middleware  
├── model  
├── utiilities  
│  ├── tokenfunc  
│  └── serializer  
├── routes  
└── service  
```     
  
* api : 用于定义接口函数    
* middleware : 应用中间件    
* model : 应用数据库模型  
* utilities/tokenfunc : token对应工具  
* routes : 路由逻辑处理  
* utilities/serializer : 将数据序列化为 json 的函数      
* service : 接口函数的实现(业务逻辑)       

# 项目运行  
**此项目使用Go Mod管理依赖。**    
## 下载依赖    
`go mod tidy`  
## 运行  
`fresh`
