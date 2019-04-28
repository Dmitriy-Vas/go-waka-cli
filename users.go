package go_waka_api

import (
	"encoding/json"
)

func (w *WakaClient) Users() (*User, error) {
	data, err := w.get("users/current", nil)
	if err != nil {
		return nil, err
	}
	var user *User
	if err = json.Unmarshal(data, &user); err != nil {
		return nil, err
	}
	return user, nil
}
