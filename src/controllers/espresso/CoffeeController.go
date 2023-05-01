package espresso

import (
	"EspressoTracking/controllers/espresso/Response"
	"EspressoTracking/database/espresso"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CoffeeController(path string, app *gin.Engine) {
	routes := app.Group(path + "/coffee")
	routes.POST("/create", AddCoffee)
	routes.POST("/delete", DeleteCoffee)

}

func AddCoffee(c *gin.Context) {
	var coffee Response.CoffeeRequest
	c.BindJSON(&coffee)

	coffee.Id = uuid.New().String()

	//map request coffee to database model coffee
	recordCoffee := espresso.Coffee{
		Id:          coffee.Id,
		Name:        coffee.Name,
		Price:       coffee.Price,
		Description: coffee.Description,
		Tags:        []espresso.Tag{},
	}

	coffeeRepo := espresso.NewEspressoRepository()
	_, err := coffeeRepo.AddCoffee(coffee.BrandId, recordCoffee)

	if err != nil {
		handleError(c, err)
	}

	c.JSON(200, coffee)
}

func DeleteCoffee(c *gin.Context) {
	var coffee Response.CoffeeRequest
	c.BindJSON(&coffee)

	coffeeRepo := espresso.NewEspressoRepository()
	_, err := coffeeRepo.DeleteCoffee(coffee.BrandId, coffee.Id)
	if err != nil {
		handleError(c, err)
	}

}
