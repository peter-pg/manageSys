package service

import (
	"car/model"
	"car/serializer"
)

//订单展示服务
type OrderShowService struct {
	Id          int    `form:"id" json:"id" `
	UserName    string `form:"user_name" json:"user_name" `
	UserPhone   string `form:"user_phone" json:"user_phone" `
	CarLicense  string `form:"car_license" json:"car_license" `
	CarBrand    string `form:"car_brand" json:"car_brand" `
	ServiceItem string `form:"service_item" json:"service_item" `
	Material    string `form:"material" json:"material" `
	Cost        int    `form:"cost" json:"cost" `
	Employee    string `form:"employee" json:"employee" `
	Operator    string `form:"operator" json:"operator" `
	CreatedAt   string `form:"created_at" json:"created_at" `
}

//订单展示方法
func (service *OrderShowService) ShowOrders() serializer.Response {
	orders := []model.Order{}
	var total int64

	//得到订单数
	if err := model.DB.Model(model.Order{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	//得到订单集合
	if err := model.DB.Find(&orders).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "数据库连接失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildOrders(orders), uint(total)),
	}
}
