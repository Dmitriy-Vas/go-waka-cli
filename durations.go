package go_waka_api

import "encoding/json"

func (w *WakaClient) Durations(data *Data) (*Durations, error) {
	query := make(map[string]string)
	query["date"] = data.format()
	bytes, err := w.get("users/current/durations", query)
	if err != nil {
		return nil, err
	}
	var durations *Durations
	if err = json.Unmarshal(bytes, &durations); err != nil {
		return nil, err
	}
	return durations, nil
}
