package log

import (
	"nodeWatch/def"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func NewLogger() *glog.Logger {
	logger := glog.New()
	logger.SetConfigWithMap(g.Map{
		//"path": "D:/golang/mypath/src/goProject/nodesWatch",
		"path":     def.CurPath,
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
	})
	return logger
}
