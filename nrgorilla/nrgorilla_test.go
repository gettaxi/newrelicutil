package nrgorilla_test

import (
	"testing"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gettaxi/newrelicutil/nrgorilla"
)

func TestRouteName(t *testing.T) {
	r := mux.NewRouter()
	h := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})

	var tt = []struct {
		route *mux.Route
		exp string
	}{
		{
			route: nil,
			exp: "",
		},{
			route: r.Handle("/api/", h).Methods("GET").Name("FOO"),
			exp: "FOO",
		},{
			route: r.Handle("/api/v1/{env}/callers/{phone}", h).Methods("GET"),
			exp: "GET /api/v1/{env}/callers/{phone}",
		},{
			route: r.Methods("GET"),
			exp: "",
		},
	}
	for _, tc := range tt {
		if want, have := tc.exp, nrgorilla.RouteName(tc.route); want != have {
			t.Errorf("want %+v, have %+v", want, have)
		}
	}
}