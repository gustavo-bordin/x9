package repository

import (
	_ "embed"
	"fmt"
)

//go:embed queries/create_track.sql
var createTrackSql string

func (r Repository) CreateUrl(url string) error {
	query := fmt.Sprintf(createTrackSql, url)
	_, err := r.db.Exec(query)
	fmt.Println(err)

	return err
}
