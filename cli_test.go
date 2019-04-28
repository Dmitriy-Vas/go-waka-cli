package go_waka_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

type Config struct {
	Token string `json:"token"`
}

var (
	config Config
	client *WakaClient
)

// Для начала создайте файл config.json
// Добавьте туда ваш токен от WakaTime (который вы используете с плагинами)
// Пример как должен выглядеть файл:
// {
// 	   "token": "put-your-token-here"
// }
func init() {
	// Загружаем файл config.json
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	// Пытаемся считать данные из файла
	if err = json.Unmarshal(data, &config); err != nil {
		fmt.Println(err)
	}
	// Создаём новый инстанс API
	client = New(config.Token)
}

func TestCommits(t *testing.T) {
	_, err := client.Commits("go-waka-api")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDurations(t *testing.T) {
	date := &Data{Day: 29, Month: 04, Year: 2019}
	_, err := client.Durations(date)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGoals(t *testing.T) {
	_, err := client.Goals()
	if err != nil {
		t.Fatal(err)
	}
}

func TestHeartbeats(t *testing.T) {
	date := &Data{Day: 29, Month: 04, Year: 2019}
	_, err := client.Heartbeats(date)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHeartbeatPost(t *testing.T) {
	heartbeat := &Heartbeat{
		Entity:       "D:/Projects/go-waka-api/.-time/structures.go",
		Type:         "file",
		Category:     "coding",
		Time:         float64(time.Now().Unix()),
		Project:      "go-waka-api",
		Branch:       "master",
		Language:     "Go",
		Dependencies: []string{},
		Lines:        665,
		IsWrite:      false,
	}
	heartbeatResponse, err := client.HeartbeatPost(heartbeat)
	if err != nil {
		t.Fatal(err, heartbeatResponse)
	}
}

func TestLeaders(t *testing.T) {
	_, err := client.Leaders()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMeta(t *testing.T) {
	_, err := client.Meta()
	if err != nil {
		t.Fatal(err)
	}
}

// Untested
/*func TestOrgDashboardMemberSummaries(t *testing.T) {
	start := &..Data{Day: 21, Month: 04, Year: 2019}
	end := &..Data{Day: 29, Month: 04, Year: 2019}
	_, err := client.OrgDashboardMemberSummaries("some-org", "some-dash", "some-member", start, end)
	if err != nil {
		t.Fatal(err)
	}
}*/

// Untested
/*func TestOrgDashboardMembers(t *testing.T) {
	_, err := client.OrgDashboardMembers("some-org", "some-dash")
	if err != nil {
		t.Fatal(err)
	}
}*/

// Untested
/*func TestOrgDashboards(t *testing.T) {
	_, err := client.OrgDashboards("some-org")
	if err != nil {
		t.Fatal(err)
	}
}*/

func TestOrgs(t *testing.T) {
	_, err := client.Orgs()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrivateLeaderboards(t *testing.T) {
	_, err := client.PrivateLeaderboards()
	if err != nil {
		t.Fatal(err)
	}
}

// Untested
/*func TestPrivateLeaderboardsLeaders(t *testing.T) {
	_, err := client.PrivateLeaderboardsLeaders("some-board")
	if err != nil {
		t.Fatal(err)
	}
}*/

func TestProjects(t *testing.T) {
	_, err := client.Projects()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStats(t *testing.T) {
	_, err := client.Stats(RANGE_7_DAYS)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSummaries(t *testing.T) {
	start := &Data{Day: 21, Month: 04, Year: 2019}
	end := &Data{Day: 29, Month: 04, Year: 2019}
	_, err := client.Summaries(start, end)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserAgents(t *testing.T) {
	_, err := client.UserAgents()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUsers(t *testing.T) {
	_, err := client.Users()
	if err != nil {
		t.Fatal(err)
	}
}
