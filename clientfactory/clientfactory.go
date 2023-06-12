package clientfactory

import (
	"backapper-client/profile"
	"h12.io/socks"
	"net/http"
)

func CreateClient(p *profile.AppProfile) *http.Client {

	proxyHost := p.Proxy.Host
	proxyPort := p.Proxy.Port

	if proxyHost != "" && proxyPort != "" {
		dialSocksProxy := socks.Dial("socks5://" + proxyHost + ":" + proxyPort + "?timeout=72000s")
		tr := &http.Transport{Dial: dialSocksProxy}
		return &http.Client{Transport: tr}
	} else {
		return &http.Client{}
	}
}
