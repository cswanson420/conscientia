package internalMetrics

import "sync"

const (
	MetricReceived = iota
	MetricProcessed = iota
	ProcessRegexMiss = iota)

type InternalMetrics struct {
	globalScalarMap map[int] int
	globalMutex *sync.Mutex
}

func (m *InternalMetrics) Init() {
	m.globalScalarMap = make(map[int] int)
	m.globalMutex = &sync.Mutex{}
}

func (m *InternalMetrics) StoreGlobalScalar(key int, value int) {
	m.globalMutex.Lock()
	defer m.globalMutex.Unlock()

	m.globalScalarMap[key] = value
}

func (m *InternalMetrics) IncrementGlobalScalar(key int) {
	m.globalMutex.Lock()
	defer m.globalMutex.Unlock()

	m.globalScalarMap[key]++
}

func (m *InternalMetrics) GetGlobalScalar(key int) int {
	m.globalMutex.Lock()
	defer m.globalMutex.Unlock()

	var v int

	if  _, ok := m.globalScalarMap[key]; !ok {
		v = 0
	} else {
		v = m.globalScalarMap[key]
	}

	return v
}

