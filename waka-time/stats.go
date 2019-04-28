package api

import (
	"encoding/json"
)

var (
	RANGE_7_DAYS   = &timeRange{"last_7_days"}
	RANGE_30_DAYS  = &timeRange{"last_30_days"}
	RANGE_6_MONTHS = &timeRange{"last_6_months"}
	RANGE_YEAR     = &timeRange{"last_year"}
)

func (w *WakaClient) Stats(timeRange *timeRange) (*Stats, error) {
	data, err := w.get("users/current/stats/"+timeRange.string, nil)
	if err != nil {
		return nil, err
	}
	var stats *Stats
	if err = json.Unmarshal(data, &stats); err != nil {
		return nil, err
	}
	return stats, nil
}
