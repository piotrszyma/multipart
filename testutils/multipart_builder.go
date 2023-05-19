package testutils

import (
	"bytes"
	"mime/multipart"
)

func BuildMultipartForm() (*multipart.Form, error) {
	form := &multipart.Form{}

	// Add string field
	formValue := &bytes.Buffer{}
	formWriter := multipart.NewWriter(formValue)

	fieldWriter, err := formWriter.CreateFormField("stringField")
	if err != nil {
		return nil, err
	}
	fieldWriter.Write([]byte("Hello, World!"))


	formWriter.Close()

	return form, nil
}
