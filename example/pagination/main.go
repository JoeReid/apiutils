package main

import (
	"net/http"

	"github.com/JoeReid/apiutils/paginate"
	"github.com/JoeReid/apiutils/render"
	"github.com/JoeReid/apiutils/render/jsoncodec"
	"github.com/JoeReid/apiutils/render/yamlcodec"
)

type PaginationExample struct {
	// DB etc...
}

func (p *PaginationExample) ServeCodec(c render.Codec, w http.ResponseWriter, r *http.Request) {
	count, skip, err := paginate.Vars(r, paginate.DefaultCount(10), paginate.MaxCount(10))
	if err != nil {
		c.Respond(r.Context(), w, http.StatusBadRequest, err)
		return
	}

	// build a slice of ints to simulate paginated data
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = (count * skip) + i
	}

	// Respond with this fragment of the data
	c.Respond(r.Context(), w, http.StatusOK, data)
}

func main() {
	// configure all the codec options
	codecSelector, _ := render.NewRequestSelector(
		render.RegisterCodec(jsoncodec.New(), "json", "application/json"),
		render.RegisterCodec(yamlcodec.New(), "yaml", "application/x-yaml"),
	)

	paginate := &PaginationExample{}
	http.Handle("/paginate", render.HandlerWithSelector(codecSelector, paginate))
	http.ListenAndServe(":8080", nil)
}
