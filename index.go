package piscine

func Index(s string, toFind string) int {
	lenS := 0
	for range s {
		lenS++
	}

	lenSub := 0
	for range toFind {
		lenSub++
	}

	if lenSub == 0 {
		return 0
	}

	for i := 0; i <= lenS-lenSub; i++ {
		if s[i:i+lenSub] == toFind {
			return i
		}
	}

	return -1
}
