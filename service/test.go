package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsyodyssey/agoge/dao"
	"github.com/hsyodyssey/agoge/model"
)

func GetAllTests(c *gin.Context) {
	var testtables []dao.Testtable
	err := model.GetAllTests(&testtables)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, testtables)
	}
}

func GetFirstTest(c *gin.Context) {
	var testtable dao.Testtable
	err := model.GetFirstTest(&testtable)
	fmt.Println(testtable)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, testtable)
	}
}
