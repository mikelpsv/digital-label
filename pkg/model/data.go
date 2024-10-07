package model

import (
	app "github.com/mlplabs/app-utils"
	"time"
)

type Code struct {
	KeyLink   string `json:"key_link"`
	KeyData   string `json:"key_data"`
	Type      int    `json:"type"`
	Payload   string `json:"payload"`
	Action    string
	CreatedAt string `json:"-"`
}

func (c *Code) Write() error {
	sqlWrite := "INSERT INTO data_links (key_link, key_data, created_at, data_type, payload) VALUES($1, $2, $3, $4, $5)"
	_, err := app.Db.Exec(sqlWrite, c.KeyLink, c.KeyData, time.Now(), c.Type, c.Payload)
	if err != nil {
		return err
	}
	return nil
}

func (c *Code) Get(keyLink string) error {
	sqlGet := "SELECT key_data, data_type, payload FROM data_links WHERE key_link = $1"
	rows, err := app.Db.Query(sqlGet, keyLink)
	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&c.KeyData, &c.Type, &c.Payload)
	if err != nil {
		return err
	}
	return nil
}

func (c *Code) Render() (string, error) {

	return "", nil

}
