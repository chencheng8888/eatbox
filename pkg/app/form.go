package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}
type ValidErrors []*ValidError

func (ve ValidError) Error() string {
	return ve.Message
}
func (ve ValidErrors) Errors() []string {
	var errs []string
	for _, err := range ve {
		errs = append(errs, err.Error())
	}
	return errs
}
func (ve ValidErrors) Error() string {
	return strings.Join(ve.Errors(), ",")
}
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return true, nil
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{Key: key, Message: value})
		}
		return true, errs
	}
	return false, nil
}
