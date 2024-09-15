package helpers

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IsEqual(s1 string, s2 string) bool {
	return s1 == s2
}
