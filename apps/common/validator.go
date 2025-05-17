package common

import (
	"errors"
	"fmt"
	"gopkg/common"
	"gopkg/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

/*
   Purpose : This method is used to validate request data from the client and validate token is valid or not.
   Parameter : pDebug - *helpers.HelperStruct, pRequestData RequestStruct
   Response :

   On Success:
   ===========
	In case of a successful execution of this method, you will get nil as a response.

   On Error:
   ===========
   	In case of any exception during the execution of this method you will send the error details.

   Author : VIJAY
   Date : 01-04-2025
*/
// ValidateRequest validates input data
func ValidateRequest(log *utils.Logger, pRequestData any, pHttpRequest *http.Request) error {
	log.Log(common.INFO, "GlobalDBInit (+)")
	lValidate := validator.New()
	lErr := lValidate.Struct(pRequestData)
	if lErr != nil {
		var validationErrors string
		for _, err := range lErr.(validator.ValidationErrors) {
			validationErrors += fmt.Sprintf("The field '%s' failed validation: it must satisfy the '%s' rule%s.",
				err.Field(),
				err.ActualTag(),
				func() string {
					if err.Param() != "" {
						return fmt.Sprintf(" with parameter '%s'", err.Param())
					}
					return ""
				}(),
			)
		}
		log.Log(common.ERROR, "Validation failed: ", validationErrors)
		log.Log(common.ERROR, "ValidateRequest : 001 (PCVR-001) request validation failed", lErr.Error())
		return errors.New(validationErrors)
	}
	log.Log(common.INFO, "ValidateRequest (-)")
	return nil
}
