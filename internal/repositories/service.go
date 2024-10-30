package repositories

import (
	"database/sql"
	"github.com/mikelpsv/digital-label/internal/repositories/dbo"
	"time"
)

type ServiceRepository struct {
	Db *sql.DB
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{
		db,
	}
}

func (r *ServiceRepository) GetLink(keyLink string) (*dbo.LinkData, error) {
	ld := dbo.LinkData{}
	sqlGet := "SELECT key_link, key_data, data_type, payload FROM data_links WHERE key_link = $1"
	rows, err := r.Db.Query(sqlGet, keyLink)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&ld.KeyLink, &ld.KeyData, &ld.Type, &ld.Payload)
		if err != nil {
			return &ld, err
		}
	}
	return &ld, nil
}

func (r *ServiceRepository) WriteData(data *dbo.LinkData) error {
	sqlWrite := "INSERT INTO data_links (key_link, key_data, created_at, data_type, payload) VALUES($1, $2, $3, $4, $5)"
	_, err := r.Db.Exec(sqlWrite, data.KeyLink, data.KeyData, time.Now(), data.Type, data.Payload)
	if err != nil {
		return err
	}
	return nil
}
