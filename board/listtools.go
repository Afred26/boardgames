package board

// ContainsOnly erwartet eine Liste von Strings und liefert true,
// wenn all diese Strings gleich dem gegebenen String s sind, sonst false.
func ContainsOnly(list []string, s string) bool {
	result := true
	for _, el := range list {
		if el != s {
			result = false
		}
	}
	return result
}

// ContainsAny erwartet eine Liste von Strings und liefert true,
// wenn darin mindestens ein String gleich dem gegebenen String s ist, sonst false.
func ContainsAny(list []string, s string) bool {
	result := false
	for _, el := range list {
		if el == s {
			result = true
		}
	}
	return result
}
