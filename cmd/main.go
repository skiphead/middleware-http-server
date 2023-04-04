package main

import "github.com/SkipHead/middleware-http-server/internal/server"

func main() {

	srv := server.NewConfig()

	srv.Run()
}
