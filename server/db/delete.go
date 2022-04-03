package db

import (
	"context"
	"time"
)

func GetToDelete() ([]string, float64, error) {
	files := make([]string, 0)
	rows, err := Env.DB.Query(context.Background(), " select url, upload_time, duration from files")
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var count float64

	for rows.Next() {

		count++
		var file string
		var t time.Time
		var duration int

		err = rows.Scan(&file, &t, &duration)
		if err != nil {
			return nil, 0, err
		}
		if time.Now().Sub(t) > time.Minute*time.Duration(duration) {
			files = append(files, file)
			tag, err := Env.DB.Exec(context.Background(), "delete from files where url = $1", file)
			if err != nil {
				return nil, 0, err
			}
			if !tag.Delete() {
				return nil, 0, err
			}
		}

	}

	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}

	return files, count, nil

}
