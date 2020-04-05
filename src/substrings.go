package src

import s "strings"

func CheckString(input string) string {
	i := 0
	AB := false
	BA := false

	for i < len(input)-1 {
		if input[i] == 'A' && input[i+1] == 'B' && !AB {
			AB = true
			i += 2
		}
		if input[i] == 'B' && input[i+1] == 'A' && !BA {
			BA = true
			i += 2
		}
		if AB && BA {
			return "YES"
		}
		i++
	}
	return "NO"
}

//One more solution
func CheckString2(input string) string {
	indexAB := s.Index(input, "AB") + 1
	indexBA := s.LastIndex(input, "BA")
	if s.Contains(input, "AB") && s.Contains(input, "BA") && indexAB != indexBA {
		return "YES"
	}
	return "NO"
}
