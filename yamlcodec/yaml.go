package yamlcodec

import (
	"context"
	"io"
	"net/http"

	"github.com/JoeReid/apiutils"
	"gopkg.in/yaml.v2"
)

type yamlCodec struct{}

func (y *yamlCodec) encoder(w io.Writer) *yaml.Encoder { return yaml.NewEncoder(w) }

func (y *yamlCodec) decoder(r io.Reader) *yaml.Decoder { return yaml.NewDecoder(r) }

func (y *yamlCodec) Respond(ctx context.Context, w http.ResponseWriter, code int, data interface{}) {
	if data != nil {
		// Special case for calling with an error type
		if err, ok := data.(error); ok {
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, err.Error(), code)
			return
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/x-yaml")

		err := y.encoder(w).Encode(data)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(code)
	}
}

func (y *yamlCodec) Read(ctx context.Context, r *http.Request, data interface{}) error {
	return y.decoder(r.Body).Decode(data)
}

func New() apiutils.Codec {
	return &yamlCodec{}
}
