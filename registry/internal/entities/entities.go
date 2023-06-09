package entities

type (
	RemoteService struct {
		ProcessName string `json:"process_name"`
		Address     string `json:"address"`
	}

	Order struct {
		ID      string `json:"id,omitempty"`
		Process string `json:"process,omitempty"`
	}
)
