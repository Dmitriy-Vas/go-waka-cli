# go-waka-cli

Simple API wrapper for WakaTime built with Go.

### Table of Contents

+ [Description](https://github.com/Dmitriy-Vas/go-waka-cli#Description)
+ [Requirements](https://github.com/Dmitriy-Vas/go-waka-cli#Requirements)
+ [Docs](https://github.com/Dmitriy-Vas/go-waka-cli#Docs)
+ [Usage](https://github.com/Dmitriy-Vas/go-waka-cli#Usage)

### Description

This program has learning purposes and aimed only for self education.
<br>
__Program temporary supports only *current* user.__

### Requirements

+ [Go](https://golang.org/dl/)

### Docs

Official documentation may be found in the [WakaTime](https://wakatime.com/developers) website.

Endpoints list:

- Commits
- Durations
- Goals
- Heartbeats (both POST and GET)
- Leaders
- Meta
- Org-Dashboard-Member-Summaries (Not tested)
- Org-Dashboard-Members (Not tested)
- Org-Dashboards (Not tested)
- Orgs
- Private-Leaderboards
- Private-Leaderboards-Leaders (Not tested)
- Projects
- Stats
- Summaries
- User-Agents
- Users

### Usage

+ Install Go
+ Clone repo to your computer

```
go get -u github.com/Dmitriy-Vas/go-waka-cli
```

+ Create new Go file and import `github.com/Dmitriy-Vas/go-waka-cli`
+ Use built-in method "New" to get wrapper instance
<br>
Example:

```go
client := New("put-your-token-here")

user, err := client.Users()
if err != nil {
	log.Fatal(err)
}
fmt.Println("Logged in as", user.Data.UserName)
```
