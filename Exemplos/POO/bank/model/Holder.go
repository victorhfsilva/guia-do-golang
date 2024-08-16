package model

type Holder struct {
	name       string
	cpf        string
	profession string
}

func CreateHolder(name string, cpf string, profession string) Holder {
	return Holder{
		name:       name,
		cpf:        cpf,
		profession: profession,
	}
}

func (h Holder) GetName() string {
	return h.name
}

func (h Holder) GetCpf() string {
	return h.cpf
}

func (h Holder) GetProfession() string {
	return h.profession
}

func (h *Holder) SetName(name string) {
	h.name = name
}

func (h *Holder) SetProfession(profession string) {
	h.profession = profession
}
