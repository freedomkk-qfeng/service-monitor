package funcs

import (
	"testing"
)

const (
	Addr     = "192.168.11.136:22"
	Username = ""
	Password = ""
	Authdb   = ""
)

func Test_mongo_stat(t *testing.T) {
	serverStatus, err := mongo_serverStatus(Addr, Authdb, Username, Password)
	t.Log(serverStatus)
	t.Error(err)
	ver := mongo_version(serverStatus)
	t.Log(ver)
	CounterMetrics, GaugeMetrics := mongo_Metrics(serverStatus)
	t.Log(CounterMetrics)
	t.Log(GaugeMetrics)
}
