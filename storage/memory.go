package storage

var urls = map[string]string{}

func GetStorage(short string) (string, error) {
	url, err := urls[short]
	if !err {
		return "", NotFound
	}
	return url, nil
}

func PostStorage(short string, original string) error {
	_, err := urls[short]
	if err {
		return AlreadyExists
	}
	urls[short] = original
	return nil
}