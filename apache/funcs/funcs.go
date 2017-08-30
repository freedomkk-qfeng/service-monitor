package funcs

import (
	"github.com/freedomkk-qfeng/service-monitor/apache/g"
	"github.com/open-falcon/common/model"
)

type FuncsAndInterval struct {
	Fs       []func() []*model.MetricValue
	Interval int
}

var Mappers []FuncsAndInterval

func BuildMappers() {
	interval := g.Config().Transfer.Interval
	Mappers = []FuncsAndInterval{
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				ApacheMetrics,
			},
			Interval: interval,
		},
	}
}
