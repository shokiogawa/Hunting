package repository

import (
	"awesomeProject/src/app/domain/entity"
	"awesomeProject/src/app/infrastructure"
)

type PracticeRepository struct {
	database infrastructure.Mysql
}

func NewPracticeRepository(database infrastructure.Mysql) *PracticeRepository {
	repo := new(PracticeRepository)
	repo.database = database
	return repo
}

func (pr *PracticeRepository) CreatePractice(practice entity.Practice) (err error) {

	db, err := pr.database.Connect()
	if err != nil {
		return
	}

	query := `INSERT INTO practices(name, age) VALUE(?,?)`
	result := db.MustExec(query, practice.Name, practice.Age)
	_, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
