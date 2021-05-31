package controller

import (
	"awesomeProject/src/app/domain/entity"
	"awesomeProject/src/app/usecase/command"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type OwnerController struct {
	createPracticeUseCase command.CreatePracticeUseCase
}

func NewOwnerController(createPracticeUseCase command.CreatePracticeUseCase,
) *OwnerController {
	oc := new(OwnerController)
	oc.createPracticeUseCase = createPracticeUseCase
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
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, practice)
}

func (oc *OwnerController) CreateCompany(c echo.Context) (err error) {
	userId, err := strconv.Atoi(c.Param("userId"))
	type CompanyMapper struct {
		id       int    `json:"id"`
		userId   int    `json:"user_id"`
		statusId int    `json:"status_id"`
		name     string `json:"name"`
		detail   string `json:"detail"`
	}
	var mapperCompany CompanyMapper
	err = c.Bind(&mapperCompany)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	company := *entity.NewCompany(mapperCompany.id, userId, mapperCompany.statusId, mapperCompany.name, mapperCompany.detail)
	fmt.Println(company)
	//TODO: commandでcompanyを作成。
	return c.JSON(http.StatusOK, "ok")

}
