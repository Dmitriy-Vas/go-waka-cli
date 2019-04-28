package go_waka_api

import "encoding/json"

func (w *WakaClient) Goals() (*Goals, error) {
	data, err := w.get("users/current/goals", nil)
	if err != nil {
		return nil, err
	}
	var goals *Goals
	if err = json.Unmarshal(data, &goals); err != nil {
		return nil, err
	}
	return goals, nil
}
