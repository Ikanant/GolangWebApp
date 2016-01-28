package viewmodels

import ()

type About struct {
	Title   string
	Active  string
	Skills  []Skill
	Summary string
	Quote string
	Image string
}

type Skill struct {
	Title string
}

func GetAbout() About {
	result := About{
		Title:   "About me:",
		Active:  "about",
		Quote: "First solve the problem, then write the code",
		Image: "7",
		Summary: `As a software developer I want to work with problems that
					will allow me grow as an engineer and also learn how to become a
					successful leader in today’s tech world. Computer Science is no
					longer a single-player game and I am well aware that in order to
					succeed with a project, collaboration with a group of individuals
					who share the same passion and energy as I do for the field is
					essential. I look forward to meeting colleagues that I can count on
					for help and mentorship. I want to write code for a purpose by
					striving to develop efficient and relevant programs to best help
					clients. I love what I do and I will continue to better myself in
					my work.`,
	}

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

	result.Skills = skillSlice

	return result
}
