package local

import (
	"os"
	"strings"
	"widesocks/common/slog"
)

var servers map[string]interface{}
var ipAddr []string

func init() {
	var addrs string
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		if args[i] == "-i" {
			if (i + 1) < len(args) {
				i++
				addrs = args[i]
			} else {
				slog.Emergency("请输入正确的服务器IP")
			}
		} else if args[i] == "-v" {
			slog.SetLogLevel(slog.LOG_DEBUG)
		}
	}

	if addrs == "" {
		addrs = "127.0.0.1:8090"
	}

	servers = map[string]interface{}{
		"addr": addrs,
	}

	ipAddr = strings.Split(servers["addr"].(string), ",")
}
