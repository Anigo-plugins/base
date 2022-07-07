package tools

func Contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {

			return true
		}

	}

	return false
}
