package parser

import (
	"paramparser/custom"
	"strings"
)

func Blacklist(url string) bool {
	if len(url) < 5 {
		return true
	}
	if !strings.Contains(url, MainDomain) {
		return true
	}
	if custom.SliceStrContains(Urls, url) {
		return true
	}
	if len(url) > 4 {
		extension := url[len(url)-4:]
		switch extension {
		case ".png":
			return true
		case ".jpg":
			return true

		}
	}
	return false
}
