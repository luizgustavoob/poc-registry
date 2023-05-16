package entities

import "encoding/json"

type RemoteService struct {
	ProcessName string `json:"process_name"`
	Address     string `json:"address"`
}

type Target struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Create struct {
	Type      string           `json:"type"`
	Target    Target           `json:"target"`
	Process   string           `json:"process"`
	Status    string           `json:"status"`
	Params    *json.RawMessage `json:"params,omitempty"`
	Assignees []string         `json:"assignees"`
	ID        string           `json:"id"`
}
