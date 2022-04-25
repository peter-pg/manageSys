package service

import (
	"car/model"
	"car/serializer"
)

// UpdateGuestService 更新客户信息的服务
//客户某些必要信息决定后，不应该轻易个改变，因此只能改变电话和是否是会员
type GuestUpdateService struct {
	UserPhone string `form:"user_phone" json:"user_phone" `
	IsVip     string `form:"is_vip" json:"is_vip"`
}

// Update 更新客户信息
func (service *GuestUpdateService) UpdateGuestById(id string) serializer.Response {
	var guest model.Guest

	//查询实例是否存在
	err := model.DB.First(&guest, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "客户不存在",
			Error: err.Error(),
		}
	}

	//更新电话和是否是会员
	guest.UserPhone = service.UserPhone
	guest.IsVip = service.IsVip

	//保存修改到数据库
	err = model.DB.Save(&guest).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "客户保存失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildGuest(guest),
	}
}
