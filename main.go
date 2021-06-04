package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/akamensky/argparse"
	"github.com/gin-gonic/gin"
)

const ENVFolder = "FOLDER"
const ENVPort = "PORT"
const DefaultPort = 8080

func parseEnvironment() (int, string) {
	env := os.Getenv(ENVPort)
	port := 0
	folder := ""
	if len(env) > 0 {
		port_, err := strconv.Atoi(env)
		if err == nil {
			port = port_
		}
	}
	env = os.Getenv(ENVFolder)
	if len(env) > 0 {
		_, err := os.Stat(env)
		if os.IsNotExist(err) {
			fmt.Printf("ERROR: Folder not found: %s\n", env)
		} else {
			folder = env
		}
	}

	return port, folder
}

func main() {
	parser := argparse.NewParser("http-static", "HTTP web server for static files")
	envPort, envFolder := parseEnvironment()
	defaultPort := DefaultPort
	if envPort > 0 {
		defaultPort = envPort
	}

	var port *int = parser.Int("p", "port", &argparse.Options{Required: envPort < 1, Help: fmt.Sprintf("Port for listening (%s)", ENVPort), Default: defaultPort})
	var folder = parser.String("f", "folder", &argparse.Options{Required: len(envFolder) == 0, Help: fmt.Sprintf("Folder to serve (%s)", ENVFolder), Default: envFolder})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	router := gin.Default()
	router.StaticFS("/", http.Dir(*folder))

	router.Run(":" + strconv.Itoa(*port))
}
