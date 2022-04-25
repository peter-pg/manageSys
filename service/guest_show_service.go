package service

import (
	"car/model"
	"car/serializer"
)

//展示客户信息的服务
type GuestShowService struct {
	Id          int    `form:"id" json:"id" `
	UserName    string `form:"user_name" json:"user_name" `
	UserPhone   string `form:"user_phone" json:"user_phone" `
	CarLicense  string `form:"car_license" json:"car_license" `
	CarBrand    string `form:"car_brand" json:"car_brand" `
	ServiceNums int    `form:"service_nums" json:"service_nums" `
	Consumption int    `form:"consumption" json:"consumption" `
	IsVip       string `form:"is_vip" json:"is_vip" `
}

//展示客户信息的方法
func (service *GuestShowService) ShowGuests() serializer.Response {
	guests := []model.Guest{}
	var total int64

	//得到客户数
	if err := model.DB.Model(model.Guest{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	//得到客户集合
	if err := model.DB.Find(&guests).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "数据库连接失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildGuests(guests), uint(total)),
	}
}
