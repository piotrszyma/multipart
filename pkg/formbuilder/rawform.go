package formbuilder

import (
	"bytes"
	"mime/multipart"
)

type rawForm struct {
	boundary boundary
	buffer   *bytes.Buffer
}

func (r *rawForm) IntoForm() (*multipart.Form, error) {
	form, err := multipart.
		NewReader(r.buffer, string(r.boundary)).
		ReadForm(2048)

	if err != nil {
		return nil, err
	}

	return form, nil
}
