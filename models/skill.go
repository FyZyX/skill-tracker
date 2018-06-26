package models

type Skill struct{
	Name string
	mastery float32
}

func NewSkill(name string, mastery float32) Skill {
	return Skill{
		Name: name,
		mastery: mastery,
	}
}

func (skill *Skill) SetMastery(mastery float32) {
	skill.mastery = mastery
}

func (skill Skill) Mastery() int {
	return int(skill.mastery * 100)
}
