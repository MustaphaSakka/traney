package dto

import (
	"strings"

	"github.com/MustaphaSakka/traney-lib/exception"
)

type AccountRequest struct {
	ClientId    string  `json:"client_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r AccountRequest) Validate() *exception.AppException {
	if r.Amount < 1000 {
		return exception.ValidationError("To open a new account you need to deposit atleast 1000.00")
	}
	if strings.ToLower(r.AccountType) != "epargne" && strings.ToLower(r.AccountType) != "courant" {
		return exception.ValidationError("Account type should be `courant` or `epargne`")
	}
	return nil
}
