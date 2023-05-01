package espresso

import (
	"EspressoTracking/controllers/espresso/Response"
	"EspressoTracking/database/errorlogging"
	"EspressoTracking/database/espresso"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"runtime"
)

func BrandController(path string, app *gin.Engine) {
	routes := app.Group(path + "/brand")
	routes.POST("/create", CreateBrand)
	routes.GET("/all", GetBrands)
	routes.GET("/:id", GetBrandById)
	routes.PUT("/:id", UpdateBrand)
	routes.DELETE("/:id", DeleteBrand)
}

func CreateBrand(c *gin.Context) {
	var brand Response.BrandRequest
	c.BindJSON(&brand)

	recordBrand := espresso.Brand{
		Id:     uuid.New().String(),
		Name:   brand.Name,
		Coffee: []espresso.Coffee{},
	}

	brandRepo := espresso.NewEspressoRepository()
	err := brandRepo.CreateBrand(recordBrand)

	if err != nil {
		handleError(c, err)
	}

	c.JSON(200, brand)
}

func GetBrands(c *gin.Context) {
	brandRepo := espresso.NewEspressoRepository()
	brands := brandRepo.GetAllBrands()
	c.JSON(200, brands)
}

func GetBrandById(c *gin.Context) {
	id := c.Param("id")
	brandRepo := espresso.NewEspressoRepository()
	brands, err := brandRepo.GetById(id)
	if err != nil {
		handleError(c, err)
	}
	c.JSON(200, brands)
}

func UpdateBrand(c *gin.Context) {
	id := c.Param("id")
	brandRepo := espresso.NewEspressoRepository()
	brands, err := brandRepo.GetById(id)
	if err != nil {
		handleError(c, err)
	}

	var brand Response.BrandRequest
	c.BindJSON(&brand)

	brands.Name = brand.Name
	_, err = brandRepo.UpdateBrand(brands)
	if err != nil {
		handleError(c, err)
	}

	c.JSON(200, brands)
}

func DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	brandRepo := espresso.NewEspressoRepository()
	err := brandRepo.DeleteBrand(id)
	if err != nil {
		handleError(c, err)
	}
	c.JSON(200, gin.H{"message": "Brand deleted"})
}

func handleError(c *gin.Context, err error) {
	ec := errorlogging.NewErrorLogRepository()
	_, file, line, _ := runtime.Caller(0)
	errRecord := errorlogging.NewErrorRecord(c.ClientIP(), file+" "+string(line), err.Error(), err.Error())
	ec.CreateErrorLogRecord(*errRecord)
	c.JSON(500, gin.H{"error": "Internal Error"})
}
