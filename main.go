package main

import "github.com/mariojuzar/go-user-auth/cmd"

// @title User Auth API
// @version 1.0.0
// @description User Auth API
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
