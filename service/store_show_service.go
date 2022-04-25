package service

import (
	"car/model"
	"car/serializer"
)

//创建订单服务
type StoreShowService struct {
	Id           int    `form:"id" json:"id" `
	MaterialName string `form:"material_name" json:"material_name" `
	UnitPrice    string `form:"unit_price" json:"unit_price" `
	Number       string `form:"number" json:"number" `
	TotalPrice   string `form:"total_price" json:"total_price" `
	CreatedAt    string `form:"created_at" json:"created_at" `
}

//订单展示方法
func (service *StoreShowService) ShowStore() serializer.Response {
	stores := []model.Store{}
	var total int64

	//得到订单数
	if err := model.DB.Model(model.Store{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	//得到订单集合
	if err := model.DB.Find(&stores).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "数据库连接失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildStores(stores), uint(total)),
	}
}
