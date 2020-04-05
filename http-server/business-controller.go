package http_server

import "C"
import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang-proj-distributed-systems/domain"
	"golang-proj-distributed-systems/errors"
	service "golang-proj-distributed-systems/sevice"
)

var(
	router = gin.Default()
)

func AddBusiness(c *gin.Context) {
	var b *domain.BusinessRep
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil{

	}
	service.AddBusiness(b)
	c.JSON(201, b)
}

func GetBusinessById(c *gin.Context) {
	business, err := service.GetBusinessById(c.Param("id"))
	if err != nil {
		c.JSON(500, errors.BusinessServiceError{Errorcode: "errorCode100" , Message:"business " + c.Param("id") + " cannot be fetched please see the " +
			"cause", Cause: err.Error() })
	}
	c.JSON(200, business)
}

func GetBusinessByFilter(c *gin.Context){
	filterIds :=make(map[string]string)
	filterIds["id"] = c.Param("id")
	filterIds["city"] = c.Query("city")
	filterIds["state"] = c.Query("state")
	logrus.Info("parameters passed for the get request, id: " + c.Param("id") + " " + c.Query("city") + " " + c.Query("state"))
	business := service.GetBusinessData(filterIds)
	logrus.Info(business)
	c.JSON(200, business)

}

func UpdateBusiness(c *gin.Context){
	var b *domain.BusinessRep
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil{

	}
	service.UpdateBusiness(b)
	c.JSON(200,b)
}

func DeleteBusiness(c *gin.Context){
	filterIds :=make(map[string]string)
	filterIds["id"] = c.Param("id")
	filterIds["city"] = c.Query("city")
	filterIds["state"] = c.Query("state")
	service.DeleteBusiness(filterIds)
	c.JSON(200, gin.H{
		"message": "Business with id: " + c.Params.ByName("id") + " deleted successfully",
	})
}
