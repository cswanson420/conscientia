package metricCollection

import (
	"sync"
	"conscientia/metricRecord"
)

type MetricCollection struct {
	m map[string]metricRecord.MetricRecord
	mutex *sync.Mutex
}

func (mc *MetricCollection) Init() {
	mc.m = make(map[string]metricRecord.MetricRecord)
	mc.mutex = &sync.Mutex{}
}

func (mc *MetricCollection) Get(label string) (metricRecord.MetricRecord, bool) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	metric, ok := mc.m[label]

	if !ok {
		return metricRecord.MetricRecord{}, false
	}

	return metric, true
}

func (mc *MetricCollection) Put(record metricRecord.MetricRecord) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mc.m[record.Label] = record
}

func (mc *MetricCollection) Size() int64 {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	return int64(len(mc.m))
}