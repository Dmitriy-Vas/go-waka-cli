package api

import "encoding/json"

func (w *WakaClient) Orgs() (*Orgs, error) {
	data, err := w.get("users/current/orgs", nil)
	if err != nil {
		return nil, err
	}
	var orgs *Orgs
	if err = json.Unmarshal(data, &orgs); err != nil {
		return nil, err
	}
	return orgs, nil
}
