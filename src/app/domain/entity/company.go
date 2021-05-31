package entity

type Company struct {
	Id       int
	UserId   int
	StatusId int
	Name     string
	Detail   string
	Color    CompanyColor
}

func NewCompany(id int, userId int, statusId int, name string, detail string) *Company {
	company := new(Company)
	company.Id = id
	company.UserId = userId
	company.StatusId = statusId
	company.Name = name
	company.Detail = detail
	company.Color = CompanyColors.Black
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
