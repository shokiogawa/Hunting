package main

import (
	"awesomeProject/src/app/infrastructure"
	"awesomeProject/src/app/infrastructure/repository"
	"awesomeProject/src/app/presentation"
	controller2 "awesomeProject/src/app/presentation/controller"
	"awesomeProject/src/app/usecase/command"
)

func main() {
	//TODO: DI
	mysql, err := infrastructure.NewMysql()
	if err != nil {
		return
	}
	prepo := repository.NewPracticeRepository(*mysql)
	prusecase := *command.NewCreatePracticeUseCase(prepo)
	crepo := repository.NewCompanyRepository(*mysql)
	ccuc := *command.NewCreateCompanyUseCase(crepo)
	controller := controller2.NewOwnerController(prusecase, ccuc)
	lineController := controller2.NewLineController()

	//echoフレームワーク
	e := presentation.NewEchoRouter(controller, lineController)
	e.Logger.Fatal(e.Start(":3000"))
}

//server := presentation.DefaultServer(controller)
//
//if err = http.ListenAndServe(":3000", server); err != nil {
//	log.Fatalf("could not listen on port 3000 : %v", err)
//}
