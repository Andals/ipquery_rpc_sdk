package ipquery

import (
    "net"
    "fmt"
    "os"
    "runtime"
	"time"

    "git.apache.org/thrift.git/lib/go/thrift"
)

const DEFAULT_TIMEOUT = 3

type IpqueryClientWrapper struct {
    client *IpqueryClient
}

func NewIpqueryClientWrapperUnix (name string, path string, timeoutSeconds int) *IpqueryClientWrapper {
    var trans thrift.TTransport
	unixAddr, err := net.ResolveUnixAddr(name, path)
	if err != nil {
        fmt.Fprintln(os.Stderr, "error resolving address:", err)
        os.Exit(1)
	}
	trans = thrift.NewTSocketFromAddrTimeout(unixAddr, time.Duration(timeoutSeconds) * time.Second)
	return getIpqueryClientWrapper(trans)
}

func NewIpqueryClientWrapper (host string, port string, timeoutSeconds int) *IpqueryClientWrapper {
    var trans thrift.TTransport
    trans, err := thrift.NewTSocketTimeout(net.JoinHostPort(host, port), time.Duration(timeoutSeconds) * time.Second)
    if err != nil {
        fmt.Fprintln(os.Stderr, "error resolving address:", err)
        os.Exit(1)
    }   
	return getIpqueryClientWrapper(trans)
}

func (this *IpqueryClientWrapper) Ipquery(ip string, tag string) (string, error) {
    return this.client.Ipquery(ip, tag)
}

func (this *IpqueryClientWrapper) IpqueryEx(ipList []string, tag string) ([]string, error) {
    return this.client.IpqueryEx(ipList, tag)
}

func (this *IpqueryClientWrapper) IpqueryVersion() (string, error) {
    return this.client.IpqueryVersion()
}

func (this *IpqueryClientWrapper) Close() {
    this.client.Transport.Close()
}

func getIpqueryClientWrapper(trans thrift.TTransport) *IpqueryClientWrapper {
    trans = thrift.NewTFramedTransport(trans)
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    client := NewIpqueryClientFactory(trans, protocolFactory)
    if err := trans.Open(); err != nil {
        fmt.Fprintln(os.Stderr, "Error opening socket", err)
        os.Exit(1)
    }
    runtime.SetFinalizer(client, func(client *IpqueryClient) {
        client.Transport.Close()
    })
    return &IpqueryClientWrapper{client : client}
}
