package testingcodec

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type testingCodec struct {
	mock.Mock
}

func (t *testingCodec) Respond(ctx context.Context, w http.ResponseWriter, code int, data interface{}) {
	t.MethodCalled("Respond", ctx, w, code, data)
}

func (t *testingCodec) Read(ctx context.Context, r *http.Request, data interface{}) error {
	args := t.MethodCalled("Read", ctx, r, data)

	// TODO: there should be a method here to write data to the data interface
	// reflection? unsafe?
	// for now we can just not support tests that read data from the request

	return args.Error(0)
}

func New() *testingCodec { return &testingCodec{} }
