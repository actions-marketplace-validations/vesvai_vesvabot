package github

import (
	"encoding/json"
	"fmt"
	"os"
)

type Event struct {
	Action            string          `json:"action"`
	Issue             *Issue          `json:"issue"`
	PullRequest       *PullRequest    `json:"pull_request"`
	Comment           *Comment        `json:"comment"`
	Review            *Review         `json:"review"`
	Label             *Label          `json:"label"`
	Sender            *User           `json:"sender"`
	RequestedReviewer *User           `json:"requested_reviewer"`
	Release           *ReleasePayload `json:"release"`
	Raw               json.RawMessage `json:"-"`
}

type Issue struct {
	Number      int             `json:"number"`
	Title       string          `json:"title"`
	Body        string          `json:"body"`
	State       string          `json:"state"`
	User        *User           `json:"user"`
	PullRequest json.RawMessage `json:"pull_request"`
}

type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	State  string `json:"state"`
	Draft  bool   `json:"draft"`
	Merged bool   `json:"merged"`
	User   *User  `json:"user"`
	Head   *Ref   `json:"head"`
	Base   *Ref   `json:"base"`
}

type Ref struct {
	Ref  string `json:"ref"`
	SHA  string `json:"sha"`
	Repo *struct {
		FullName string `json:"full_name"`
	} `json:"repo"`
}

type Comment struct {
	ID        int64  `json:"id"`
	Body      string `json:"body"`
	User      *User  `json:"user"`
	IssueURL  string `json:"issue_url"`
	HTMLURL   string `json:"html_url"`
	CreatedAt string `json:"created_at"`
}

type Review struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
	User *User  `json:"user"`
}

type Label struct {
	Name string `json:"name"`
}

type User struct {
	Login string `json:"login"`
	Type  string `json:"type"`
}

type ReleasePayload struct {
	ID     int64  `json:"id"`
	Tag    string `json:"tag_name"`
	Name   string `json:"name"`
	Body   string `json:"body"`
	Draft  bool   `json:"draft"`
	URL    string `json:"html_url"`
	Author *User  `json:"author"`
}

func LoadEvent(path string) (*Event, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("github: read event file %s: %w", path, err)
	}
	var ev Event
	if err := json.Unmarshal(data, &ev); err != nil {
		return nil, fmt.Errorf("github: parse event file %s: %w", path, err)
	}
	ev.Raw = data
	return &ev, nil
}

func (e *Event) IsPullRequestConversation() bool {
	return e.Issue != nil && len(e.Issue.PullRequest) > 0
}

func (e *Event) IssueNumber() int {
	if e.Issue != nil {
		return e.Issue.Number
	}
	if e.PullRequest != nil {
		return e.PullRequest.Number
	}
	return 0
}

func (e *Event) CommentNumber() int {
	if e.Comment != nil && e.Issue != nil {
		return e.Issue.Number
	}
	return e.IssueNumber()
}

func (e *Event) Actor() string {
	if e.Sender != nil {
		return e.Sender.Login
	}
	return ""
}

func Mentions(body, botLogin string) bool {
	return botLogin != "" && len(body) >= len(botLogin)+1 &&
		containsWord(body, botLogin)
}

func containsWord(body, word string) bool {
	for i := 0; i+len(word) <= len(body); i++ {
		if body[i:i+len(word)] != word {
			continue
		}
		beforeOK := i == 0 || body[i-1] == '@'
		after := i + len(word)
		afterOK := after == len(body) || isPunct(body[after])
		if beforeOK && afterOK {
			return true
		}
	}
	return false
}

func isPunct(b byte) bool {
	switch b {
	case ' ', '\n', '\t', ',', '.', ';', ':', '!', '?', ')', ']', '}', '>':
		return true
	}
	return false
}
