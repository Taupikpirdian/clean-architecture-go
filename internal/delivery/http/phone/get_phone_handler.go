package phone

import (
	response2 "challenge/clean-architecture-go/internal/delivery/response"
	"challenge/clean-architecture-go/pkg/response"
	"encoding/json"
	"net/http"
)

/*
	- nama method disini bebas, ini yang akan di panggil di routes, tidak bergantung dengan interface
	- dibagian sini hanya berkaitan dengan sisi usecasenya apa, lalu kita ubah output expectednya
*/
func (handler *PhoneHandlerInteractor) GetPhoneController(w http.ResponseWriter, r *http.Request) {
	listPhone, errGetPhoneFromUseCase := handler.PhoneUseCase.GetPhone()

	if errGetPhoneFromUseCase != nil {
		// error status here
	}

	var phonesResponse = make([]response2.PhoneAttributs, 0)
	for _, phone := range listPhone {
		phoneResponse := response2.NewPhoneResponse(phone)
		phonesResponse = append(phonesResponse, phoneResponse)
	}

	responseJson, _ := json.Marshal(response.NewJsonResponse(phonesResponse, "SUCCESS", 0))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
