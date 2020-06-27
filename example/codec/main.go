package main

import (
	"net/http"
	"time"

	"github.com/JoeReid/apiutils/render"
	"github.com/JoeReid/apiutils/render/jsoncodec"
	"github.com/JoeReid/apiutils/render/yamlcodec"
)

type HelloEndpoint struct {
	// DB etc...
}

func (h *HelloEndpoint) ServeCodec(c render.Codec, w http.ResponseWriter, r *http.Request) {
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

	http.Handle("/json", render.HandlerWithCodec(jsoncodec.New(), hello))
	http.Handle("/yaml", render.HandlerWithCodec(yamlcodec.New(), hello))
	http.ListenAndServe(":8080", nil)
}
