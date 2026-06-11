package raindrops

import (
	"strconv"
	"strings"
)

type rule struct {
	factor int
	sound  string
}

func Convert(number int) string {

	rules := []rule{
		{factor: 3, sound: "Pling"},
		{factor: 5, sound: "Plang"},
		{factor: 7, sound: "Plong"},
	}

	var result strings.Builder

	for _, r := range rules {
		if number%r.factor == 0 {
			result.WriteString(r.sound)
		}
	}

	if result.String() == "" {
		return strconv.Itoa(number)
	}

	return result.String()
}
