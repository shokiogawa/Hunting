package repository

import "awesomeProject/src/app/domain/entity"

type PracticeRepository interface {
	CreatePractice(practice entity.Practice) error
}
