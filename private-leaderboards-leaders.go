package go_waka_api

import "encoding/json"

func (w *WakaClient) PrivateLeaderboardsLeaders(board string) (*PrivateLeaderboardsLeaders, error) {
	data, err := w.get("users/current/leaderboards/"+board, nil)
	if err != nil {
		return nil, err
	}
	var privateLeaderboardsLeaders *PrivateLeaderboardsLeaders
	if err = json.Unmarshal(data, &privateLeaderboardsLeaders); err != nil {
		return nil, err
	}
	return privateLeaderboardsLeaders, nil
}
