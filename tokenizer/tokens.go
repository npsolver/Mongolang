package tokenizer

import (
	"fmt"
	"regexp"
)

// Length limits in bytes
const (
	databaseLengthLimit = 64
)

const (
	// Non-leaf Tokens
	RESULTS = iota
	QUERY
	EXP
	DBMETHOD
	COLMETHOD
	KEY
	COLON

	// Leaf Tokens
	DATABASE
	COLLECTION
	DBMETHODNAME
	COLMETHODNAME
	DOT
	ID
	VALUE
	FIRSTBRACKETLEFT
	FIRSTBRACKETRIGHT
	SECONDBRACKETLEFT
	SECONDBRACKETRIGHT
	THIRDBRACKETLEFT
	THIRDBRACKETRIGHT
)

type TokenType int

type Token struct {
	tokenType TokenType
	val       string
}

func newRegex(s string) *regexp.Regexp {
	regex, err := regexp.Compile(s)
	if err != nil {
		panic(fmt.Sprintf("Invalid Regex: %s\n", s))
	}
	return regex
}

func isToken(r string, b []byte) (bool, error) {
	regex := newRegex(r)
	return regex.Match(b), nil
}

func (tk *TokenType) CreateToken(val string) *Token { return &Token{*tk, val} }

// For leaf tokens
func MatchToken(s string, tk int) (bool, error) {
	b := []byte(s)
	switch tk {
	case DATABASE:
		return isDatabaseToken(b)
	case COLLECTION:
		return true, nil
	case FIRSTBRACKETLEFT:
		return isFirstBracketLeft(b)
	case FIRSTBRACKETRIGHT:
		return isFirstBracketRight(b)
	case SECONDBRACKETLEFT:
		return isSecondBracketLeft(b)
	case SECONDBRACKETRIGHT:
		return isSecondBracketRight(b)
	case THIRDBRACKETLEFT:
		return isThirdBracketLeft(b)
	case THIRDBRACKETRIGHT:
		return isThirdBracketRight(b)
	case DOT:
		return isDot(b)
	case ID:
		return isID(b)
	}
	return false, nil
}
