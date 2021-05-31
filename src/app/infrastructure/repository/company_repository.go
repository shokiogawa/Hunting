package repository

import (
	"awesomeProject/src/app/domain/entity"
	"awesomeProject/src/app/infrastructure"
)

type CompanyRepository struct {
	database infrastructure.Mysql
}

func NewCompanyRepository(database infrastructure.Mysql) *CompanyRepository {
	cr := new(CompanyRepository)
	cr.database = database
	return cr
}

func (cr CompanyRepository) CreateCompany(company entity.Company) (err error) {
	db, err := cr.database.Connect()
	if err != nil {
		return
	}
	query := `INSERT INTO companies(user_id, status_id, name, detail, color)
VALUE (?,?,?,?,?);`
	result := db.MustExec(query, company.UserId, company.StatusId, company.Name, company.Detail, company.Color)
	_, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
