package main

import (
	"fc-golang/router"
	"fmt"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"runtime"
)

var r router.Router

func init() {
	r = router.Router{}
}

func main() {
	PrintMemUsage()
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/api/v1/compiler/golang/run", r.Exec)

	app.Listen(4000)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
