package service

import (
	"IM_Project/utils"
	"github.com/gin-gonic/gin"
)

// GetPing
// @Tags swaggerTest
// @Success 200 {string} ping
// @Router /ping [get]
func GetPing(c *gin.Context) {
	utils.Success(c, 1, "success", "ping")
}
