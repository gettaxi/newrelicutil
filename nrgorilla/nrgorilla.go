// Package nrgorilla provides New Relic instruments for gorilla/mux tool.
package nrgorilla

import (
	"strings"
	"github.com/gorilla/mux"
)

// RouteName returns the name that would be used as transaction name in New Relic
func RouteName(route *mux.Route) string {
	if nil == route {
		return ""
	}
	if n := route.GetName(); n != "" {
		return n
	}
	if n, _ := route.GetPathTemplate(); n != "" {
		ms, _ := route.GetMethods()
		return strings.Join(ms, "/") + " " + n
	}
	n, _ := route.GetHostTemplate()
	return n
}