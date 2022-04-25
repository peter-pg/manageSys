package serializer

import "car/model"

//客人序列化器
type Guest struct {
	Id          int    `json:"id" `
	UserName    string `json:"user_name" `
	UserPhone   string `json:"user_phone" `
	CarLicense  string `json:"car_license" `
	CarBrand    string `json:"car_brand"`
	ServiceNums int    `json:"service_nums"`
	Consumption int    `json:"consumption" `
	IsVip       string `json:"is_vip"`
}

// BuildGuest 客人信息序列化
func BuildGuest(g model.Guest) Guest {
	return Guest{
		Id:          g.Id,
		UserName:    g.UserName,
		UserPhone:   g.UserPhone,
		CarLicense:  g.CarLicense,
		CarBrand:    g.CarBrand,
		ServiceNums: g.ServiceNums,
		Consumption: g.Consumption,
		IsVip:       g.IsVip,
	}
}

// BuildGuests 客人信息集合序列化
func BuildGuests(items []model.Guest) []Guest {
	guests := []Guest{}
	for _, item := range items {
		g := BuildGuest(item)
		guests = append(guests, g)
	}

	return guests
}
