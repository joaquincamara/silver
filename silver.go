package silver

import (
	"fmt"
	"log"
	"net/http"
)

/*func main() {
	router := AlchemyDoor()
	router.Use(middleware.BasicAuth)
	router.GET("/", HomeRoute)
	log.Fatal(http.ListenAndServe(":8080", router))
} */

//
func AlchemyDoor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome fellow Alchemist to your first route by Silver.go")
}

// Router serves http
type Router struct {
	handlers    map[string]func(http.ResponseWriter, *http.Request)
	middlewares []func(http.Handler) http.Handler
}

// NewRouter creates instance of Router
func TransmuteRouter() *Router {
	router := new(Router)
	router.handlers = make(map[string]func(http.ResponseWriter, *http.Request))
	return router
}

// ServeHTTP is called for every connection
func (s *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := s.handlers[key(r.Method, r.URL.Path)]
	if !ok {
		bad(w)
		return
	}
	f(w, r)
}

func (s *Router) Start(port string, router *Router) {
	p := ":" + port
	log.Fatal(http.ListenAndServe(p, router))
}

// Catch all the middlewares, with a error message if there are not handlers implemented
func (s *Router) Use(middlewares ...func(http.Handler) http.Handler) {
	if s.handlers != nil {
		panic("Silver: all middlewares must be defined before routes ")
	}
	s.middlewares = append(s.middlewares, middlewares...)
}

// GET sets get handler
func (s *Router) GET(path string, f http.HandlerFunc) {
	s.handlers[key("GET", path)] = f
}

// POST sets post handler
func (s *Router) POST(path string, f http.HandlerFunc) {
	s.handlers[key("POST", path)] = f
}

// DELETE sets delete handler
func (s *Router) DELETE(path string, f http.HandlerFunc) {
	s.handlers[key("DELETE", path)] = f
}

// PUT sets put handler
func (s *Router) PUT(path string, f http.HandlerFunc) {
	s.handlers[key("PUT", path)] = f
}

// PATCH set patch handler
func (s *Router) PATCH(path string, f http.HandlerFunc) {
	s.handlers[key("PATCH", path)] = f
}

func key(method, path string) string {
	return fmt.Sprintf("%s:%s", method, path)
}

func bad(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"error":"not found"}`))
}
