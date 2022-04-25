package service

import (
	"car/model"
	"car/serializer"
	"strconv"
)

//新建员工信息的服务
type StoreCreateService struct {
	MaterialName string `form:"material_name" json:"material_name" `
	Manufacturer string `form:"manufacturer" json:"manufacturer" `
	UnitPrice    string `form:"unit_price" json:"unit_price" `
	Number       string `form:"number" json:"number" `
}

//新建员工信息的方法
func (service *StoreCreateService) CreateStore() serializer.Response {
	unitPrice, _ := strconv.Atoi(service.UnitPrice)
	number, _ := strconv.Atoi(service.Number)
	//get 员工信息 from form
	newStore := model.Store{
		MaterialName: service.MaterialName,
		Manufacturer: service.Manufacturer,
		UnitPrice:    unitPrice,
		Number:       number,
		TotalPrice:   unitPrice * number,
	}

	//入库
	err := model.DB.Create(&newStore).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "员工信息保存失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildStore(newStore),
	}
}
