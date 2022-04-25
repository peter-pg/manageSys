package service

import (
	"car/model"
	"car/serializer"
	"strconv"
)

//创建订单服务
type ServiceCostCreateService struct {
	Id           int    `form:"id" json:"id" `
	CarLicense   string `form:"car_license" json:"car_license" `
	ServiceItem  string `form:"service_item" json:"service_item" `
	Material     string `form:"material" json:"material" `
	MaterialCost string `form:"material_cost" json:"material_cost" `
}

//创建订单方法
func (service *ServiceCostCreateService) CreateServiceCost() serializer.Response {
	cost, _ := strconv.Atoi(service.MaterialCost)

	//接收form传来的值
	newServiceCost := model.ServiceCost{
		Id:           service.Id,
		CarLicense:   service.CarLicense,
		ServiceItem:  service.ServiceItem,
		MaterialCost: cost,
		Material:     service.Material,
	}

	//入库
	err := model.DB.Create(&newServiceCost).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "支出信息保存失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildServiceCost(newServiceCost),
	}
}
