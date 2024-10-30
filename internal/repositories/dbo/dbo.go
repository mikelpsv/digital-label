package dbo

type LinkData struct {
	KeyLink   string `json:"key_link"`
	KeyData   string `json:"key_data"`
	Type      int    `json:"type"`
	Payload   string `json:"payload"`
	Action    string
	CreatedAt string `json:"-"`
}
