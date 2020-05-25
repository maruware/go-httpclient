package httpclient

import (
	"encoding/json"
	"io"
)

func JsonBody(data interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(data)
}
