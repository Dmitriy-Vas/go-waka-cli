package api

import "encoding/json"

func (w *WakaClient) Summaries(start *Data, end *Data) (*Summaries, error) {
	query := make(map[string]string)
	query["start"] = start.format()
	query["end"] = end.format()
	data, err := w.get("users/current/summaries", query)
	if err != nil {
		return nil, err
	}
	var summaries *Summaries
	if err = json.Unmarshal(data, &summaries); err != nil {
		return nil, err
	}
	return summaries, nil
}
