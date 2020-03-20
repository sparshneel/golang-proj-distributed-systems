package http_server

import "C"
import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang-proj-distributed-systems/domain"
	service "golang-proj-distributed-systems/sevice"
)

var(
	router = gin.Default()
)

func AddBusiness(c *gin.Context) {
	var b *domain.Business
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil{

	}
	service.AddBusiness(b)
	c.JSON(201, gin.H{
			"message": "Business added successfully",
		})
}

func GetBusiness(c *gin.Context) {
	//service.GetBusinessData()
	c.JSON(200, gin.H{
		"message": "Implement me ",
	})
}

func GetBusinessById(c *gin.Context){
	filterIds :=make(map[string]string)
	filterIds["id"] = c.Param("id")
	filterIds["city"] = c.Query("city")
	filterIds["state"] = c.Query("state")
	logrus.Info("parameters passed for the get request, id: " + c.Param("id") + " " + c.Query("city") + " " + c.Query("state"))
	business := service.GetBusinessDataById(filterIds)
	logrus.Info(business)
	c.JSON(200, business)

}

func UpdateBusiness(c *gin.Context){
	var b *domain.Business
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil{

	}
	service.UpdateBusiness(b)
	c.JSON(200, gin.H{
		"message": "Implement me ",
	})
}

func DeleteBusiness(c *gin.Context){
	service.DeleteBusiness(c.Params.ByName("id"))
	c.JSON(200, gin.H{
		"message": "Business with id: " + c.Params.ByName("id") + " deleted successfully",
	})
}
