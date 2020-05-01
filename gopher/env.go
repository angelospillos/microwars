package main

import (
	"log"
	"os"
)

func getEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

func requireEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("environment variable '%s' not set", key)
	}
	return val
}
