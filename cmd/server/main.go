package main

import httpserver "Gearjot/pkg/api/http_server"

func main() {
	r := httpserver.CreateRouter()
	r.Run()
}
