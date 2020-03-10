package metricCollection

import (
	"testing"
	"conscientia/metricRecord"
)

func TestMetricCollectionDefinedNaive(t *testing.T) {
	mc := MetricCollection{}

	mc.m = make(map[string]metricRecord.MetricRecord, 10)
}

func TestMetricCollectionInitNaive(t *testing.T) {
	expected := 1
	mc := MetricCollection{}

	mc.Init()

	mc.m["metric"] = metricRecord.MetricRecord{}

	actual := len(mc.m)

	if actual != expected {
		t.Errorf("Expected value: %d does not equal actual value: %d.", expected, actual)
	}
}

func TestMetricCollectionGetNaive(t *testing.T) {
	expected := int64(1)
	label := "metric"

	mc := MetricCollection{}
	mc.Init()
	mc.m[label] = metricRecord.MetricRecord{ValueInt: 1}

	record, _ := mc.Get(label)

	actual := record.ValueInt

	if actual != expected {
		t.Errorf("Expected value: %d does not equal actual value: %d.", expected, actual)
	}
}

func TestMetricCollectionGetInvalidLabel(t *testing.T) {
	expected := 0
	label := "metric"

	mc := MetricCollection{}
	mc.Init()

	_, ok := mc.Get(label)

	actual := 1

	if !ok {
		actual = 0
	}

	if actual != expected {
		t.Errorf("Expected value: %d does not equal actual value: %d.", expected, actual)
	}
}

func TestMetricCollectionSizeNaive(t *testing.T) {
	expected := int64(1)
	label := "metric"

	mc := MetricCollection{}
	mc.Init()
	mc.m[label] = metricRecord.MetricRecord{ValueInt: 1}

	actual := mc.Size()

	if actual != expected {
		t.Errorf("Expected value: %d does not equal actual value: %d.", expected, actual)
	}
}

func TestMetricCollectionPutNaive(t *testing.T) {
	expected := int64(1)
	label := "metric"

	mc := MetricCollection{}
	mc.Init()
	mc.Put(metricRecord.MetricRecord{Label: label, ValueInt: 1})

	actual := mc.Size()

	if actual != expected {
		t.Errorf("Expected value: %d does not equal actual value: %d.", expected, actual)
	}
}