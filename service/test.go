package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsyodyssey/agoge/dao"
	"github.com/hsyodyssey/agoge/model"
)

func GetAllTests(c *gin.Context) {
	var testtable []dao.Testtable
	err := model.GetAllTests(&testtable)

	// fmt.Println(testtable[1])

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, testtable)
	}
}

func GetFirstTest(c *gin.Context) {
	var testtable dao.Testtable
	err := model.GetFirstTest(&testtable)

	// fmt.Println(testtable[1])

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, testtable)
	}
}
