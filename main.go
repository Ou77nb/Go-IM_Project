package main

import (
	_ "IM_Project/docs"
	"IM_Project/router"
	"IM_Project/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title IM_Project
// @description 通讯系统
// @version 1.0
// @host localhost:8080
// @BasePath
func main() {

	r := router.Router()
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()

	// 使用gin-swagger生成Swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
