package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

// FastlyBillingPlugin mackerel plugin for Fastly Billing
type FastlyBillingPlugin struct {
	apiKey string
	prefix string
}

// FetchMetrics interface for PluginWithPrefix
func (p FastlyBillingPlugin) FetchMetrics() (map[string]interface{}, error) {
	totalCost, err := thisMonthTotalCost(p.apiKey)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"totalCost": totalCost,
	}, nil
}

func thisMonthTotalCost(apiKey string) (float64, error) {
	client := &http.Client{Timeout: time.Duration(5) * time.Second}

	req, err := http.NewRequest("GET", fastlyEndPoint(), nil)
	if err != nil {
		return 0.0, err
	}

	req.Header.Add("Fastly-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0.0, errors.New(resp.Status)
	}

	return pickTotalCost(resp)
}

func fastlyEndPoint() string {
	now := time.Now()

	// See https://docs.fastly.com/api/account#billing
	return fmt.Sprintf("https://api.fastly.com/billing/year/%s/month/%s",
		// See https://golang.org/pkg/time/#pkg-constants
		now.Format("2006"), now.Format("01"))
}

func pickTotalCost(resp *http.Response) (float64, error) {
	type Total struct {
		Cost float64
	}

	type Billing struct {
		Total Total `json:"total"`
	}

	var billing Billing

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&billing)

	if err != nil {
		return 0.0, err
	}

	return billing.Total.Cost, nil
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
