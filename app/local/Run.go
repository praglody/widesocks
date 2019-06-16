package local

import (
	"encoding/json"
	"fmt"
	"syscall"
	"widesocks/common/eventloop"
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
	var fd int
	var err error
	var proxyAddr syscall.SockaddrInet4
	if fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP); err != nil {
		slog.Emergency("create socket error")
	}

	if err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		slog.Emergency("set socket opt error")
	}

	if err = syscall.SetNonblock(fd, true); err != nil {
		slog.Emergency("set socket non block error")
	}

	proxyAddr.Addr = [4]byte{0, 0, 0, 0}
	proxyAddr.Port = 14213
	if err = syscall.Bind(fd, &proxyAddr); err != nil {
		slog.Emergency("Bind local port 14213 error")
	}

	if err = syscall.Listen(fd, 512); err != nil {
		slog.Emergencyf("Listening fd %d error", fd)
	}

	if err = eventloop.EpollAdd(fd, eventloop.EPOLL_READ); err != nil {
		slog.Emergency(err)
	}

	if ev, err := eventloop.EpollWait(); err != nil {
		slog.Emergency(err)
	} else {
		var ufd int
		var uaddr syscall.Sockaddr
		for _, e := range ev {
			if e.Events&eventloop.EPOLL_ERR != 0 {
				// 处理 socket 错误，这里分为监听的端口错误事件和用户连接的错误事件
			}
			if e.Fd == int32(fd) {
				// 新连接事件
				ufd, uaddr, err = syscall.Accept(int(e.Fd))
				slog.Infof("new conn，fd: %d, addr: %v", ufd, uaddr)

				// 将新的连接加入epoll监听队列
				if err = eventloop.EpollAdd(ufd, eventloop.EPOLL_READ); err != nil {
					slog.Warningf("fd %d add epoll err", ufd)
					_ = syscall.Close(fd)
				}
			} else {
				// 用户连接的读写事件
			}
		}
	}
}
