// Package datanode is to define and declare all functionalities of the namenode server
package datanode

import (
	"fmt"
	"os"
)

func InitDatanode() {
	port := os.Getenv("DATANODE_PORT")
	fmt.Printf("Initializing datanode on port %s", port)
}
