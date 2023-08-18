package common

import "strings"

func GetLastSlashValue(url string) string {
	urlParts := strings.Split(strings.TrimSuffix(url, "/"), "/")
	return urlParts[len(urlParts)-1]

}
