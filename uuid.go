package benchmark

import (
	"github.com/bwmarrin/snowflake"
	"github.com/satori/go.uuid"
)

// go.uuid.v1
func uuidv1() string {
	return uuid.Must(uuid.NewV1()).String()
}

// go.uuid.v4
func uuidv4() string {
	return uuid.Must(uuid.NewV4()).String()
}

// snowflake
func uuid3() string {
	// Create a new Node with a Node number of 1
	node, _ := snowflake.NewNode(1)
	// Generate a snowflake ID.
	id := node.Generate()
	return id.String()
}
