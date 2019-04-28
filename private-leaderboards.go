package go_waka_api

import "encoding/json"

func (w *WakaClient) PrivateLeaderboards() (*PrivateLeaderboards, error) {
	data, err := w.get("users/current/leaderboards", nil)
	if err != nil {
		return nil, err
	}
	var privateLeaderboards *PrivateLeaderboards
	if err = json.Unmarshal(data, &privateLeaderboards); err != nil {
		return nil, err
	}
	return privateLeaderboards, nil
}
