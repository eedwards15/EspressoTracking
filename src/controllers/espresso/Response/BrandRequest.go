package Response

type BrandRequest struct {
	Name string `json:"name"`
}

type CoffeeRequest struct {
	BrandId     string  `json:"brandId"`
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
