package app

import(
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartAppliaction() {

	router.Run("8080")
}
