package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mobilejavierg/apirestgo/libs"
)

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/despachos/years/", GetYears)
		v1.GET("/despachos/year/:id", GetYear)
		v1.GET("/despachos/yearmonth/:yearmonth", GetYearMonth)
	}

	r.Run(":20678")
}

func GetYears(c *gin.Context) {

	resp := consultaDespachos.GetYears()

	c.JSON(200, resp)

	// curl -i http://localhost:8080/api/v1/users
}

func GetYear(c *gin.Context) {

	id := c.Params.ByName("id")

	yearId, _ := strconv.Atoi(id)

	resp := consultaDespachos.GetYear(yearId)

	c.JSON(200, resp)

	// curl -i http://localhost:8080/api/v1/users/1
}

func GetYearMonth(c *gin.Context) {

	id := c.Params.ByName("yearmonth")

	resp := consultaDespachos.GetYearMonth(id)

	c.JSON(200, resp)

	// curl -i http://localhost:8080/api/v1/users/1
}
