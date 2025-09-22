// Package namenode is to define and declare all functionalities of the namenode server
package namenode

import (
	"fmt"
	"os"
)

type Namenode struct{}

func InitNamenode() {
	port := os.Getenv("NAMENODE_PORT")
	fmt.Printf("Initializing namenode on %s", port)
}
