package profile

type AppProfile struct {
	Name  string
	App   string
	Path  string
	Host  string
	Port  string
	Proxy *ProxyConfig
}

type ProxyConfig struct {
	Host string
	Port string
}
