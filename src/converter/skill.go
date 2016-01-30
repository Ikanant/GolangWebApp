package converter

import (
	"models"
	"viewmodels"
)

func ConvertSkillToViewModel(modelSkills models.Skill) viewmodels.Skill {

	result := viewmodels.Skill{
		Skill: modelSkills.Skill(),
	}

	return result

}
