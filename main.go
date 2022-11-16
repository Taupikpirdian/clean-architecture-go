package main

import (
	phone_handler "clean-architecture-go/internal/delivery/http/phone"
	"clean-architecture-go/internal/repository/mocks"
	internal_usecase "clean-architecture-go/internal/usecase/phone"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	repoMocks    = new(mocks.PhoneRepositoryMock)
	useCasePhone = internal_usecase.NewPhoneUseCaseInteractor(repoMocks)
)

func main() {
	/*
		DETAIL APLIKASI ADA PADA FILE README.md
	*/
	// os package
	r := mux.NewRouter()

	// routes
	handlerPhone := phone_handler.NewPhoneHandler(useCasePhone)
	r.HandleFunc("/list-phone", handlerPhone.GetPhoneController).Methods(http.MethodGet)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}
