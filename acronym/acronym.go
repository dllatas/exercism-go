// Package acronym provides a function to transform strings
package acronym

import(
    "strings"
)

// Abbreviate gets a string and returns its acronym
func Abbreviate(s string) string {
        s = strings.Replace(s, "-", " ", -1)
        words := strings.Split(s, " ")
        abv := ""
        for _, word := range words {
            abv += strings.ToUpper(string(word[0]))
        }	
        return abv
}
