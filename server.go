package main

import (
	"net/http"
	"unicode"

	"github.com/labstack/echo/v4"
)

type CreditCardRequest struct {
	CreditCardNumbers []string `json:"creditCardNumbers"`
}

type CreditCardResult struct {
	CreditCardNumber string `json:"creditCardNumber"`
	IsValid          bool   `json:"isValid"`
}

type CreditCardResponse struct {
	CreditCard []*CreditCardResult `json:"CreditCardValidations"`
}

func main() {
	e := echo.New()
	/*
		curl -X POST http://localhost:1323/credit_card \
			-H 'Content-Type: application/json' \
			-d '{"creditCardNumbers": ["123", "3379 5135 6110 8795", "3379 5135 6110 8794"]}'
	*/
	e.POST("/credit_card", func(c echo.Context) (err error) {
		results := GetCreditCardValidation(c)

		return c.JSON(http.StatusOK, results)

	})
	e.Logger.Fatal(e.Start(":1323"))

}

func GetCreditCardValidation(c echo.Context) CreditCardResponse {
	credit_card_request := new(CreditCardRequest)
	c.Bind(credit_card_request)

	var credit_card_validations []*CreditCardResult
	for _, credit_card_number := range credit_card_request.CreditCardNumbers {
		credit_card_validations = append(credit_card_validations, &CreditCardResult{
			CreditCardNumber: credit_card_number,
			IsValid:          GetCardValidation(credit_card_number),
		})
	}

	results := CreditCardResponse{CreditCard: credit_card_validations}

	return results
}

func GetCardValidation(number string) bool {
	var factor int = 2
	var sum int = 0
	var product int = 1
	var count int = 0
	for _, ch := range number {
		if unicode.IsDigit(ch) {
			count += 1
			product = (int(ch) - '0') * factor

			if product >= 10 {
				sum = sum + 1 + product%10
			} else {
				sum = sum + product
			}

			if factor == 2 {
				factor = 1
			} else {
				factor = 2
			}
		}
	}

	if sum%10 == 0 && count == 16 {
		return true
	}

	return false
}
