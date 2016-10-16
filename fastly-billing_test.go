package main

import (
	"strings"
	"testing"
	"time"
)

func TestFastlyEndPoint(t *testing.T) {
	const layout = "2006-Jan-02"
	now, _ := time.Parse(layout, "2015-May-03")

	actual := fastlyEndPoint(now)
	const expected = "https://api.fastly.com/billing/year/2015/month/05"
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func TestPickTotalCost_validJSON(t *testing.T) {
	testJSON := `
{
  "total": {
    "cost": 123.456789
  }
}
`
	actual, err := pickTotalCost(strings.NewReader(testJSON))
	const expected = 123.456789
	if actual != expected {
		t.Errorf("expected %f but got %f", expected, actual)
	}
	if err != nil {
		t.Error("`err` should be nil")
	}
}

func TestPickTotalCost_brokenJSON(t *testing.T) {
	testJSON := `
{
  "total": {
    "cost": 123.456789,
  }
}
`
	_, err := pickTotalCost(strings.NewReader(testJSON))
	if err == nil {
		t.Error("`err` should not be nil")
	}
}

func TestPickTotalCost_nonExistentField(t *testing.T) {
	testJSON := `
{
  "foo": {
    "bar": 123.456789
  }
}
`
	_, err := pickTotalCost(strings.NewReader(testJSON))
	if err == nil {
		t.Error("`err` should not be nil")
	}
}
