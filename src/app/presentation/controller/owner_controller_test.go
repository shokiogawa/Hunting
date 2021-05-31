package controller

import (
	"awesomeProject/src/app/infrastructure"
	"awesomeProject/src/app/infrastructure/repository"
	"awesomeProject/src/app/presentation"
	"awesomeProject/src/app/usecase/command"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOwnerController_CreatePractice(t *testing.T) {
	mysql, err := infrastructure.NewMysql()
	assert.Error(t, err)
	prepo := repository.NewPracticeRepository(*mysql)
	prusecase := *command.NewCreatePracticeUseCase(prepo)
	controller := NewOwnerController(prusecase)
	lineController := NewLineController()
	e := presentation.NewEchoRouter(controller, lineController)

	tests := []struct {
		name         string
		practice     PracticeMapper
		responseCode int
	}{
		{
			name:         "path this test",
			practice:     PracticeMapper{Name: "Test", Age: 20},
			responseCode: http.StatusCreated,
		},
		{
			name:         "unable to path this test",
			practice:     PracticeMapper{Name: "Test rejection", Age: 21},
			responseCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jsonPractice, err := json.Marshal(test.practice)
			assert.Nil(t, err)
			req := httptest.NewRequest(http.MethodPost, "practice", strings.NewReader(string(jsonPractice)))
			rec := httptest.NewRecorder()
			//レsポンスとリクエストを行う。
			e.ServeHTTP(rec, req)
			//TODO: httpステータスコードの判定
			assert.Equal(t, rec.Code, test.responseCode)
			//TODO: 返り値の判定
			//assert.Equal(t, rec.Body.String(), "")
		})
	}

}
