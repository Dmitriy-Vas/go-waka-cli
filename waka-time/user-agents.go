package api

import "encoding/json"

func (w *WakaClient) UserAgents() (*UserAgents, error) {
	data, err := w.get("users/current/user_agents", nil)
	if err != nil {
		return nil, err
	}
	var userAgents *UserAgents
	if err = json.Unmarshal(data, &userAgents); err != nil {
		return nil, err
	}
	return userAgents, nil
}
