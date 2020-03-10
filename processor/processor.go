package processor

import (
	"conscientia/internalMetrics"
	"fmt"
	"regexp"
	"conscientia/metricCollection"
	"conscientia/metricRecord"
)


func Process(mCh chan []byte, mc metricCollection.MetricCollection,
	intM *internalMetrics.InternalMetrics) {

	const metricRegex = "^([0-9a-zA-Z-._:]+)\\s+([0-9.-]+)\\s+([0-9]+)"

	mRegex, err := regexp.Compile(metricRegex)
	if err != nil {
		fmt.Println("Can not compile regular expression.")
		return
	}

	for {
		m := <-mCh

		result := mRegex.FindSubmatch(m)
		if result != nil {
			if met, ok := mc.Get(string(result[1])); ok {
				met.SetValue(string(result[2]))
				met.RecordTimestamp(string(result[3]))
				mc.Put(met)
			} else {
				met := metricRecord.MetricRecord{Label: string(result[1]), Timestamp: string(result[3])}
				met.SetValue(string(result[2]))
				mc.Put(met)
			}

			intM.IncrementGlobalScalar(internalMetrics.MetricProcessed)
		} else {
			intM.IncrementGlobalScalar(internalMetrics.ProcessRegexMiss)
		}
	}
}
