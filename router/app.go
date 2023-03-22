package router

import (
	"IM_Project/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()

	//静态资源
	r.Static("/asset", "asset/")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//r.StaticFS()
	r.LoadHTMLGlob("views/**/*")

	r.GET("/ping", service.GetPing)

	//前端
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	r.GET("/chat", service.Chat)

	//上传文件
	r.POST("/attach/upload", service.Upload)
	//用户模块
	r.POST("/user/createUser", service.CreateUser)
	r.POST("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	r.GET("/user/getUserList", service.GetUserList)
	//发送消息
	r.GET("/user/sendMsg", service.SendMsg)
	//发送消息
	r.GET("/user/sendUserMsg", service.SendUserMsg)
	//添加好友
	r.POST("/contact/addfriend", service.AddFriend)
	//好友列表
	r.POST("/searchFriends", service.SearchFriends)
	//前端发过来的消息再receProc里面处理
	//r.POST("/user/redisMsg", service.RedisMsg)
	return r
}
