package go_waka_api

import (
	"encoding/json"
)

func (w *WakaClient) HeartbeatPost(heartbeat *Heartbeat) (*HeartbeatResponse, error) {
	data, err := json.Marshal(heartbeat)
	if err != nil {
		return nil, err
	}
	resp, err := w.post("users/current/heartbeats", data)
	if err != nil {
		return nil, err
	}
	var heartbeatResponse *HeartbeatResponse
	if err = json.Unmarshal(resp, &heartbeatResponse); err != nil {
		return nil, err
	}
	return heartbeatResponse, nil
}

func (w *WakaClient) Heartbeats(data *Data) (*Heartbeats, error) {
	query := make(map[string]string)
	query["date"] = data.format()
	bytes, err := w.get("users/current/heartbeats", query)
	if err != nil {
		return nil, err
	}
	var heartbeats *Heartbeats
	if err = json.Unmarshal(bytes, &heartbeats); err != nil {
		return nil, err
	}
	return heartbeats, nil
}
