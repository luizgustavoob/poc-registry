package entities

type (
	RemoteService struct {
		ProcessName string `json:"process_name"`
		Address     string `json:"address"`
	}

	Create struct {
		ID      string `json:"id"`
		Process string `json:"process"`
	}
)
