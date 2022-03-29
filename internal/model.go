package internal

type Product struct {
	Product string `json:"product"`
	Price   int    `json:"price"`
	Rating  int    `json:"rating"`
}

type IProduct interface {
	CompareJSON(listProduct map[string]interface{}) *Product
	CompareCSV(listProduct []string) *Product
}

func NewProduct() *Product {
	product := new(Product)
	return product
}
