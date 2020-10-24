package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hsyodyssey/agoge/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/test-api")
	{
		grp1.GET("user", service.GetFirstTest)
		grp1.GET("users", service.GetAllTests)
	}
	grp2 := r.Group("/test")
	{
		grp2.GET("user", service.GetFirstTest)
		grp2.GET("users", service.GetAllTests)
	}
	return r
}
