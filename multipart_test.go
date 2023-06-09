package multipart_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/piotrszyma/multipart"
	"github.com/piotrszyma/multipart/pkg/formbuilder"
)

func ReadToString(r io.Reader) string {
	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func TestBindMultipart_StringPointer_Binds(t *testing.T) {
	// Arrange.
	var formData struct {
		StringPtr1 *string `multipart:"stringPtr1"`
		StringPtr2 *string `multipart:"stringPtr2"`
	}

	rawForm, err := formbuilder.
		NewMultipart().
		AddStringField("stringPtr1", "foo").
		AddStringField("stringPtr2", "bar").
		Build()

	require.NoError(t, err)

	form, err := rawForm.IntoForm()
	require.NoError(t, err)

	// Act.
	err = multipart.Bind(form, &formData)
	require.NoError(t, err)

	// Assert.
	require.Equal(t, "foo", *formData.StringPtr1)
	require.Equal(t, "bar", *formData.StringPtr2)
}

func TestMultipartRequest_String_Binds(t *testing.T) {
	// Arrange.
	var formData struct {
		String1 string `multipart:"string1"`
		String2 string `multipart:"string2"`
	}
	rawForm, err := formbuilder.
		NewMultipart().
		AddStringField("string1", "foo").
		AddStringField("string2", "bar").
		Build()
	require.NoError(t, err)

	form, err := rawForm.IntoForm()
	require.NoError(t, err)

	// Act.
	err = multipart.Bind(form, &formData)
	require.NoError(t, err)

	// Assert.
	require.Equal(t, "foo", formData.String1)
	require.Equal(t, "bar", formData.String2)
}

func TestMultipartRequest_File_Binds(t *testing.T) {
	// Arrange.
	var formData struct {
		File1 *multipart.FileHeader `multipart:"file1"`
	}
	rawForm, err := formbuilder.
		NewMultipart().
		AddFileField("file1", "file.txt", []byte("test file data")).
		Build()
	require.NoError(t, err)

	form, err := rawForm.IntoForm()
	require.NoError(t, err)

	// Act.
	err = multipart.Bind(form, &formData)
	require.NoError(t, err)

	// Assert.
	require.Equal(t, "file.txt", formData.File1.Filename)
	file, err := formData.File1.Open()
	require.NoError(t, err)
	require.Equal(t, "test file data", ReadToString(file))
}
