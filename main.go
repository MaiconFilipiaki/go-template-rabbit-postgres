package main

import (
	"golangNetHttp/external/db/config"
	"golangNetHttp/external/rabbit/mount"
	"golangNetHttp/http"
	"golangNetHttp/task"
)

func main() {

	mountRabbit.Mount()
	db := config.CreateConnection()
	task.Run(db)
	s := http.NewServer()
	s.Run()
}
