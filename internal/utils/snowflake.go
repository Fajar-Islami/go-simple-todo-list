package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitSnowflake() {
	snowNode, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println("err snowflake", err)
		return
	}

	node = snowNode
}

func IDGenerator() int64 {
	return node.Generate().Int64()
}
