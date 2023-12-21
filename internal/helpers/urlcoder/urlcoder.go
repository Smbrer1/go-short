package urlcoder

import (
	"fmt"
	"math"
	"unicode"
)

// true random short url
// code from https://gist.github.com/bhelx/778542

const (
	base            = 62
	uppercaseOffset = 55
	lowercaseOffset = 61
	digitOffset     = 48
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func trueOrd(char rune) (int64, error) {
	if unicode.IsDigit(char) {
		return int64(char) - digitOffset, nil
	} else if 'A' <= char && char <= 'Z' {
		return int64(char) - uppercaseOffset, nil
	} else if 'a' <= char && char <= 'z' {
		return int64(char) - lowercaseOffset, nil
	} else {
		return 0, fmt.Errorf("%c id not a valid character", char)
	}
}

func trueChar(integer int64) (rune, error) {
	if integer < 10 {
		return rune(integer + digitOffset), nil
	} else if 10 <= integer && integer <= 35 {
		return rune(integer + uppercaseOffset), nil
	} else if 36 <= integer && integer <= 62 {
		return rune(integer + lowercaseOffset), nil
	} else {
		return 0, fmt.Errorf("%d id not a valid integr in the range of base %d", integer, base)
	}
}

func Decode(key string) (int64, error) {
	var intSum int64
	reversedKey := reverse(key)
	for idx, char := range reversedKey {
		val, err := trueOrd(char)
		if err != nil {
			return 0, err
		}
		intSum += val * int64(math.Pow(base, float64(idx)))
	}
	return intSum, nil
}

func Encode(integer int64) (string, error) {
	if integer == 0 {
		return "0", nil
	}

	str := ""
	for integer > 0 {
		remainder := integer % base
		char, err := trueChar(remainder)
		if err != nil {
			return "", err
		}
		str = string(char) + str
	}
	return str, nil
}
