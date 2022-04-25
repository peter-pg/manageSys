package api

import (
	"car/service"

	"github.com/gin-gonic/gin"
)

//新建员工接口
func CreateEmployee(c *gin.Context) {
	service := service.EmployeeCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateEmployee()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//展示员工信息接口
func ShowAllEmployeeInfo(c *gin.Context) {
	service := service.EmployeeShowService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ShowEmployees()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//删除员工信息接口
func DeleteEmployee(c *gin.Context) {
	service := service.EmployeeDeleteService{}
	res := service.DeleteEmplyeeById(c.Param("id"))
	c.JSON(200, res)
}

//展示员工绩效接口
func ShowEmployeePerformance(c *gin.Context) {
	service := service.EmployeeShowPerformanceService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ShowPerformance()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

//更新员工信息接口
func UpdateEmployee(c *gin.Context) {
	service := service.EmployeeUpdateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateEmployeeById(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
