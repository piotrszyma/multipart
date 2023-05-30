package fileutil

import (
	"io"
	"strings"
)

func ToString(file io.Reader) string {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, file)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
