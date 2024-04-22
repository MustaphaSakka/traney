package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MustaphaSakka/traney/dto"
	"github.com/MustaphaSakka/traney/exception"
	"github.com/MustaphaSakka/traney/mocks/service"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"
)

var router *mux.Router
var ch ClientHandlers
var mockService *service.MockClientService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockClientService(ctrl)
	ch = ClientHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/clients", ch.getAllClients)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_clients_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	dummyClients := []dto.ClientResponse{
		{"1", "Pierre", "Lyon", "69007", "1990-04-03", "1"},
		{"2", "Samuel", "Paris", "91000", "1993-11-01", "1"},
	}
	mockService.EXPECT().GetAllClient().Return(dummyClients, nil)
	request, _ := http.NewRequest(http.MethodGet, "/clients", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().GetAllClient().Return(nil, exception.InternalServerException("database is HS"))
	request, _ := http.NewRequest(http.MethodGet, "/clients", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
