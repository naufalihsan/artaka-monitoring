package controllers

import (
	"github.com/gunturbudikurniawan/Artaka/api/middlewares"
)

func (s *Server) initialRoutes() {
	v1 := s.Router.Group("/api/admin")
	{
		v1.GET("/getall", s.GetMerchants)
		v1.GET("/transactionSales", s.GetLastMerchant)
		v1.GET("/transactionsaved", s.GetLastSaved)
		v1.GET("/transactionOnline", s.GetLastOnline)
		v1.GET("/ShowSleep", s.GetShow)

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

	}
	v4 := s.Router.Group("/api/transaction")
	{
		v4.POST("/savedorder", s.CreateSavedOrder)
		v4.POST("/onlinesales", s.CreateOnlineSales)
		v4.POST("/sales", s.CreateSales)
	}
}
