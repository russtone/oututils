package oututils_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/russtone/oututils"
)

func TestNDJSON(t *testing.T) {
	var buf bytes.Buffer

	w := oututils.NDJSON(&buf)

	type Data struct {
		Field1 string `json:"field1"`
	}

	w.Write(&Data{"test"})
	w.Write(&Data{"test2"})

	expected := `{"field1":"test"}
{"field1":"test2"}
`

	assert.Equal(t, expected, buf.String())
}
