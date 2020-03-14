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
	var b *domain.Business
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil{
		//errors.NewB
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
	filterIds["city"] = c.Param("city")
	filterIds["state"] = c.Param("state")
	//business := service.GetBusinessData(filterIds)
	//err := json.NewEncoder(c.).Encode(business)
	c.JSON(200, gin.H{
		"message": "Implement me ",
	})
}

func UpdateBusiness(c *gin.Context){
	var b *domain.Business
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil{

	}
	//service.UpdateBusiness(b)
	c.JSON(200, gin.H{
		"message": "Implement me ",
	})
}

func DeleteBusiness(c *gin.Context){
	//service.DeleteBusiness(c.Params.ByName("id"))
	c.JSON(200, gin.H{
		"message": "Business with id: " + c.Params.ByName("id") + " deleted successfully",
	})
}
