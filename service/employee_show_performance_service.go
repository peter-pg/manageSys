package service

import (
	"car/model"
	"car/serializer"
)

//展示员工绩效的服务
type EmployeeShowPerformanceService struct {
	Id          int    `form:"id" json:"id"`
	Name        string `form:"name" json:"name" `
	ServiceItem string `form:"service_item" json:"service_item" `
	Performance int    `form:"performance" json:"performance" `
	CreatedAt   string `form:"created_at" json:"created_at" `
}

//展示员工绩效的方法
func (service *EmployeeShowPerformanceService) ShowPerformance() serializer.Response {
	performances := []model.Performance{}
	var total int64

	//得到员工绩效数
	if err := model.DB.Model(model.Performance{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	//得到员工绩效集合
	if err := model.DB.Find(&performances).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "数据库连接失败",
			Error: err.Error(),
		}
	}

	//返回json数据
	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildEmployeePerformance(performances), uint(total)),
	}
}
