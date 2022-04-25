package service

import (
	"car/model"
	"car/serializer"
)

// UpdateEmployeeService 更新员工信息的服务
type EmployeeUpdateService struct {
	Phone      string `form:"phone" json:"phone" `
	Address    string `form:"address" json:"address" `
	IsHead     string `form:"is_head" json:"is_head" `
	Department string `form:"department" json:"department"`
	Salary     int    `form:"salary" json:"salary" `
}

// Update 更新员工信息的方法
func (service *EmployeeUpdateService) UpdateEmployeeById(id string) serializer.Response {
	var employee model.Employee

	//查询实例是否存在
	err := model.DB.First(&employee, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "员工不存在",
			Error: err.Error(),
		}
	}

	//更新员工信息
	employee.Phone = service.Phone
	employee.Address = service.Address
	employee.IsHead = service.IsHead
	employee.Department = service.Department
	employee.Salary = service.Salary

	//入库
	err = model.DB.Save(&employee).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "员工保存失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildEmployee(employee),
	}
}
