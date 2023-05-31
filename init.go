package main

import (
	"log"
	"os"
)

func Init() {
	envs := []string{"GITURL"}

	for _, v := range envs {
		if os.Getenv(v) == "" {
			log.Fatalf("%s not set but required", v)
		}
	}

	optEnvs := map[string]string{
		"FETCH_INTERVAL":    "300",
		"PORT":              "5555",
		"GIT_REPO_LOCATION": "./repo",
	}

	for k, v := range optEnvs {
		if os.Getenv(k) == "" {
			if err := os.Setenv(k, v); err != nil {
				log.Fatal(err)
			}
		}
	}
}
