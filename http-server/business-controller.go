package http_server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-proj-distributed-systems/domain"
	service "golang-proj-distributed-systems/sevice"
)

var(
	router = gin.Default()
)

func AddBusiness(c *gin.Context) {
	var b domain.Business
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil{
		//errors.NewB
	}
	service.AddBusiness()
	c.JSON(201, gin.H{
			"message": "Business added successfully",
		})
}

func GetBusiness(c *gin.Context){

}

func GetBusinessById(c *gin.Context){}

func UpdateBusiness(c *gin.Context){}

func DeleteBusiness(c *gin.Context){}
