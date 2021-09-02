package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	AddSuccessMultiCreateHeroesCounter(val float64)
	IncSuccessCreateHeroCounter()
	IncSuccessListHeroesCounter()
	IncSuccessDescribeHeroCounter()
	IncSuccessRemoveHeroCounter()
	IncSuccessUpdateHeroCounter()
}

type metrics struct {
	successMultiCreateHeroCounter prometheus.Counter
	successListHeroesCounter      prometheus.Counter
	successCreateHeroCounter      prometheus.Counter
	successDescribeHeroCounter    prometheus.Counter
	successRemoveHeroCounter      prometheus.Counter
	successUpdateHeroCounter      prometheus.Counter
}

func NewApiMetrics(namespace, subsystem string) Metrics {
	return &metrics{
		successMultiCreateHeroCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_multi_create_heroes_request_count",
			Help:      "The total number of success multi created heroes",
		}),
		successCreateHeroCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_create_hero_request_count",
			Help:      "The total number of success created heroes",
		}),
		successListHeroesCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_list_heroes_request_count",
			Help:      "The total number of success list heroes",
		}),
		successDescribeHeroCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_describe_hero_request_count",
			Help:      "The total number of success describe heroes",
		}),
		successRemoveHeroCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_remove_hero_request_count",
			Help:      "The total number of success removed heroes",
		}),
		successUpdateHeroCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_update_hero_request_count",
			Help:      "The total number of success removed heroes",
		}),
	}
}

func (m *metrics) AddSuccessMultiCreateHeroesCounter(val float64) {
	m.successMultiCreateHeroCounter.Add(val)
}
func (m *metrics) IncSuccessCreateHeroCounter() {
	m.successCreateHeroCounter.Inc()
}
func (m *metrics) IncSuccessListHeroesCounter() {
	m.successListHeroesCounter.Inc()
}
func (m *metrics) IncSuccessDescribeHeroCounter() {
	m.successDescribeHeroCounter.Inc()
}
func (m *metrics) IncSuccessRemoveHeroCounter() {
	m.successRemoveHeroCounter.Inc()
}
func (m *metrics) IncSuccessUpdateHeroCounter() {
	m.successUpdateHeroCounter.Inc()
}
