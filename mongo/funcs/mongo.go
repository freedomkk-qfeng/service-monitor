package funcs

import (
	"log"
	"strings"

	"github.com/freedomkk-qfeng/service-monitor/mongo/g"
	"github.com/open-falcon/common/model"
)

func MongoMetrics() (L []*model.MetricValue) {
	if !g.Config().Mongo.Enabled {
		log.Println("Mongo Monitor is disabled")
		return
	}
	Addr := g.Config().Mongo.Addr
	Username := g.Config().Mongo.Username
	Password := g.Config().Mongo.Password
	Authdb := g.Config().Mongo.Authdb
	Port := strings.Split(Addr, ":")[1]
	serverStatus, err := mongo_serverStatus(Addr, Authdb, Username, Password)
	if err != nil {
		L = append(L, GaugeValue("Mongo.Alive", -1, "Port="+Port))
		log.Println(err)
		return
	}

	version := mongo_version(serverStatus)
	if err == nil {
		log.Println("Mongo version is: ", version)
	} else {
		log.Println(err)
	}

	CounterMetrics, GaugeMetrics := mongo_Metrics(serverStatus)
	if connections_current, ok := GaugeMetrics["connections_current"]; ok {
		if connections_available, ok := GaugeMetrics["connections_available"]; ok {
			if connections_available != 0 {
				connections_used_percent := 100 * float64(connections_current) / float64(connections_available)
				L = append(L, GaugeValue("Mongo.connections_used_percent", connections_used_percent, "Port="+Port))
			}
		}
	}
	for metric, value := range GaugeMetrics {
		L = append(L, GaugeValue("Mongo."+metric, value, "Port="+Port))
	}
	for metric, value := range CounterMetrics {
		L = append(L, CounterValue("Mongo."+metric, value, "Port="+Port))
	}
	return
}
