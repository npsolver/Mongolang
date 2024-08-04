package tokenizer

import "errors"

func isID(b []byte) (bool, error)  { return isToken(``, b) }
func isDot(b []byte) (bool, error) { return isToken(`\.`, b) }

func isFirstBracketLeft(b []byte) (bool, error)  { return isToken(`(`, b) }
func isFirstBracketRight(b []byte) (bool, error) { return isToken(`(`, b) }

func isSecondBracketLeft(b []byte) (bool, error)  { return isToken(`{`, b) }
func isSecondBracketRight(b []byte) (bool, error) { return isToken(`}`, b) }

func isThirdBracketLeft(b []byte) (bool, error)  { return isToken(`[`, b) }
func isThirdBracketRight(b []byte) (bool, error) { return isToken(`]`, b) }

func isDatabaseToken(b []byte) (bool, error) {
	if len(b) > databaseLengthLimit {
		return false, errors.New("database name length more than 64 bytes")
	}
	return isToken(`[^/\. "$]`, b)
}
