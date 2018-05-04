package newrelicutil_test

import (
	"github.com/gettaxi/newrelicutil"
	"net/http"
)

func ExampleWrapHandlerCtx() {
	nrapp := newApp()
	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
	wrapped_handler := newrelicutil.WrapHandlerCtx(nrapp, "HandlerName", handler)
	mux := http.NewServeMux()
	mux.Handle("/some/path", wrapped_handler)
}
