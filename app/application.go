package app

import(
	"github.com/gin-gonic/gin"
	http_server "golang-proj-distributed-systems/http-server"
)

var (
	router = gin.Default()
)

func StartApplication() {

	router.Run("8080")
}

func mapurls() {

	router.GET("/business/id", http_server.GetBusinessById)
	router.POST("/business", http_server.AddBusiness)
	router.PUT("/business/id", http_server.UpdateBusiness)
	router.DELETE("/business/id", http_server.DeleteBusiness)
	router.PATCH("/business/id", http_server.UpdateBusiness)
}
