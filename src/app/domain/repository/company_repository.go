package repository

import "awesomeProject/src/app/domain/entity"

type CompanyRepository interface {
	CreateCompany(entity.Company) error
}
