package oututils

import (
	"encoding/json"
	"io"
)

const (
	lf = 0x0A
)

// ndjson is newline delimited JSON writer, see: http://ndjson.org.
type ndjson struct {
	out io.Writer
	enc *json.Encoder
}

// NDJSON returns new ndjson Writer.
func NDJSON(out io.Writer) Writer {
	enc := json.NewEncoder(out)
	enc.SetEscapeHTML(false)
	return &ndjson{
		out: out,
		enc: enc,
	}
}

// Write allows ndjson to implement Writer interface.
func (w *ndjson) Write(value interface{}) error {
	return w.enc.Encode(value)
}
