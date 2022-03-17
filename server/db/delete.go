package db

import (
	"context"
	"time"
)

func GetToDelete() ([]string, error) {
	files := make([]string, 0)
	rows, err := Env.DB.Query(context.Background(), " select url, upload_time from files")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var file string
		var t time.Time

		err = rows.Scan(&file, &t)
		if err != nil {
			return nil, err
		}
		if time.Now().Sub(t) > time.Minute*10 {
			files = append(files, file)
			tag, err := Env.DB.Exec(context.Background(), "delete from files where url = $1", file)
			if err != nil {
				return nil, err
			}
			if !tag.Delete() {
				return nil, err
			}
		}

	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return files, nil

}
