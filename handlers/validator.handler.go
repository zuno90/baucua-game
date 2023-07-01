package handlers

import (
	"github.com/go-playground/validator"
	st "github.com/zuno90/go-ws/structs"
)



var validate = validator.New()

func ValidateStruct[T any](vT T) []*st.ErrorResponse {
	errors := []*st.ErrorResponse{}
	err := validate.Struct(vT)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := &st.ErrorResponse{}
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
	}
	if len(errors) == 0 {
		return nil
	}
	return errors
}
