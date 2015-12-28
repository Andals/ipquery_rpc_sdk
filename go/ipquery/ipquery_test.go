package ipquery

import (
    "fmt"
    "testing"
)

const HOST = "localhost"
const PORT = "9090"
const TIMEOUT = 3

const DOMAIN_SOCKET_PATH = "/tmp/ipquery_socket"
const DOMAIN_SOCKET_NAME = "unix"

func TestIpquery(t *testing.T) {
    client := NewIpqueryClientWrapper(HOST, PORT, TIMEOUT)
    result, err := client.Ipquery("223.103.15.1", "prov_city")
	if err != nil {
		t.Errorf("test Ipquery failed:%s", err)
	}
    fmt.Println(result)
}


func TestIpqueryEx(t *testing.T) {
    client := NewIpqueryClientWrapper(HOST, PORT, TIMEOUT)
    ips := []string{
        "223.103.15.1",
        "58.30.15.255",
    }
    result, err := client.IpqueryEx(ips, "prov_city")
	if err != nil {
		t.Errorf("test IpqueryEx failed:%s", err)
	}

    for _, value := range result {
        fmt.Println(value)
    }
}

func TestIpqueryVersion(t *testing.T) {
    client := NewIpqueryClientWrapper(HOST, PORT, TIMEOUT)
    result, err := client.IpqueryVersion()
	if err != nil {
		t.Errorf("test IpqueryVersion failed:%s", err)
	}
    fmt.Println(result)
}

func TestIpqueryUnix(t *testing.T) {
	client := NewIpqueryClientWrapperUnix(DOMAIN_SOCKET_NAME, DOMAIN_SOCKET_PATH, TIMEOUT)
	result, err := client.Ipquery("223.103.15.1", "prov_city")
	if err != nil {
		t.Errorf("test Ipquery failed:%s", err)
	}
    fmt.Println(result)
}

func TestIpqueryExUnix(t *testing.T) {
	client := NewIpqueryClientWrapperUnix(DOMAIN_SOCKET_NAME, DOMAIN_SOCKET_PATH, TIMEOUT)
    ips := []string{
        "223.103.15.1",
        "58.30.15.255",
    }
    result, err := client.IpqueryEx(ips, "prov_city")
	if err != nil {
		t.Errorf("test IpqueryEx failed:%s", err)
	}

    for _, value := range result {
        fmt.Println(value)
    }
}
func TestIpqueryVersionEx(t *testing.T) {
	client := NewIpqueryClientWrapperUnix(DOMAIN_SOCKET_NAME, DOMAIN_SOCKET_PATH, TIMEOUT)
    result, err := client.IpqueryVersion()
	if err != nil {
		t.Errorf("test IpqueryVersion failed:%s", err)
	}
    fmt.Println(result)
}
