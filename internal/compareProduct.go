package internal

import "strconv"

func (product *Product) CompareJSON(listProduct map[string]interface{}) *Product {

	if product.Price <= int(listProduct["price"].(float64)) && product.Rating < int(listProduct["rating"].(float64)) {
		product.Product = listProduct["product"].(string)
		product.Price = int(listProduct["price"].(float64))
		product.Rating = int(listProduct["rating"].(float64))
	}
	return product

}

func (product *Product) CompareCSV(listProduct []string) *Product {

	listProductPrice, _ := strconv.Atoi(listProduct[1])
	listProductRating, _ := strconv.Atoi(listProduct[2])
	if product.Price <= listProductPrice && product.Rating < listProductRating {
		product.Product = listProduct[0]
		product.Price = listProductPrice
		product.Rating = listProductRating
	}

	return product

}
