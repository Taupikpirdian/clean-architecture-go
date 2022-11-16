package phone_test

import (
	"clean-architecture-go/internal/delivery/http/phone"
	"clean-architecture-go/internal/usecase/mocks"
	"clean-architecture-go/testdata"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	useCasePhoneMocks = new(mocks.PhoneUseCaseMocks)
	req               = httptest.NewRequest("GET", "http://example.com/foo", nil)
	w                 = httptest.NewRecorder()
)

func TestHandlerGetPhone(t *testing.T) {
	useCasePhoneMocks.On("GetPhone").Return(testdata.GenerateListPhoneByEntity(), nil)

	handlerInteractorPhone := phone.NewPhoneHandler(useCasePhoneMocks)
	handlerInteractorPhone.GetPhoneController(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotNil(t, string(body))
}
