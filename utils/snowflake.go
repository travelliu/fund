package utils

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

// snow flake id
var node *snowflake.Node

func init() {
	// Create a new Node with a Node number of 1
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatalln("init snowflake node with err", err)
	}
}

// GenerateID generate a Twitter snowflake ID
func GenerateID() int64 {
	// Generate a snowflake ID.
	return node.Generate().Int64()
}
