package espresso

type Brand struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Coffee []Coffee `json:"coffee"`
}

type Coffee struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Tags        []Tag   `json:"tags"`
}

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Dose struct {
	Id       string  `json:"id"`
	BrandId  string  `json:"brandId"`
	CoffeeId string  `json:"coffeeId"`
	Weight   float64 `json:"weight"`
	Grind    float64 `json:"grind"`
	Water    float64 `json:"water"`
	Rating   float64 `json:"rating"`
	Notes    string  `json:"notes"`
}
