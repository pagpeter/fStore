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

func isInMap(m map[string]string, key string) string {
	for k, v := range m {
		if v == key {
			return k
		}
	}
	return ""
}

func (d *Database) SaveHash(val string) string {
	r := isInMap(d.hashValues, val)
	if r != "" {
		return r
	}
	h := fmt.Sprintf("h_%v", len(d.hashValues))
	d.hashValues[h] = val
	return h
}

func (d *Database) SaveKey(val string) string {
	r := isInMap(d.hashKeys, val)
	if r != "" {
		return r
	}
	h := fmt.Sprintf("h_%v", len(d.hashKeys))
	d.hashKeys[h] = val
	return h
}
