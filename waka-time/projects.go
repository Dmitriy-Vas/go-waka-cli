package api

import (
	"encoding/json"
)

func (w *WakaClient) Projects() (*Projects, error) {
	data, err := w.get("users/current/projects", nil)
	if err != nil {
		return nil, err
	}
	var projects *Projects
	if err = json.Unmarshal(data, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}
