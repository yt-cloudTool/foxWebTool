package global_id

import (
	"github.com/bwmarrin/snowflake"
	"fmt"
)

var nodeVar *snowflake.Node

func init () {
	fmt.Println("global_id init =>")
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Sprintf("global_id init err => %s", err)
	}
	
	nodeVar = node
}

// 生成id
func Generate () string {
	return nodeVar.Generate().String()
}