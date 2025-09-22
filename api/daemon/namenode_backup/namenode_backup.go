// Package namenodebackup serves all functionality for backup namenode server
package namenodebackup

import (
	"fmt"
	"os"
)

func InitNamenodeBackup() {
	port := os.Getenv("NAMENODE_BACKUP_PORT")
	fmt.Printf("Initializing namenode_backup on %s", port)
}
