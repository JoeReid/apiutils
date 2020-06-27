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
	hello := &HelloEndpoint{}

	http.Handle("/json", apiutils.HandlerWithCodec(jsoncodec.New(), hello))
	http.Handle("/yaml", apiutils.HandlerWithCodec(yamlcodec.New(), hello))
	http.ListenAndServe(":8080", nil)
}
