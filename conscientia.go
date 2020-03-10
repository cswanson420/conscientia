package main

import (
	"conscientia/collector"
	"conscientia/processor"
	"time"
	"conscientia/internalMetrics"
	"conscientia/metricCollection"
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
)

func index() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, nil)
	})
}

func countKeys(mc metricCollection.MetricCollection, mCh chan []byte,
	intM *internalMetrics.InternalMetrics) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		keys := mc.Size()

		fmt.Fprintf(w, "Key count: %d<br>", keys)
		fmt.Fprintf(w, "Channel Size: %d<br>", len(mCh))
		fmt.Fprintf(w, "Metrics Received: %d<br>", intM.GetGlobalScalar(internalMetrics.MetricReceived))
		fmt.Fprintf(w, "Metrics Processed: %d<br>", intM.GetGlobalScalar(internalMetrics.MetricProcessed))
		fmt.Fprintf(w, "Processor Regex Miss: %d<br>", intM.GetGlobalScalar(internalMetrics.ProcessRegexMiss))
		fmt.Fprint(w, "<br><a href=\"/\">Home</a>")
	})
}

func getKey(mc metricCollection.MetricCollection) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params, ok := r.URL.Query()["m"]

		if !ok || len(params[0]) < 1 {
			fmt.Fprintf(w, "[]")
			return
		}

		value, ok := mc.Get(params[0])

		if !ok {
			fmt.Fprintf(w, "[]")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		enc.Encode(value)
	})
}


func registerHandlers(mc metricCollection.MetricCollection, mCh chan []byte,
	intM *internalMetrics.InternalMetrics) {
	http.HandleFunc("/", index())
	http.HandleFunc("/metrics", countKeys(mc, mCh, intM))
	http.HandleFunc("/g", getKey(mc))
}

func webserver() {
	for {
		http.ListenAndServe(":3005", nil)
	}
}


func main() {
	intM := internalMetrics.InternalMetrics{}
	intM.Init()
	mCh := make(chan []byte)

	mc := metricCollection.MetricCollection{}
	mc.Init()

	registerHandlers(mc, mCh, &intM)
	go webserver()
	go collector.Collect(mCh, &intM)

	for i := 0; i < 20; i++ {
		go processor.Process(mCh, mc, &intM)
	}

	for {
		time.Sleep(time.Second * 60)
	}
}
