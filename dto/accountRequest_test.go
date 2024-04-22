package dto

import (
	"testing"
)

func Test_should_return_error_when_account_type_is_different_to_epargne_or_courant(t *testing.T) {
	// Arrange
	request := AccountRequest{
		AccountType: "invalid_account_type",
		Amount:      2000.00,
	}
	// Act
	appException := request.Validate()

	// Assert
	if appException.Message != "Account type should be `courant` or `epargne`" {
		t.Error(" Invalid message while testing account type")
	}
}

func Test_should_return_error_when_amount_is_less_than_min_authorized_amount(t *testing.T) {
	// Arrange
	request := AccountRequest{
		AccountType: "epargne",
		Amount:      123.45,
	}
	// Act
	appException := request.Validate()

	// Assert
	if appException.Message != "To open a new account you need to deposit atleast 1000.00" {
		t.Error(" Invalid message while testing amount opening value")
	}
}
