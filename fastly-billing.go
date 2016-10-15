package main

import (
	"flag"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

// FastlyBillingPlugin mackerel plugin for Fastly Billing
type FastlyBillingPlugin struct {
	apiKey string
	prefix string
}

// FetchMetrics interface for PluginWithPrefix
func (p FastlyBillingPlugin) FetchMetrics() (map[string]interface{}, error) {
	return map[string]interface{}{
		"totalCost": 100.0,
	}, nil
}

// GraphDefinition interface for PluginWithPrefix
func (p FastlyBillingPlugin) GraphDefinition() map[string](mp.Graphs) {
	labelPrefix := strings.Title(p.prefix)

	// metric value structure
	var graphdef = map[string](mp.Graphs){
		"billing": {
			Label: (labelPrefix + " Billing"),
			Unit:  "float",
			Metrics: [](mp.Metrics){
				{Name: "totalCost", Label: "Total Cost", Type: "float64"},
			},
		},
	}

	return graphdef
}

// MetricKeyPrefix interface for PluginWithPrefix
func (p FastlyBillingPlugin) MetricKeyPrefix() string {
	if p.prefix == "" {
		p.prefix = "fastly"
	}
	return p.prefix
}

func main() {
	optAPIKey := flag.String("api-key", "", "Fastly API Key")
	optMetricKeyPrefix := flag.String("metric-key-prefix", "fastly", "Metric Key Prefix")
	flag.Parse()

	var fastlyBilling FastlyBillingPlugin

	fastlyBilling.apiKey = *optAPIKey
	fastlyBilling.prefix = *optMetricKeyPrefix

	helper := mp.NewMackerelPlugin(fastlyBilling)
	helper.Run()
}
