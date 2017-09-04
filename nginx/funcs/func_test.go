package funcs

import (
	"testing"
)

const (
	nginx_url = "http://127.0.0.1/status"
	pid       = "/var/run/nginx.pid"
)

func Test_nginx(t *testing.T) {
	if text, code, err := httpGet(nginx_url); err != nil {
		t.Error(err)
	} else {
		//t.Log("text :", text)
		t.Log("code:", code)
		status, _ := nginx_status(text)
		t.Log("status:", status)
		version, err := nginx_version()
		t.Log("version:", version)
		t.Error(err)
		uptime, err := pid_uptime(pid)
		t.Log("uptime:", uptime)
		t.Error(err)
	}
}
