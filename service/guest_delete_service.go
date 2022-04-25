package service

import (
	"car/model"
	"car/serializer"
)

//删除客户信息的服务
type GuestDeleteService struct {
}

//删除客户信息的方法
func (service *GuestDeleteService) DeleteGuestById(id string) serializer.Response {
	var guest model.Guest

	//查询实例是否存在
	if err := model.DB.First(&guest, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "客户不存在",
			Error: err.Error(),
		}
	}

	//销库
	if err := model.DB.Delete(&guest).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "客户删除失败",
			Error: err.Error(),
		}
	}

	//delete success, return ""
	return serializer.Response{}
}
