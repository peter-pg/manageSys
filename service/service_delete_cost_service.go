package service

import (
	"car/model"
	"car/serializer"
)

// DeleteServiceCostById 删除成本信息的服务
type ServiceCostDeleteService struct {
}

//删除成本信息方法（通过成本信息id）
func (d *ServiceCostDeleteService) DeleteServiceCostById(id string) serializer.Response {
	var cost model.ServiceCost

	//查询实例是否存在
	if err := model.DB.First(&cost, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "成本信息不存在",
			Error: err.Error(),
		}
	}

	//销库
	if err := model.DB.Delete(&cost).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "成本信息删除失败",
			Error: err.Error(),
		}
	}

	//delete success, return ""
	return serializer.Response{}

}
