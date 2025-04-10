package custom

func SliceStrContains(yourslice []string, elemento string) bool {
	for _, value := range yourslice {
		if value == elemento {
			return true
		}
	}
	return false
}
