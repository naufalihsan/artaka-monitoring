package controllers

import (
	"github.com/gunturbudikurniawan/Artaka/api/middlewares"
)

func (s *Server) initialRoutes() {
	v1 := s.Router.Group("/api/admin")
	{
		v1.GET("/transactionsaved", s.GetLastSaved)
		v1.GET("/transactionOnline", s.GetLastOnline)
		v1.GET("/NotYetContact", s.NotAll)
		v1.GET("/Already", s.Already)
		v1.GET("/ShowSleep", s.Showall)
		v1.GET("/NotRespon", s.LateRespon)
		v1.GET("/ShowSalesPayment", s.ShowSalesPayment)
		v1.GET("/ShowOnlineSalesPayment", s.ShowOnlineSalesPayment)
		v1.GET("/GetAllSubcribers", s.GetCertainSubscribers)

		v1.GET("/getall/:id", s.GetMerchant)
		v1.POST("/register", s.CreateAdmin)
		v1.POST("/login", s.LoginAdmin)
		v1.PUT("/update/:id", middlewares.TokenAuthMiddleware(), s.UpdateAdmin)

	}
	v2 := s.Router.Group("/api/merchant")
	{
		v2.POST("/register", s.CreateMerchants)
		v2.POST("/login", s.LoginMerchant)
		v2.PUT("/update/:id", middlewares.TokenAuthMiddleware(), s.UpdateMerchant)
	}
	v3 := s.Router.Group("/api/post")
	{
		v3.POST("/create", middlewares.TokenAuthMiddleware(), s.CreatePost)
		v3.PUT("/:id", middlewares.TokenAuthMiddleware(), s.UpdatePost)
		v3.GET("/getpost/:id", s.GetPost)

	}
	v4 := s.Router.Group("/api/transaction")
	{
		v4.POST("/savedorder", s.CreateSavedOrder)
		v4.POST("/onlinesales", s.CreateOnlineSales)
		v4.POST("/sales", s.CreateSales)
	}
}
