package service

import (
	"car/model"
	"car/serializer"
)

//创建订单服务
type ServiceCostShowService struct {
	Id           int    `form:"id" json:"id" `
	CarLicense   string `form:"car_license" json:"car_license" `
	ServiceItem  string `form:"service_item" json:"service_item" `
	Material     string `form:"material" json:"material" `
	MaterialCost string `form:"material_cost" json:"material_cost" `
	CreatedAt    string `form:"created_at" json:"created_at" `
}

//订单展示方法
func (service *ServiceCostShowService) ShowServiceCost() serializer.Response {
	costs := []model.ServiceCost{}
	var total int64

	//得到订单数
	if err := model.DB.Model(model.ServiceCost{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	//得到订单集合
	if err := model.DB.Find(&costs).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "数据库连接失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildServiceCosts(costs), uint(total)),
	}
}
