package main

import (
	"video-conf/core/cmd"
)

// @securityDefinitions.apikey Bearer
// @in 						   header
// @name 					   Authorization
func main() {
	cmd.Execute()
}
