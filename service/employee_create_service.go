package service

import (
	"car/model"
	"car/serializer"
	"strconv"
)

//新建员工信息的服务
type EmployeeCreateService struct {
	Name       string `form:"name" json:"name" `
	Gender     string `form:"gender" json:"gender" `
	IdCard     string `form:"id_card" json:"id_card" `
	Birthday   string `form:"birthday" json:"birthday" `
	Phone      string `form:"phone" json:"phone" `
	Address    string `form:"address" json:"address" `
	IsHead     string `form:"is_head" json:"is_head" `
	Department string `form:"department" json:"department"`
	Salary     string `form:"salary" json:"salary" `
}

//新建员工信息的方法
func (service *EmployeeCreateService) CreateEmployee() serializer.Response {
	salary, _ := strconv.Atoi(service.Salary)

	//get 员工信息 from form
	newEmployee := model.Employee{
		Name:       service.Name,
		Gender:     service.Gender,
		IdCard:     service.IdCard,
		Birthday:   service.Birthday,
		Phone:      service.Phone,
		Address:    service.Address,
		IsHead:     service.IsHead,
		Department: service.Department,
		Salary:     salary,
	}

	//入库
	err := model.DB.Create(&newEmployee).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "员工信息保存失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildEmployee(newEmployee),
	}
}
