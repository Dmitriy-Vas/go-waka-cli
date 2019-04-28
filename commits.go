package go_waka_api

import (
	"encoding/json"
	"errors"
	"regexp"
)

func (w *WakaClient) Commits(project string) (*Commits, error) {
	data, err := w.get("users/current/projects/"+project+"/commits", nil)
	if err != nil {
		return nil, err
	}
	matched, _ := regexp.Match("error", data)
	if matched {
		return nil, errors.New(string(data))
	}
	var commits *Commits
	if err = json.Unmarshal(data, &commits); err != nil {
		return nil, err
	}
	return commits, nil
}
