package service

import (
	"car/model"
	"car/serializer"
)

//展示员工信息的服务
type EmployeeShowService struct {
	Id         int    `form:"id" json:"id"`
	Name       string `form:"name" json:"name" `
	Gender     string `form:"gender" json:"gender" `
	IdCard     string `form:"idcard" json:"id_card" `
	Birthday   string `form:"birthday" json:"birthday" `
	Phone      string `form:"phone" json:"phone" `
	Address    string `form:"address" json:"address" `
	IsHead     string `form:"is_head" json:"is_head" `
	Department string `form:"department" json:"department"`
	Salary     int    `form:"salary" json:"salary"`
}

//展示员工信息的方法
func (service *EmployeeShowService) ShowEmployees() serializer.Response {
	employees := []model.Employee{}
	var total int64

	//得到员工数
	if err := model.DB.Model(model.Employee{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	//得到员工集合
	if err := model.DB.Find(&employees).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "数据库连接失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildEmployees(employees), uint(total)),
	}
}
