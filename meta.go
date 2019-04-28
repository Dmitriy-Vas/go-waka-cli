package go_waka_api

import "encoding/json"

func (w *WakaClient) Meta() (*Meta, error) {
	data, err := w.get("meta", nil)
	if err != nil {
		return nil, err
	}
	var meta *Meta
	if err = json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}
	return meta, nil
}
