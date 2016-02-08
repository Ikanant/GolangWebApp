package converter

import (
	"models"
	"viewmodels"
)

func ConvertProductsToViewModel(product models.Product) viewmodels.Product {
	
	result := viewmodels.Product{
		Id: product.Id(),
		ImageUrl: product.ImageUrl(),
		Name: product.Name(),
		Type: product.Typ(),
		Description: product.Description(),
		Price: product.Price(),
	}
	
	return result
}