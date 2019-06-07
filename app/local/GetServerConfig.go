package local

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"widesocks/config"
)

func GetServerConfig(hostport string) error {
	ip, port, err := net.SplitHostPort(hostport)
	if err != nil {
		return err
	}

	IP := net.ParseIP(ip)
	if IP == nil {
		return errors.New("请输入正确的服务器IP")
	}

	portInt, err := strconv.Atoi(port)
	if err != nil || portInt < 1 || portInt > 65535 {
		return errors.New("请输入正确的服务器端口号")
	}

	// 获取服务器配置
	httpClient := &http.Client{}
	response, err := httpClient.Get(fmt.Sprintf("http://%s:%d/getssconfig", ip, portInt))
	if err != nil {
		return err
	} else {
		defer response.Body.Close()
	}

	var conf []byte
	conf, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var serverConfObj config.ServerConfig
	err = json.Unmarshal(conf, &serverConfObj)
	if err != nil {
		return err
	}

	var userver config.UpstreamServer
	for k, v := range serverConfObj.PortPassword {
		userver = config.UpstreamServer{
			Server:     IP,
			ServerPort: k,
			Password:   v,
			Weight:     1,
			Method:     serverConfObj.Method,
		}
		config.Servers = append(config.Servers, userver)
	}

	return nil
}
