package go_waka_api

import (
	waka "./waka-time"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type Config struct {
	Token string `json:"token"`
}

func TestWakaStats(t *testing.T) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		t.Fatal(err)
	}

	var cfg *Config
	if err = json.Unmarshal(data, &cfg); err != nil {
		t.Fatal(err)
	}

	client := waka.New(cfg.Token)

	stats, err := client.Stats(waka.RANGE_7_DAYS)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%+v", *stats))
}
