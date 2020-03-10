package metricRecord

import (
	"strconv"
	"strings"
	"fmt"
)

type MetricRecord struct {
	Label             string  `json: "label"`
	ValueInt          int64   `json: "valueint""`
	ValueFloat        float64 `json: "valuefloat""`
	ValueType         bool    `json: "valuetype"`
	Timestamp         string  `json: "timestamp"`
	Valid             bool    `json: "valid"`
	PreviousInt       int64   `json: "previousint""`
	PreviousFloat     float64 `json: "previousfloat""`
	PreviousTimestamp string  `json: "previoustimestamp"`
}

func (m *MetricRecord) Derivative() float64 {
	if m.Valid {
		ts2, err := strconv.ParseFloat(m.PreviousTimestamp, 64)

		if err != nil {
			return 0
		}
		ts1, err := strconv.ParseFloat(m.Timestamp, 64)

		if err != nil {
			return 0
		}

		denominator := ts2 - ts1

		if denominator != 0 {
			var d float64

			if m.ValueType {
				d = (m.PreviousFloat - m.ValueFloat) / denominator
			} else {
				d = (float64(m.PreviousInt) - float64(m.ValueInt)) / denominator
			}

			return d
		}
	}

	return 0
}

func (m *MetricRecord) RecordTimestamp(ts string) {
	m.PreviousTimestamp = m.Timestamp
	m.Timestamp = ts
}

func (m *MetricRecord) SetValue(value string) {
	decimal := strings.Index(value, ".")

	if decimal > 0 {
		m.ValueType = true
		n, err := strconv.ParseFloat(value, 64)

		if err != nil {
			m.Valid = false
			return
		}
		m.PreviousFloat = m.ValueFloat
		m.ValueFloat = n
	} else {
		m.ValueType = false
		n, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			m.Valid = false
			return
		}

		m.PreviousInt = m.ValueInt
		m.ValueInt = n
	}

	m.Valid = true
}
