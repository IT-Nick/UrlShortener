package storage

import "sync"

var urls = map[string]string{}
var m = sync.Mutex{}

func PostStorage(short string, original string) error {
	m.Lock()
	_, err := urls[short]
	if err {
		m.Unlock()
		return AlreadyExists
	}
	urls[short] = original
	m.Unlock()
	return nil
}

func GetStorage(short string) (string, error) {
	m.Lock()
	url, err := urls[short]
	if !err {
		m.Unlock()
		return "", NotFound
	}
	m.Unlock()
	return url, nil
}
