package main

import (
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Stats struct {
	memStats    *runtime.MemStats
	PollCount   int
	RandomValue float64
	mu          sync.Mutex
}

func (s *Stats) Update() {
	s.mu.Lock()

	runtime.ReadMemStats(s.memStats)
	s.PollCount++
	s.RandomValue = randomFloat64()

	s.mu.Unlock()
}

func (s *Stats) initDoctor() {
	reportInterval := 10
	pollInterval := 2

	tickerSendServer := time.NewTicker(time.Duration(reportInterval) * time.Second)
	defer tickerSendServer.Stop()

	tickerUpdateMetric := time.NewTicker(time.Duration(pollInterval) * time.Second)
	defer tickerUpdateMetric.Stop()

	for {
		select {
		case <-tickerSendServer.C:
			pickStats()
		case <-tickerUpdateMetric.C:
			s.Update()
		}
	}
}

func NewStats() *Stats {
	return &Stats{
		memStats:    &runtime.MemStats{},
		PollCount:   0,
		RandomValue: randomFloat64(),
	}
}

func randomFloat64() float64 {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return 100 * random.Float64()
}

var statistic *Stats
var check map[string]string

func init() {
	statistic = NewStats()

	// edit next
	check = map[string]string{
		"Alloc":         "uint64",
		"BuckHashSys":   "uint64",
		"Frees":         "uint64",
		"GCCPUFraction": "float64",
		"GCSys":         "uint64",
		"HeapAlloc":     "uint64",
		"HeapIdle":      "uint64",
		"HeapInuse":     "uint64",
		"HeapObjects":   "uint64",
		"HeapReleased":  "uint64",
		"HeapSys":       "uint64",
		"LastGC":        "uint64",
		"Lookups":       "uint64",
		"MCacheInuse":   "uint64",
		"MCacheSys":     "uint64",
		"MSpanInuse":    "uint64",
		"MSpanSys":      "uint64",
		"Mallocs":       "uint64",
		"NextGC":        "uint64",
		"NumForcedGC":   "uint32",
		"NumGC":         "uint32",
		"OtherSys":      "uint64",
		"PauseTotalNs":  "uint64",
		"StackInuse":    "uint64",
		"StackSys":      "uint64",
		"Sys":           "uint64",
		"TotalAlloc":    "uint64",
	}
}

func main() {
	agent := NewStats()
	agent.initDoctor()
}

func pickStats() {

	statistic.mu.Lock()

	pushServer("counter", "PollCount", strconv.Itoa(statistic.PollCount))
	pushServer("gauge", "RandomValue", strconv.FormatFloat(statistic.RandomValue, 'f', -1, 64))

	val := reflect.ValueOf(*statistic.memStats)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		if _, ok := check[fieldName]; ok {

			switch t := field.Type().Kind(); t {
			case reflect.Uint64:
				pushServer("gauge", fieldName, strconv.FormatUint(field.Uint(), 10))
			case reflect.Float64:
				pushServer("gauge", fieldName, strconv.FormatFloat(field.Float(), 'f', -1, 64))
			case reflect.Uint32:
				pushServer("gauge", fieldName, strconv.FormatUint(field.Uint(), 10))
			}
		}
	}

	statistic.mu.Unlock()
}

func pushServer(typeMetric, nameMetric, valueMetric string) {
	url := "http://localhost:8080/update/" + typeMetric + "/" + nameMetric + "/" + valueMetric

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "text/plain")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln(resp.StatusCode, req)
	}
}
