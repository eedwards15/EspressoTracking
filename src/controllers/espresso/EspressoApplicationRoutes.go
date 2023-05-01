package espresso

import (
	"github.com/gin-gonic/gin"
)

func EspressoApplicationRoutes(path string, app *gin.Engine) {
	BrandController(path, app)
	CoffeeController(path, app)

}
