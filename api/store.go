package api

import (
	"car/service"

	"github.com/gin-gonic/gin"
)

//新建订单接口
func CreateStore(c *gin.Context) {
	service := service.StoreCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateStore()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//展示订单信息接口
func ShowAllStores(c *gin.Context) {
	service := service.StoreShowService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ShowStore()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
