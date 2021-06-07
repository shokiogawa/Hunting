package repository

import "awesomeProject/src/app/domain/entity"

type CompanyRepository interface {
	CreateCompany(company entity.Company) error
}
