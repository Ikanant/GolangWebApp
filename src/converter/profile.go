package converter

import (
	"models"
	"viewmodels"
)

func ConvertMemberlToViewModel(modelMember models.Member) viewmodels.Member {

	result := viewmodels.Member{
		Email: modelMember.Email(),
		Id: modelMember.Id(),
		FirstName: modelMember.FirstName(),
	}

	return result
}