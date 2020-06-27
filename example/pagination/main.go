package main

import (
	"net/http"

	"github.com/JoeReid/apiutils"
	"github.com/JoeReid/apiutils/jsoncodec"
	"github.com/JoeReid/apiutils/yamlcodec"
)

type PaginationExample struct {
	// DB etc...
}

func (p *PaginationExample) ServeCodec(c apiutils.Codec, w http.ResponseWriter, r *http.Request) {
	count, skip, err := apiutils.Paginate(r, apiutils.DefaultCount(10), apiutils.MaxCount(10))
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
	codecSelector, _ := apiutils.NewRequestSelector(
		apiutils.RegisterCodec(jsoncodec.New(), "json", "application/json"),
		apiutils.RegisterCodec(yamlcodec.New(), "yaml", "application/x-yaml"),
	)

	paginate := &PaginationExample{}
	http.Handle("/paginate", apiutils.HandlerWithSelector(codecSelector, paginate))
	http.ListenAndServe(":8080", nil)
}
