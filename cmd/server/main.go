package main

import httpserver "Gearjot/pkg/api/http_server"

func main() {
	r := httpserver.CreateRouter()
	r.Run()

	// For future:

	// parsing config.

	// adding logger.

	// connecting to SomeDB.

	// creating an error channel and signals for shutdown.

	// creating a metric recorder.

	// starting the REST server and serving connections.

	// serving metrics.

	// blocking select to intercept interrupt signals and unexpected errors.

}
