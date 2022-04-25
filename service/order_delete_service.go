package service

import (
	"car/model"
	"car/serializer"

	"gorm.io/gorm"
)

//订单删除服务
type OrderDeleteService struct {
}

//订单删除方法（通过id定位）
func (service *OrderDeleteService) DeleteOrderById(id string) serializer.Response {
	var orders model.Order
	var performance model.Performance
	var serviceCost model.ServiceCost

	errAcid := model.DB.Transaction(func(tx *gorm.DB) error {

		//查询实例是否存在
		if err := tx.First(&orders, id).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.First(&performance, id).Error; err != nil {
			return err
		}

		if err := tx.First(&serviceCost, id).Error; err != nil {
			return err
		}

		//对应的客户总消费按订单金额减少
		if err := tx.Model(model.Guest{}).Where("car_license = ?", orders.CarLicense).UpdateColumn("consumption", gorm.Expr("consumption - ?", orders.Cost)).Error; err != nil {
			return err
		}

		//对应的客户服务次数-1
		if err := tx.Model(model.Guest{}).Where("car_license = ?", orders.CarLicense).UpdateColumn("service_nums", gorm.Expr("service_nums - ?", 1)).Error; err != nil {
			return err
		}

		//销库
		if err := tx.Delete(&orders).Error; err != nil {
			return err
		}

		if err := tx.Delete(&performance).Error; err != nil {
			return err
		}

		if err := tx.Delete(&serviceCost).Error; err != nil {
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

	//delete success, return ""
	return serializer.Response{}
}
