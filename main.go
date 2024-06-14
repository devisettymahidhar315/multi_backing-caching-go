package main

import "app/api"

func main() {

	// Initialize the Gin router and routes
	r := api.InitializeRoutes()

	// Start the server
	r.Run()

}
