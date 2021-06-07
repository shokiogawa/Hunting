package controller

import (
	"awesomeProject/src/app/domain/entity"
	"awesomeProject/src/app/usecase/command"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type OwnerController struct {
	createPracticeUseCase command.CreatePracticeUseCase
	createCompanyUseCase  command.CreateCompanyUserCase
}

func NewOwnerController(createPracticeUseCase command.CreatePracticeUseCase, createCompanyUseCase command.CreateCompanyUserCase) *OwnerController {
	oc := new(OwnerController)
	oc.createPracticeUseCase = createPracticeUseCase
	oc.createCompanyUseCase = createCompanyUseCase
	return oc
}

type PracticeMapper struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//webフレームワークの情報はなるべくcontroller内に収めないといけない
//jsonなどの情報はバックとフロントでの通信するためのもののため、controller内に収める。
func (oc *OwnerController) CreatePractice(c echo.Context) (err error) {

	var mapperPractice PracticeMapper
	err = c.Bind(&mapperPractice)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	practice := *entity.NewPractice(0, mapperPractice.Name, mapperPractice.Age)
	err = oc.createPracticeUseCase.Invoke(practice)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusCreated, practice)
	return
}

func (oc *OwnerController) CreateCompany(c echo.Context) (err error) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	type CompanyMapper struct {
		StatusId int    `json:"status_id"`
		Name     string `json:"name"`
		Detail   string `json:"detail"`
	}
	var mapperCompany CompanyMapper
	err = c.Bind(&mapperCompany)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	company := *entity.NewCompany(0, userId, mapperCompany.StatusId, mapperCompany.Name, mapperCompany.Detail)
	err = oc.createCompanyUseCase.Invoke(company)
	return c.JSON(http.StatusCreated, company)
}
