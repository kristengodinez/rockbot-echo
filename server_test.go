package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreditCardValidations(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"creditCardNumbers": ["123", "3379 5135 6110 8795", "3379 5135 6110 8794"]}`))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	expected := `{"CreditCardValidations":[{"creditCardNumber":"123","isValid":false},{"creditCardNumber":"3379 5135 6110 8795","isValid":true},{"creditCardNumber":"3379 5135 6110 8794","isValid":false}]}`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expected, recorder.Body.String())
}
