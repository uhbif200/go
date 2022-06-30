package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	words := regexp.MustCompile("-|\\s|_").Split(s, -1)
	if len(words[0]) > 1 {
		words[0] = string(words[0][0]) + strings.ToLower(words[0][1:])
	}

	for i := 1; i < len(words); i++ {
		if len(words[i]) > 1 {
			words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
		} else {
			words[i] = strings.ToUpper(string(words[i][0]))
		}
	}
	return strings.Join(words, "")
}

func main() {
	fmt.Printf("ToCamelCase(\"Hellp-pfpe_Zalupa-jopa-\"): %v\n", ToCamelCase("you_have_chosen_to_translate_this_kata_For_your_convenience_we_have_provided_the_existing_test_cases_used_for_the_language_that_you_have_already_completed_as_well_as_all_of_the_other_related_fields"))
}
