package main

import (
	"log"
	"os"

	namenode "github.com/Triyaambak/dfs/daemon/namenode"
	namenode_backup "github.com/Triyaambak/dfs/daemon/namenode_backup"
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
	case "namenode_backup":
		namenode_backup.InitNamenodeBackup()
	}
}
