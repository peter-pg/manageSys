package api

import (
	"car/service"

	"github.com/gin-gonic/gin"
)

//新建订单成本接口
func CreateServiceCost(c *gin.Context) {
	service := service.ServiceCostCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateServiceCost()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//展示订单信息接口
func ShowAllServiceCosts(c *gin.Context) {
	service := service.ServiceCostShowService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ShowServiceCost()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//展示订单信息接口
func DeleteServiceCost(c *gin.Context) {
	service := service.ServiceCostDeleteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteServiceCostById(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
