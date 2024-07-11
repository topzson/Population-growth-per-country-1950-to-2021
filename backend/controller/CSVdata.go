package controller

import (
	"github.com/topzson/Population-growth-per-country-1950-to-2021/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// 7.ดึงข้อมูลด้วย ID GET /user/id

func Getdatabyid(c *gin.Context) {

	var data entity.CSVdata

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM cs_vdata WHERE id = ? ", id).Scan(&data).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}

// GET /users

func Listdata(c *gin.Context) {

	var data []entity.CSVdata

	if err := entity.DB().Raw("SELECT * FROM cs_vdata").Scan(&data).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}
