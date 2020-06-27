package main

import (
	"net/http"
	"time"

	"github.com/JoeReid/apiutils"
	"github.com/JoeReid/apiutils/jsoncodec"
	"github.com/JoeReid/apiutils/yamlcodec"
)

type HelloEndpoint struct {
	// DB etc...
}

func (h *HelloEndpoint) ServeCodec(c apiutils.Codec, w http.ResponseWriter, r *http.Request) {
	type Hello struct {
		Hello string    `json:"hello" yaml:"yamlhello"`
		Time  time.Time `json:"timestamp" yaml:"yamltimestamp"`
	}

	// Respond back with a simple hello world message
	c.Respond(r.Context(), w, http.StatusOK, &Hello{
		Hello: "world",
		Time:  time.Now(),
	})
}

func main() {
	// configure all the codec options
	codecSelector, _ := apiutils.NewRequestSelector(
		apiutils.RegisterCodec(
			jsoncodec.New(), "json", "application/json"),

		apiutils.RegisterCodec(
			jsoncodec.New(jsoncodec.SetIndent("", "\t")),
			"json,pretty", "application/json,pretty"),

		apiutils.RegisterCodec(
			yamlcodec.New(), "yaml", "application/x-yaml"),
	)

	// serve all the codecs on a common endpoint
	hello := &HelloEndpoint{}
	http.Handle("/hello", apiutils.HandlerWithSelector(codecSelector, hello))
	http.ListenAndServe(":8080", nil)
}
