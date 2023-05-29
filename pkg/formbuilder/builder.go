package formbuilder

import (
	"bytes"
	"mime/multipart"
)

type boundary string

type builder struct {
	// form *multipart.Form
	formWriter *multipart.Writer
	buffer *bytes.Buffer
}

func NewMultipart() *builder {
	buffer := bytes.Buffer{}
	formWriter := multipart.NewWriter(&buffer)
	return &builder{formWriter, &buffer}
}

func (m *builder) AddStringField(fieldname, value string) *builder {
	m.formWriter.WriteField(fieldname, value)
	return m;
}

func (m *builder) Build() (*bytes.Buffer, boundary, error) {
	err := m.formWriter.Close()
	if err != nil {
		return nil, "", err
	}

	return m.buffer, boundary(m.formWriter.Boundary()), nil
}