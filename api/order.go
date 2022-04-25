package api

import (
	"car/service"

	"github.com/gin-gonic/gin"
)

//新建订单接口
func CreateNewOrder(c *gin.Context) {
	service := service.OrderCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateOrder()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//展示订单信息接口
func ShowAllOrders(c *gin.Context) {
	service := service.OrderShowService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ShowOrders()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//删除订单信息接口
func DeleteOrder(c *gin.Context) {
	service := service.OrderDeleteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteOrderById(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//更新订单信息接口
func UpdateOrder(c *gin.Context) {
	service := service.OrderUpdateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateOrderById(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
