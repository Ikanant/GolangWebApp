package converter

import (
	"models"
	"viewmodels"
)

func ConvertCategoryToViewModel(category models.Category) viewmodels.Category {
	
	result := viewmodels.Category{
		ImageUrl: category.ImageUrl(),
		Id: category.Id(),
		SampleProducts: category.SampleProducts(),
		Title: category.Title(),
	}
	
	return result
}