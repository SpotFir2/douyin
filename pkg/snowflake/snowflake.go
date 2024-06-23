package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
)

/*
Init 初始化雪花算法参数
*/
func Init(startTime string, id int64) (err error) {
	var nt time.Time
	nt, err = time.Parse("2006-01-02 15:04:05", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = nt.UnixMicro()
	node, err = snowflake.NewNode(id)
	return
}

/*
GenerateId 使用雪花算法生成Id
*/
func GenerateId() int64 {
	return node.Generate().Int64()
}
