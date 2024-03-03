package integtest

import (
	"log"
	"net"
	"net/url"
	"os"
	"strconv"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func GetHostPort(res *dockertest.Resource, portID string) (host string, port string) {
	host = "127.0.0.1"

	if h := os.Getenv("DOCKER_HOST"); h != "" {
		u, err := url.Parse(h)
		if err != nil {
			log.Fatalf("parse docker host: %s", err)
		}

		host = u.Hostname()
	}

	addr, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("lookup docker host (%s): %s", host, err)
	}
	if len(addr) > 0 {
		host = addr[0].String()
	}

	if res == nil {
		port = docker.Port(portID).Port()
		return
	}

	port = res.GetPort(portID)
	return
}

func GetHostPortInt(res *dockertest.Resource, portID string) (host string, port int) {
	var rawPort string
	host, rawPort = GetHostPort(res, portID)

	var err error
	if port, err = strconv.Atoi(rawPort); err != nil {
		log.Fatalf("convert port %s to int: %s", rawPort, err)
		return
	}

	return
}
