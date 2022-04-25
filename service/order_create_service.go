package service

import (
	"car/model"
	"car/serializer"
	"strconv"
	"time"

	"gorm.io/gorm"
)

//创建订单服务
type OrderCreateService struct {
	Id          int       `form:"id" json:"id" `
	UserName    string    `form:"user_name" json:"user_name" `
	UserPhone   string    `form:"user_phone" json:"user_phone" `
	CarLicense  string    `form:"car_license" json:"car_license" `
	CarBrand    string    `form:"car_brand" json:"car_brand" `
	ServiceItem string    `form:"service_item" json:"service_item" `
	Material    string    `form:"material" json:"material" `
	Cost        string    `form:"cost" json:"cost" `
	Employee    string    `form:"employee" json:"employee" `
	Operator    string    `form:"operator" json:"operator" `
	CreatedAt   time.Time `form:"created_at" json:"created_at" `
}

//创建订单方法
func (service *OrderCreateService) CreateOrder() serializer.Response {
	cost, _ := strconv.Atoi(service.Cost)

	//接收form传来的值
	newOrder := model.Order{
		UserName:    service.UserName,
		UserPhone:   service.UserPhone,
		CarLicense:  service.CarLicense,
		CarBrand:    service.CarBrand,
		ServiceItem: service.ServiceItem,
		Cost:        cost,
		Material:    service.Material,
		Employee:    service.Employee,
		Operator:    service.Operator,
		CreatedAt:   time.Now(),
	}

	errAcid := model.DB.Transaction(func(tx *gorm.DB) error {
		//入库
		if err := tx.Create(&newOrder).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		//对应的客户总消费按订单金额增加
		if err := tx.Model(model.Guest{}).Where("car_license = ?", newOrder.CarLicense).UpdateColumn("consumption", gorm.Expr("consumption + ?", newOrder.Cost)).Error; err != nil {
			return err
		}

		//对应的客户服务次数+1
		if err := tx.Model(model.Guest{}).Where("car_license = ?", newOrder.CarLicense).UpdateColumn("service_nums", gorm.Expr("service_nums + ?", 1)).Error; err != nil {
			return err
		}

		//
		p := model.Performance{
			Name:        service.Employee,
			Id:          newOrder.Id,
			ServiceItem: newOrder.ServiceItem,
			Performance: newOrder.Cost / 10,
			CreatedAt:   newOrder.CreatedAt,
		}

		if err := tx.Create(&p).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	if errAcid != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "客户保存失败",
			Error: errAcid.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildOrder(newOrder),
	}
}
