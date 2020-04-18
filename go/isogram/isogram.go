package isogram

import (
	"strings"
)

// comment to document it comment
func IsIsogram(alph string) bool {
	// some comments
	alph = strings.ToLower(alph)
	if len(alph) == 0 || len(alph) == 1 {
		return true
	}
	ff := false
	for i := 0; i < len(alph); i++ {
		//fmt.Printf("%v\n", alph[i])
		for j := i + 1; j < len(alph); j++ {
			//fmt.Println(alph[i], alph[j])
			if alph[i] == alph[j] {
				if alph[i] <= 122 && alph[i] >= 97 && alph[j] <= 122 && alph[j] >= 97 {
					ff = true
					return false
				}
			}
		}
	}
	if ff {
		return false
	} else {
		return true
	}
}
