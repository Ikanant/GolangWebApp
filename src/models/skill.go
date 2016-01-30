package models

import ()

type Skill struct {
	skill string
}

func (this *Skill) Skill() string {
	return this.skill
}
func (this *Skill) setSkill(value string) {
	this.skill = value
}

func GetSkills() []Skill {
	skillSlice := []Skill{
		Skill{"Java"},
		Skill{"Spring MVC"},
		Skill{"Maven"},
		Skill{"HTML"},
		Skill{"JavaScript"},
		Skill{"CSS"},
		Skill{"Web Services"},
		Skill{"Docker"},
		Skill{"Golang"},
		Skill{"Git"},
		Skill{"Sql"},
	}
	
	return skillSlice;
}
