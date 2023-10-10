package fstore

import "fmt"

type Database struct {
	hashValues map[string]string
	hashKeys   map[string]string
}

func GetDatabase() Database {
	return Database{
		hashValues: map[string]string{},
		hashKeys:   map[string]string{},
	}
}

func (d *Database) SaveHash(val string) string {
	h := fmt.Sprintf("h_%v", len(d.hashValues))
	d.hashValues[h] = val
	return h
}

func (d *Database) SaveKey(val string) string {
	h := fmt.Sprintf("h_%v", len(d.hashKeys))
	d.hashKeys[h] = val
	return h
}
