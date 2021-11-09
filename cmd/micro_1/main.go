package main

import "github.com/mr-chelyshkin/micro/internal/http_route"

func main() {
	httRouter := http_route.NewHttpRouter()
	httRouter.Start()
}
