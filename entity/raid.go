package entity

type Raid struct {
	SystemID   string `json:"system_id,omitempty"`
	ID         int    `json:"id,omitempty"`
	Incomplete bool   `json:"__incomplete__,omitempty"`
}
