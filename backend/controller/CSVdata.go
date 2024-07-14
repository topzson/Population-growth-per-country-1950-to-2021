package controller

import (
	"github.com/topzson/Population-growth-per-country-1950-to-2021/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Getdatabyid(c *gin.Context) {

	var data entity.CSVdata

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM cs_vdata WHERE id = ? ", id).Scan(&data).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}

func Listdata(c *gin.Context) {

	var data []entity.CSVdata

	if err := entity.DB().Raw("SELECT a.name ,a.date,a.category,a.value,b.region,b.url FROM cs_vdata a inner join regions b on a.name = b.name").Scan(&data).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}
