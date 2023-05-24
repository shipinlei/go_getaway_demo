package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shipinlei/go_getaway_demo/dto"
	"github.com/shipinlei/go_getaway_demo/middleware"
)

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}
func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	fmt.Println(params)
}
