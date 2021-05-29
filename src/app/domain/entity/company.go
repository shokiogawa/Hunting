package entity

type Company struct {
	id       int
	userId   int
	statusId int
	name     string
	detail   string
	color    CompanyColor
}

func NewCompany(id int, userId int, statusId int, name string, detail string) *Company {
	company := new(Company)
	company.id = id
	company.userId = userId
	company.statusId = statusId
	company.name = name
	company.detail = detail
	company.color = CompanyColors.Black
	return company
}

func (colors CompanyColor) String() string {
	return colors.value
}

type CompanyColor struct {
	value string
}

var CompanyColors = struct {
	Red    CompanyColor
	Blue   CompanyColor
	Yellow CompanyColor
	Green  CompanyColor
	Orange CompanyColor
	Black  CompanyColor
}{
	Red:    CompanyColor{"RED"},
	Blue:   CompanyColor{"BLUE"},
	Yellow: CompanyColor{"Yellow"},
	Green:  CompanyColor{"GREEN"},
	Orange: CompanyColor{"ORANGE"},
	Black:  CompanyColor{"BLACK"},
}
