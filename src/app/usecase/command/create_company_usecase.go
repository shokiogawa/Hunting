package command

import (
	"awesomeProject/src/app/domain/entity"
	"awesomeProject/src/app/domain/repository"
)

type CreateCompanyUserCase struct {
	repository repository.CompanyRepository
}

func NewCreateCompanyUseCase(repository repository.CompanyRepository) *CreateCompanyUserCase {
	uc := new(CreateCompanyUserCase)
	uc.repository = repository
	return uc
}

func (uc *CreateCompanyUserCase) Invoke(company entity.Company) (err error) {
	err = uc.repository.CreateCompany(company)
	if err != nil {
		return
	}
	return
}
