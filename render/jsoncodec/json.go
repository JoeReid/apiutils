package jsoncodec

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/JoeReid/apiutils/render"
)

type jsonCodec struct {
	indentPrefix          string
	indent                string
	disallowUnknownFields bool
}

func (j *jsonCodec) encoder(w io.Writer) *json.Encoder {
	enc := json.NewEncoder(w)
	enc.SetIndent(j.indentPrefix, j.indent)

	return enc
}

func (j *jsonCodec) decoder(r io.Reader) *json.Decoder {
	dec := json.NewDecoder(r)

	if j.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}

	return dec
}

func (j *jsonCodec) Respond(ctx context.Context, w http.ResponseWriter, code int, data interface{}) {
	if data != nil {
		// Special case for calling with an error type
		if err, ok := data.(error); ok {
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, err.Error(), code)
			return
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")

		err := j.encoder(w).Encode(data)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(code)
	}
}

func (j *jsonCodec) Read(ctx context.Context, r *http.Request, data interface{}) error {
	return j.decoder(r.Body).Decode(data)
}

func New(opts ...func(*jsonCodec)) render.Codec {
	jc := &jsonCodec{}

	for _, opt := range opts {
		opt(jc)
	}
	return jc
}

func SetIndent(prefix, indent string) func(j *jsonCodec) {
	return func(j *jsonCodec) {
		j.indentPrefix = prefix
		j.indent = indent
	}
}

func DisallowUnknownFields() func(j *jsonCodec) {
	return func(j *jsonCodec) {
		j.disallowUnknownFields = true
	}
}
