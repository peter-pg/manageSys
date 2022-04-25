package service

import (
	"car/model"
	"car/serializer"
)

// DeleteEmployeeById 删除员工信息的服务
type EmployeeDeleteService struct {
}

//删除员工信息方法（通过员工id）
func (d *EmployeeDeleteService) DeleteEmplyeeById(id string) serializer.Response {
	var emp model.Employee

	//查询实例是否存在
	if err := model.DB.First(&emp, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "员工不存在",
			Error: err.Error(),
		}
	}

	//销库
	if err := model.DB.Delete(&emp).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "员工删除失败",
			Error: err.Error(),
		}
	}

	//delete success, return ""
	return serializer.Response{}

}
