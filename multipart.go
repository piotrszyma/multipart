package multipart

import (
	"mime/multipart"
	"reflect"

	"github.com/pkg/errors"
)


func Bind(form *multipart.Form, target interface{}) error {
	multipartFields := map[string][]string{}

	for fieldName, fieldValues := range form.Value {
		multipartFields[fieldName] = fieldValues
	}

	formHeaders := map[string][]*multipart.FileHeader{}
	for fieldName, fieldHeaders := range form.File {
		formHeaders[fieldName] =  fieldHeaders
	}

	v := reflect.ValueOf(target)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}


	if v.Kind() != reflect.Struct {
		return errors.New("target is not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)

		if !field.IsExported() {
			// Skip non-public fields
			continue
		}

		tagValue := field.Tag.Get("multipart")
		if tagValue == "" {
			// Skip fields without multipart tag set.
			continue
		}

		structFieldValue := v.FieldByName(field.Name)

		formFieldValue, ok := multipartFields[tagValue]
		if !ok {
			return errors.Errorf("field %s is missing", tagValue)
		}

		if structFieldValue.Kind() == reflect.Ptr {
			value := formFieldValue[0]
			structFieldValue.Set(reflect.ValueOf(&value))
		} else {
			structFieldValue.Set(reflect.ValueOf(formFieldValue[0]))
		}

	}
	return nil
}
