package service

import (
	"car/model"
	"car/serializer"

	"gorm.io/gorm"
)

// UpdateOrderService 更新订单的服务
type OrderUpdateService struct {
	ServiceItem string `form:"service_item" json:"service_item" `
	Material    string `form:"material" json:"material" `
}

// Update 更新订单
func (service *OrderUpdateService) UpdateOrderById(id string) serializer.Response {
	var order model.Order
	var performance model.Performance
	err := model.DB.First(&order, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "员工不存在",
			Error: err.Error(),
		}
	}

	errAcid := model.DB.Transaction(func(tx *gorm.DB) error {

		//查询实例是否存在
		if err := tx.First(&order, id).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.First(&performance, id).Error; err != nil {
			return err
		}

		//修改
		order.ServiceItem = service.ServiceItem
		order.Material = service.Material
		performance.ServiceItem = service.ServiceItem

		//保存修改到数据库
		if err := tx.Save(&performance).Error; err != nil {
			return err
		}

		if err := tx.Save(&order).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})

	//事务发送错误，删除失败
	if errAcid != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "订单删除失败",
			Error: errAcid.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildOrder(order),
	}
}
