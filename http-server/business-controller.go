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
		c.JSON(400,
			errors.BusinessServiceError{Errorcode: "errCode101" , Message: "Error parsing the body, please check if the request body is correct", Cause: ""})
	}
	addFlag, errAdd := service.AddBusiness(b)
	if errAdd != nil && addFlag == false{
		c.JSON(500,
			errors.BusinessServiceError{Errorcode: "errCode102" , Message: "Error adding business " + b.Id, Cause: ""})
	}
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
	if business != nil {
		c.JSON(500, errors.BusinessServiceError{Errorcode: "errorCode100" , Message:"business " + c.Param("id") + " cannot be fetched please see the " +
			"cause", Cause: "" })
	}
	logrus.Info(business)
	c.JSON(200, business)
}

func UpdateBusiness(c *gin.Context){
	var b *domain.BusinessRep
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil{
		c.JSON(400,
			errors.BusinessServiceError{Errorcode: "errCode101" , Message: "Error parsing the body, please check if the request body is correct", Cause: ""})
	}

	updateFlag, errUpdate := service.UpdateBusiness(b)
	if errUpdate != nil && updateFlag == false{
		c.JSON(500,
			errors.BusinessServiceError{Errorcode: "errCode103" , Message: "Error updating business " + b.Id, Cause: ""})
	}
	c.JSON(200,b)
}

func DeleteBusiness(c *gin.Context){
	filterIds :=make(map[string]string)
	filterIds["id"] = c.Param("id")
	deleteFlag, errDelete := service.DeleteBusiness(filterIds)
	if errDelete != nil && deleteFlag == false{
		c.JSON(500,
			errors.BusinessServiceError{Errorcode: "errCode104" , Message: "Error deleting business " + c.Param("id"), Cause: ""})
	}
	c.JSON(204, "")
}
