package main

import (
	"fmt"
	"router/pkg/route"
)

func main() {

	router := route.NewRouter()
	router.Use(cheeseMiddleware)
	router.Use(beefMiddleware)
	router.Get("hello", helloHandler)
	res := &route.Message{}
	req := &route.Message{
		Identification: "hello",
		Method:         "get",
		Content:        "Gopher",
	}
	router.Run(res, req)
	fmt.Println(res.Content)
}

func helloHandler(res, req *route.Message) {
	fmt.Println("This is core")
	res.Content += req.Content
}

func cheeseMiddleware(next route.Handler) route.Handler {
	return func(res, req *route.Message) {
		res.Content += "cheese "
		fmt.Println("This is cheese")
		next(res, req)
		fmt.Println("This is cheese")
	}
}
func beefMiddleware(next route.Handler) route.Handler {
	return func(res, req *route.Message) {
		res.Content += "beef "
		fmt.Println("This is beef")
		next(res, req)
		fmt.Println("This is beef")
	}
}
