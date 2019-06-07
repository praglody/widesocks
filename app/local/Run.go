package local

import (
	"encoding/json"
	"fmt"
	"widesocks/common/slog"
	"widesocks/config"
)

func Run() {
	// 获取服务器配置文件
	for _, v := range ipAddr {
		err := GetServerConfig(v)
		if err != nil {
			slog.Warningf("Error %s with IP %s", err.Error(), v)
		}
	}

	if config.Servers == nil {
		slog.Emergency("没有可用的上游代理服务器")
	}

	s, _ := json.MarshalIndent(config.Servers, "", "  ")
	fmt.Println(string(s))

	// 监听本地 14213 端口，作为 socks5 代理端口
}
