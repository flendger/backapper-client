package clientfactory

import (
	"backapper-client/profile"
	"context"
	"h12.io/socks"
	"net"
	"net/http"
)

func CreateClient(p *profile.AppProfile) *http.Client {

	proxyHost := p.Proxy.Host
	proxyPort := p.Proxy.Port

	if proxyHost != "" && proxyPort != "" {
		dialSocksProxy := socks.Dial("socks5://" + proxyHost + ":" + proxyPort + "?timeout=72000s")

		tr := &http.Transport{DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialSocksProxy(network, addr)
		}}

		return &http.Client{Transport: tr}
	} else {
		return &http.Client{}
	}
}
