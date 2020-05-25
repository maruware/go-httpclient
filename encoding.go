package httpclient

import (
	"encoding/json"
	"io"
)

func EncodeJson(data interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(data)
}

func DecodeJson(ref interface{}, r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(ref)
}
