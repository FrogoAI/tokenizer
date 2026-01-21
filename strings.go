package tokenizer

import (
	"strings"
	"unicode"
	"unsafe"

	"github.com/twmb/murmur3"
	"golang.org/x/text/unicode/norm"
)

const (
	EmailTagStart = "+"
	EmailAt       = "@"
)

func ABTest(data, salt []byte, groups ...uint64) uint64 {
	var total uint64
	for _, group := range groups {
		total += group
	}

	if total == 0 {
		return 0
	}

	return murmur3.Sum64(append(data, salt...)) % total
}

func SanitizeEmail(email string) string {
	return strings.Join(SplitBetweenTokens(strings.ToLower(email), EmailTagStart, EmailAt), EmailAt)
}

func NFDLowerString(str string) string {
	return strings.ToLower(norm.NFD.String(strings.TrimSpace(str)))
}

func CommonString(str string) string {
	var result strings.Builder

	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// SplitBetweenTokens take string and one or two tokens, and cut everything between two tokens, or between two copies of first token
func SplitBetweenTokens(data string, keys ...string) []string {
	if data == "" {
		return []string{}
	}

	var key1, key2 string

	switch {
	case len(keys) == 1: //nolint:mnd
		key1 = keys[0]
		key2 = keys[0]
	case len(keys) >= 2: //nolint:mnd
		key1 = keys[0]
		key2 = keys[1]
	default:
		return []string{data}
	}

	if key1 == "" || key2 == "" {
		return []string{data}
	}

	s := strings.Index(data, key1)
	if s <= -1 {
		return []string{data}
	}

	part1 := data[0:s]

	s += len(key1)

	e := strings.Index(data[s:], key2)
	if e <= -1 {
		return []string{part1}
	}

	e += len(key2)

	part2 := data[s+e:]

	return []string{part1, part2}
}

// ByteSliceToString cast given bytes to string, without allocation memory
func ByteSliceToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) //nolint
}

// Between function to get content between two keys
func Between(data string, keys ...string) string {
	var key1, key2 string

	switch {
	case len(keys) == 1: //nolint:mnd
		key1 = keys[0]
		key2 = keys[0]
	case len(keys) >= 2: //nolint:mnd
		key1 = keys[0]
		key2 = keys[1]
	default:
		return ""
	}

	if key1 == "" || key2 == "" {
		return ""
	}

	s := strings.Index(data, key1)
	if s <= -1 {
		return ""
	}

	s += len(key1)

	e := strings.Index(data[s:], key2)
	if e <= -1 {
		return ""
	}

	return strings.TrimSpace(data[s : s+e])
}
