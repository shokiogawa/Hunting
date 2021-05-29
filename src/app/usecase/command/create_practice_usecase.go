package command

import (
	"awesomeProject/src/app/domain/entity"
	"awesomeProject/src/app/domain/repository"
)

type CreatePracticeUseCase struct {
	repository repository.PracticeRepository
}

func NewCreatePracticeUseCase(repository repository.PracticeRepository) *CreatePracticeUseCase {
	uc := new(CreatePracticeUseCase)
	uc.repository = repository
	return uc
}

func (uc *CreatePracticeUseCase) Invoke(practice entity.Practice) (err error) {
	err = uc.repository.CreatePractice(practice)
	if err != nil {
		return err
	}
	return err
}
