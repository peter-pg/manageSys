package api

import (
	"car/service"

	"github.com/gin-gonic/gin"
)

//新建客户接口
func CreateNewGuest(c *gin.Context) {
	service := service.GuestCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateGuest()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//展示客户信息接口
func ShowAllGuests(c *gin.Context) {
	service := service.GuestShowService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ShowGuests()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//删除客户信息接口
func DeleteGuest(c *gin.Context) {
	service := service.GuestDeleteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteGuestById(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//更新客户信息接口
func UpdateGuest(c *gin.Context) {
	service := service.GuestUpdateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateGuestById(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
