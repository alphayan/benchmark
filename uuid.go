package benchmark

import (
	"github.com/bwmarrin/snowflake"
	uuid2 "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
)

// go.uuid.v1
func uuidv1() string {
	return uuid.NewV1().String()
}

// go.uuid.v4
func uuidv4() string {
	return uuid.NewV4().String()
}

// snowflake
func uuid3() string {
	// Create a new Node with a Node number of 1
	node, _ := snowflake.NewNode(0)
	// Generate a snowflake ID.
	return node.Generate().String()
}

// google.uuid
func guuid() string {
	return uuid2.New().String()
}
