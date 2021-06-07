package controller_test

import (
	"awesomeProject/src/app/domain/entity"
	"awesomeProject/src/app/presentation"
	controller2 "awesomeProject/src/app/presentation/controller"
	"awesomeProject/src/app/usecase/command"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type InmemoryDb struct {
}
type TestPracticeRepository struct {
	database InmemoryDb
}

func (tp *TestPracticeRepository) CreatePractice(practice entity.Practice) (err error) {
	return err
}

type TestCompanyRepository struct {
	database InmemoryDb
}

func (tc *TestCompanyRepository) CreateCompany(company entity.Company) (err error) {
	return err
}

func TestOwnerController_CreatePractice(t *testing.T) {
	mysql := new(InmemoryDb)
	testPracticeRepo := new(TestPracticeRepository)
	testPracticeRepo.database = *mysql
	testCompanyRepo := new(TestCompanyRepository)
	testPracticeRepo.database = *mysql
	prusecase := *command.NewCreatePracticeUseCase(testPracticeRepo)
	ccuc := *command.NewCreateCompanyUseCase(testCompanyRepo)
	controller := controller2.NewOwnerController(prusecase, ccuc)
	lineController := controller2.NewLineController()
	e := presentation.NewEchoRouter(controller, lineController)

	tests := []struct {
		name         string
		practice     controller2.PracticeMapper
		responseCode int
	}{
		{
			name:         "path this test",
			practice:     controller2.PracticeMapper{Name: "Test", Age: 20},
			responseCode: http.StatusBadRequest,
		},
		{
			name:         "path",
			practice:     controller2.PracticeMapper{Name: "Test rejection", Age: 21},
			responseCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jsonPractice, err := json.Marshal(test.practice)
			assert.Nil(t, err)
			req := httptest.NewRequest(http.MethodPost, "/practice", bytes.NewBuffer(jsonPractice))
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			//TODO: httpステータスコードの判定
			assertStatus(t, rec.Code, test.responseCode)
			fmt.Println(rec.Body)
		})
	}

}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
