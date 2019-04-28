package api

import "encoding/json"

func (w *WakaClient) Leaders() (*Leaders, error) {
	data, err := w.get("leaders", nil)
	if err != nil {
		return nil, err
	}
	var leaders *Leaders
	if err = json.Unmarshal(data, &leaders); err != nil {
		return nil, err
	}
	return leaders, nil
}
