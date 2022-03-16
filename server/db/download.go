package db

import (
	"context"
	"errors"
)

func GetData(url string) (string, bool, bool, string, int, error) {
	var name string
	var isProtected, isLimit bool
	var password string
	var limit int
	err := Env.DB.QueryRow(context.Background(), "select name, is_protected, is_limit, password, \"limit\" from files where url = $1", url).Scan(&name, &isProtected, &isLimit, &password, &limit)
	if err != nil {
		return "", false, false, "", 0, err
	}

	return name, isProtected, isLimit, password, limit, nil
}

func ReduceLimit(url string) error {
	tag, err := Env.DB.Exec(context.Background(), "update files set \"limit\" = \"limit\" -1 where url = $1", url)
	if err != nil {
		return err
	}
	if !tag.Update() {
		return errors.New("Not update")
	}
	return nil
}
