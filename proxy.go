package curl

import (
	"golang.org/x/net/proxy"
	"net"
)

type Socks5Proxy struct {
	addr string
	auth *proxy.Auth
}

func (this *Request) SetSocks5Proxy(addr string) *Request {
	if this.Socks5 == nil {
		this.Socks5 = &Socks5Proxy{}
	}
	this.Socks5.addr = addr
	return this
}

func (this *Request) SetSocksProxyAuth(user, password string) *Request {
	if this.Socks5 == nil {
		this.Socks5 = &Socks5Proxy{}
	}
	this.Socks5.auth = &proxy.Auth{User: user, Password: password}
	return this
}

// 基于forward创建Socks5的Dial
func (this *Request) newSocks5ProxyDial(forward func(network, addr string) (net.Conn, error)) (func(network, addr string) (c net.Conn, err error), error) {
	var dialer proxy.Dialer
	dialer.Dial = forward
	dialer, err := proxy.SOCKS5("tcp", this.Socks5.addr,
		this.Socks5.auth,
		dialer,
	)
	if err != nil {
		return nil, err
	}
	return dialer.Dial, nil
}
