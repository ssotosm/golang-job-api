package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"witpgh-jobapi-go/app/route"
	"witpgh-jobapi-go/app/shared/database"

	"github.com/joho/godotenv"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	configureEnvironments()
	database.ConnectWITJobBoard()

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	fmt.Println("Application is running ... in port : ", port)

	http.ListenAndServe(":"+port, route.LoadRoutes())

}

func configureEnvironments() {
	os.Clearenv()

	err := godotenv.Load("doc.env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}
