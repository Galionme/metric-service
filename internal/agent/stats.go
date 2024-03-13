package agent

import (
	"github.com/Galionme/metric-service.git/internal/agent/picker"
	"github.com/Galionme/metric-service.git/internal/util"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Stats struct {
	memStats       *runtime.MemStats
	PollCount      int
	RandomValue    float64
	minRandomValue float64
	maxRandomValue float64
	mu             sync.Mutex
	send           func(a, b, c, address string)
	filter         *picker.Picker
	address        string
}

func NewStats(minRandomValue, maxRandomValue float64, send func(a, b, c, address string), address string) *Stats {
	return &Stats{
		memStats:       &runtime.MemStats{},
		PollCount:      0,
		minRandomValue: minRandomValue,
		maxRandomValue: maxRandomValue,
		RandomValue:    util.RandomFloat64(minRandomValue, maxRandomValue),
		send:           send,
		filter:         picker.NewPicker(),
		address:        address,
	}
}

func (s *Stats) update() {
	s.mu.Lock()

	runtime.ReadMemStats(s.memStats)
	s.PollCount++
	s.RandomValue = util.RandomFloat64(s.minRandomValue, s.maxRandomValue)

	s.mu.Unlock()
}

func (s *Stats) InitDoctor(reportInterval, pollInterval int) {

	tickerSendServer := time.NewTicker(time.Duration(reportInterval) * time.Second)
	defer tickerSendServer.Stop()

	tickerUpdateMetric := time.NewTicker(time.Duration(pollInterval) * time.Second)
	defer tickerUpdateMetric.Stop()

	for {
		select {
		case <-tickerSendServer.C:
			s.push()
		case <-tickerUpdateMetric.C:
			s.update()
		}
	}
}

func (s *Stats) push() {
	s.mu.Lock()

	s.send("counter", "PollCount", strconv.Itoa(s.PollCount), s.address)

	strRandom, _ := util.GaugeToString(s.RandomValue)
	s.send("gauge", "RandomValue", strRandom, s.address)

	val := reflect.ValueOf(*s.memStats)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		if str, err := s.filter.GetString(fieldName, field); err == nil {
			s.send("gauge", fieldName, str, s.address)
		}

	}

	s.mu.Unlock()
}
