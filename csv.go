package oututils

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/gocarina/gocsv"
)

// StringSlice is an alias for []string, which is correctly encoded into csv.
type StringSlice []string

// MarshalCSV allows StringSlice to implement gocsv.TypeMarshaller interface.
func (ss StringSlice) MarshalCSV() (string, error) {
	return strings.Join([]string(ss), "\n"), nil
}

// StringMap is an alias for map[string]interface{}, which is correctly encoded into csv.
type StringMap map[string]interface{}

// MarshalCSV allows StringMap to implement gocsv.TypeMarshaller interface.
func (sm StringMap) MarshalCSV() (string, error) {
	s := ""
	for k, v := range sm {
		s += fmt.Sprintf("%s: %v\n", k, v)
	}
	return s, nil
}

// csv is csv writer.
type csv struct {
	headerWritten bool
	out           io.Writer
}

// CSV returns a new instance of csv writer.
func CSV(out io.Writer) Writer {
	return &csv{
		out: out,
	}
}

// Write allows csv to implement Writer interface.
func (w *csv) Write(value interface{}) error {
	values := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(value)), 0, 0)

	if !w.headerWritten {
		if err := gocsv.Marshal(values.Interface(), w.out); err != nil {
			return err
		}
		w.headerWritten = true
	}

	values = reflect.Append(values, reflect.ValueOf(value))

	return gocsv.MarshalWithoutHeaders(values.Interface(), w.out)
}
