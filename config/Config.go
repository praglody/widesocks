package config

import (
	"net"
)

type ServerConfig struct {
	PortPassword map[uint16]string `json:"port_password"`
	Method       string            `json:"method"`
}

type UpstreamServer struct {
	Server     net.IP
	ServerPort uint16
	Password   string
	Weight     int
	Method     string
}

var Servers []UpstreamServer
