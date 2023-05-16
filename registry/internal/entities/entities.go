package entities

import (
	"encoding/json"
	"fmt"
)

type (
	Status string

	RemoteService struct {
		ProcessName  string `json:"process_name"`
		Address      string `json:"address"`
		FinalAddress string `json:"final_address,omitempty"`
	}

	WorkOrder struct {
		ID      string           `json:"id,omitempty"`
		Process string           `json:"process,omitempty"`
		Status  string           `json:"status,omitempty"`
		Params  *json.RawMessage `json:"params,omitempty"`
		Result  *json.RawMessage `json:"result,omitempty"`
	}

	Create struct {
		WorkOrder `json:"work_order,omitempty"`
	}
)

func (r *RemoteService) FormatFinalAddress() {
	r.FinalAddress = fmt.Sprintf("http://www.%s.com", r.ProcessName)
}
