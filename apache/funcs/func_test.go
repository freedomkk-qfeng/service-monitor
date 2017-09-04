package funcs

import (
	"strings"
	"testing"
)

const (
	apache_url = "https://www.apache.org/server-status?auto"
)

func Test_apache(t *testing.T) {
	apache_url := strings.Split(apache_url, "?")[0]
	url := apache_url + "?auto"
	if text, code, err := httpGet(url); err != nil {
		t.Error(err)
	} else {
		//t.Log("text :", text)
		t.Log("code:", code)
		status, _ := apache_status(text)
		t.Log("status:", status)
		version, err := apache_version(apache_url)
		t.Log("version:", version)
		t.Error(err)
	}
}
