package oututils_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/russtone/oututils"
)

func TestCSV(t *testing.T) {
	var buf bytes.Buffer

	w := oututils.CSV(&buf)

	type Data struct {
		Field1 string               `csv:"field1"`
		Field2 oututils.StringSlice `csv:"field2"`
		Field3 oututils.StringMap   `csv:"field3"`
	}

	w.Write(&Data{
		Field1: "test",
		Field2: []string{"test1", "test2"},
		Field3: map[string]interface{}{
			"test1": "test2",
		},
	})

	expected := `field1,field2,field3
test,"test1
test2","test1: test2
"
`

	assert.Equal(t, expected, buf.String())
}
