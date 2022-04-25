package service

import (
	"car/model"
	"car/serializer"
)

//新建客户的服务
type GuestCreateService struct {
	Id          int    `form:"id" json:"id" `
	UserName    string `form:"user_name" json:"user_name" `
	UserPhone   string `form:"user_phone" json:"user_phone" `
	CarLicense  string `form:"car_license" json:"car_license" `
	CarBrand    string `form:"car_brand" json:"car_brand" `
	ServiceNums int    `form:"service_nums" json:"service_nums" `
	Consumption int    `form:"consumption" json:"consumption" `
	IsVip       string `form:"is_vip" json:"is_vip" `
}

//新建客户的方法
func (service *GuestCreateService) CreateGuest() serializer.Response {
	newGuest := model.Guest{
		UserName:   service.UserName,
		UserPhone:  service.UserPhone,
		CarLicense: service.CarLicense,
		CarBrand:   service.CarBrand,
		IsVip:      service.IsVip,
	}

	//入库
	err := model.DB.Create(&newGuest).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "员工信息保存失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildGuest(newGuest),
	}
}
