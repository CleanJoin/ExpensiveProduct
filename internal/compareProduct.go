package internal

import "strconv"

func CompareJSON(listProduct []map[string]interface{}) *Product {
	product := new(Product)
	for _, k := range listProduct {

		if product.Price <= int(k["price"].(float64)) && product.Rating < int(k["rating"].(float64)) {
			product.Product = k["product"].(string)
			product.Price = int(k["price"].(float64))
			product.Rating = int(k["rating"].(float64))
		}

	}
	return product

}

func CompareCSV(listProduct [][]string) *Product {

	product := new(Product)
	for i, listProduct := range listProduct {
		if i != 0 {
			listProductPrice, _ := strconv.Atoi(listProduct[1])
			listProductRating, _ := strconv.Atoi(listProduct[2])
			if product.Price <= listProductPrice && product.Rating < listProductRating {
				product.Product = listProduct[0]
				product.Price = listProductPrice
				product.Rating = listProductRating
			}
		}

	}
	return product

}
