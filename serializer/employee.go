package serializer

import (
	"car/model"
	"time"
)

// Employee 员工信息序列化器
type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	IdCard     string `json:"id_card"`
	Birthday   string `json:"birthday"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	IsHead     string `json:"is_head"`
	Department string `json:"department"`
	Salary     int    `json:"salary" `
}

//Performance 员工绩效序列化器
type Performance struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	ServiceItem string    `json:"service_item" `
	Performance int       `json:"performance" `
	CreatedAt   time.Time `json:"created_at" `
}

// BuildEmployee 员工信息序列化
func BuildEmployee(emp model.Employee) Employee {
	return Employee{
		Id:         emp.Id,
		Name:       emp.Name,
		Gender:     emp.Gender,
		IdCard:     emp.IdCard,
		Birthday:   emp.Birthday,
		Phone:      emp.Phone,
		Address:    emp.Address,
		IsHead:     emp.IsHead,
		Department: emp.Department,
		Salary:     emp.Salary,
	}
}

// BuildEmployees 员工信息集合序列化
func BuildEmployees(items []model.Employee) []Employee {
	employees := []Employee{}
	for _, item := range items {
		emp := BuildEmployee(item)
		employees = append(employees, emp)
	}

	return employees
}

// BuildSalary 绩效序列化
func BuildPerformance(p model.Performance) Performance {
	return Performance{
		Id:          p.Id,
		Name:        p.Name,
		ServiceItem: p.ServiceItem,
		Performance: p.Performance,
		CreatedAt:   p.CreatedAt,
	}
}

// BuildEmployees 绩效集合序列化
func BuildEmployeePerformance(items []model.Performance) []Performance {
	performances := []Performance{}
	for _, item := range items {
		p := BuildPerformance(item)
		performances = append(performances, p)
	}

	return performances
}
