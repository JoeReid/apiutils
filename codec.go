package apiutils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type Handler interface {
	ServeCodec(c Codec, w http.ResponseWriter, r *http.Request)
}

type HandlerFunc func(c Codec, w http.ResponseWriter, r *http.Request)

func (h *HandlerFunc) ServeCodec(c Codec, w http.ResponseWriter, r *http.Request) {
	(*h)(c, w, r)
}

func HandlerWithCodec(c Codec, h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeCodec(c, w, r)
	})
}

func HandlerWithSelector(s CodecSelector, h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeCodec(s.For(r), w, r)
	})
}

type CodecSelector interface {
	For(r *http.Request) Codec
}

type requestSelector struct {
	codecKey string
	codecs   map[string]Codec
}

func (r *requestSelector) For(req *http.Request) Codec {
	codecName := req.URL.Query().Get(r.codecKey)

	if c, ok := r.codecs[codecName]; ok {
		return c
	}
	return &unsuportedCodec{codecName}
}

func NewRequestSelector(opts ...func(r *requestSelector)) (*requestSelector, error) {
	const defaultCodecKey = "codec"

	r := &requestSelector{
		codecKey: defaultCodecKey,
		codecs:   make(map[string]Codec),
	}

	for _, opt := range opts {
		opt(r)
	}

	if _, ok := r.codecs[""]; !ok {
		return nil, errors.New("need at least one codec registered")
	}
	return r, nil
}

func SetCodecURLKey(k string) func(r *requestSelector) {
	return func(r *requestSelector) {
		r.codecKey = k
	}
}

func RegisterCodec(c Codec, values ...string) func(r *requestSelector) {
	return func(r *requestSelector) {
		// set default if not already set
		if _, ok := r.codecs[""]; !ok {
			r.codecs[""] = c
		}

		for _, v := range values {
			r.codecs[v] = c
		}
	}
}

func SetDefaultCodec(c Codec) func(r *requestSelector) {
	return func(r *requestSelector) {
		r.codecs[""] = c
	}
}

type unsuportedCodec struct {
	codecName string
}

func (u *unsuportedCodec) Respond(ctx context.Context, w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "text/plain")
	http.Error(w, fmt.Sprintf("unsuported codec: %s", u.codecName), http.StatusBadRequest)
}

func (u *unsuportedCodec) Read(ctx context.Context, r *http.Request, data interface{}) error {
	return fmt.Errorf("unsuported codec: %s", u.codecName)
}

type Codec interface {
	// Encode and write the given data to the response writer with the requested status code
	// Will change the status code apropriately if there is an encoding error
	Respond(ctx context.Context, w http.ResponseWriter, code int, data interface{})

	// Read will attempt to decode the body of a request onto the given data interface
	Read(ctx context.Context, r *http.Request, data interface{}) error
}
