package multipart_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/piotrszyma/multipart"
	"github.com/piotrszyma/multipart/testutils"
)

func TestMultipartRequest(t *testing.T) {
	// Create a new multipart buffer
	form, err := testutils.BuildMultipartForm()
	require.NoError(t, err)

	var foo struct {
		Bar string `multipart:"bar"`
	}

	multipart.Bind(form, &foo)

	require.Equal(t, "Hello", foo.Bar)

}
