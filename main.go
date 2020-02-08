package main

import (
	"os"

	"github.com/bcelenza/carrier/httpsrv"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Setup logging
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	port := getEnvOrDefault("HTTP_PORT", "8080")
	cert := getEnvOrDefault("TLS_CERT_FILE", "")
	key := getEnvOrDefault("TLS_KEY_FILE", "")
	log.Info("Starting Carrier on port ", port)

	httpserver := httpsrv.New(port, cert, key)
	log.Fatalln(httpserver.Start())
}

func getEnvOrDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
