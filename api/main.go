package main

import (
	"fmt"
	"log"
	"os"

	namenode "github.com/Triyaambak/dfs/daemon/namenode"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Something went wrong while reading .env file :", err)
	}

	role := os.Getenv("ROLE")
	switch role {
	case "namenode":
		namenode.InitNamenode()
	}
	fmt.Println(role)
}
