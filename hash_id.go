package hash_id

import (
	"log"
	"regexp"
	"unicode"
)

func isAlpha(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}
func IsNum(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
func isAlphaNum(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

// Returns possible hash algorithm for given hash
func Identify(hash string) []string {
	var possible_hashes []string

	var ok bool
	var length int
	var format string

	for name, info := range ALGORITHMS {
		// Check length of algorithm adn sent hash
		length, ok = info["Length"].(int)
		if !ok {
			log.Fatalf("...")
		}
		if len(hash) != length {
			continue
		}
		// Check format of algorithm and sent hash
		format, ok = info["Format"].(string)
		if !ok {
			log.Fatalf("---")
		}
		switch format {
		case "AlphaNum":
			if !isAlphaNum(hash) {
				continue
			}
		case "Alpha":
			if !isAlpha(hash) {
				continue
			}
		case "Num":
			if !IsNum(hash) {
				continue
			}
		default:
			log.Fatalf("Unsupported format for %s algorithm", name)
		}
		possible_hashes = append(possible_hashes, name)
	}
	return possible_hashes
}
