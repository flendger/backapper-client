package profileholder

import (
	"log"
	"testing"
)

const configFile = "backapper-client-test.json"

func TestRead(t *testing.T) {
	logger := *log.Default()
	holder := Read(configFile, &logger)

	if len(holder.profiles) != 1 {
		t.Errorf("Profiles size must be 1")
	}

	profileName := "dev"
	p, err := holder.GetProfile(profileName)
	if err != nil {
		t.Errorf("Profile not found: %s", profileName)
		return
	}

	appName := "test-app"
	if p.App != appName {
		t.Errorf("App name must be %s", appName)
	}

	path := "./test.txt"
	if p.Path != path {
		t.Errorf("Path maust be %s", path)
	}

	host := "localhost"
	if p.Host != host {
		t.Errorf("Host maust be %s", host)
	}

	port := "8787"
	if p.Port != port {
		t.Errorf("Port maust be %s", port)
	}

	proxy := p.Proxy

	proxyHost := "localhost"
	if proxy.Host != proxyHost {
		t.Errorf("Proxy host maust be %s", proxyHost)
	}

	proxyPort := "8989"
	if proxy.Port != proxyPort {
		t.Errorf("Proxy port maust be %s", proxyPort)
	}
}
