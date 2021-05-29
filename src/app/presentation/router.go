package presentation

import (
	"awesomeProject/src/app/presentation/controller"
	"fmt"
	"net/http"
)

func NewServer(controller *controller.OwnerController) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/practices", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			fmt.Println("MethodGetですよ")
		case http.MethodPost:
			fmt.Println("MeyhodはPOSTですよ")
		case http.MethodDelete:
			fmt.Println("Methodは削除ですよ")
		}
	})
	mux.HandleFunc("/company", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("会社を取得するよ")
	})
	return mux
}

func DefaultServer(controller *controller.OwnerController) *http.ServeMux {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			helloWorld(w, req)
		case http.MethodPost:
			fmt.Println("MeyhodはPOSTですよ")
		case http.MethodDelete:
			fmt.Println("Methodは削除ですよ")
		}
	})
	http.HandleFunc("/users", getUser)
	return http.DefaultServeMux
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}

func getUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}
