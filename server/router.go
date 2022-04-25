package server

import (
	"car/api"
	"car/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		//employee api
		//create new employee
		v1.POST("employee", api.CreateEmployee)
		//show all employees
		v1.GET("employees", api.ShowAllEmployeeInfo)
		//show employees performances
		v1.GET("employees/performance", api.ShowEmployeePerformance)
		//delete employee
		v1.DELETE("employee/:id", api.DeleteEmployee)
		//update employee
		v1.PUT("employee/:id", api.UpdateEmployee)

		//guest api
		//create new guest
		v1.POST("guest", api.CreateNewGuest)
		//show all guests information
		v1.GET("guests", api.ShowAllGuests)
		//delete guest information
		v1.DELETE("guest/:id", api.DeleteGuest)
		//update guest
		v1.PUT("guest/:id", api.UpdateGuest)

		//order api
		//create new order
		v1.POST("order", api.CreateNewOrder)
		//show all guests information
		v1.GET("orders", api.ShowAllOrders)
		//delete guest information
		v1.DELETE("order/:id", api.DeleteOrder)
		//update order
		v1.PUT("order/:id", api.UpdateOrder)

		//service_cost api
		//create a service cost
		v1.POST("service/cost", api.CreateServiceCost)
		//show all service_costs information
		v1.GET("service/costs", api.ShowAllServiceCosts)
		//delete service cost information
		v1.DELETE("service/cost/:id", api.DeleteServiceCost)

		//store api
		//create a store
		v1.POST("store", api.CreateStore)
		//show stores
		v1.GET("stores", api.ShowAllStores)

		v1.POST("ping", api.Ping)

		//user api
		//user register
		v1.POST("user/register", api.UserRegister)
		//user login
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
