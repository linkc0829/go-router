package route

import (
	"errors"
)

// Message is message
type Message struct {
	Identification string
	Method         string
	Time           string
	Content        string
	Size           int
}

// Handler is handler
type Handler func(res, req *Message)

// Middleware is public middleware
type middleware func(Handler) Handler

// Router is router
type Router struct {
	middleWareChain []middleware
	mux             map[string]Handler
}

// NewRouter new a router
func NewRouter() *Router {
	return &Router{
		mux: make(map[string]Handler),
	}
}

func (r *Router) Use(m middleware) {
	r.middleWareChain = append(r.middleWareChain, m)
}

// Add a route
func (r *Router) add(route string, h Handler) {
	var mergeHandler = h

	for i := len(r.middleWareChain) - 1; i >= 0; i-- {
		mergeHandler = r.middleWareChain[i](mergeHandler)
	}
	r.mux[route] = mergeHandler
}

// Get add a get method pattern
func (r *Router) Get(route string, h Handler) {
	r.add("get:"+route, h)
}

// Put add a put method pattern
func (r *Router) Put(route string, h Handler) {
	r.add("put:"+route, h)
}

// Post add a post method pattern
func (r *Router) Post(route string, h Handler) {
	r.add("post:"+route, h)
}

// Delete add a delete method pattern
func (r *Router) Delete(route string, h Handler) {
	r.add("delete:"+route, h)
}

func (r *Router) Run(res, req *Message) error {
	route := req.Method + ":" + req.Identification
	handler, exists := r.mux[route]
	if exists {
		handler(res, req)
		return nil
	}
	return errors.New("route not exists")
}
