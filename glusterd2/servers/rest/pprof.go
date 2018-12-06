package rest

// Based on ideas from https://github.com/mistifyio/negroni-pprof
import (

	//	"net/http"
	"net/http"
	"net/http/pprof"
	runtime_pprof "runtime/pprof"

	"github.com/gorilla/mux"
)

var (
	basePath string = "/debug/pprof/"
)

func addPath(router *mux.Router, name string, handler http.Handler) {
	router.Path(basePath + name).Name(name).Handler(handler)
	//fmt.Fprintf(os.Stderr, "DEBUG: registered profiling handler on %s\n", basePath+name)
}
func EnableProfiling(router *mux.Router) {
	for _, profile := range runtime_pprof.Profiles() {
		name := profile.Name()
		handler := pprof.Handler(name)
		addPath(router, name, handler)
	}
	// static profiles as listed in net/http/pprof/pprof.go:init()
	addPath(router, "cmdline", http.HandlerFunc(pprof.Cmdline))
	addPath(router, "profile", http.HandlerFunc(pprof.Profile))
	addPath(router, "symbol", http.HandlerFunc(pprof.Symbol))
	addPath(router, "trace", http.HandlerFunc(pprof.Trace))
	// The index handler only lists the runtime_pprof.Profiles() which does
	// not include all the endpoints we have. Thus we opt not to use the
	// common prof.Index endpoint here.
}
