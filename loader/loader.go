package loader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"unicode/utf8"
)

const (
	minLength = 8
	maxLength = 64
)

type commonPasswords struct {
	list []string
}

// load in the common password file, returns a pointer to the struct containing list of strings.
func LoadCommon(path string) (*commonPasswords, error) {
	s := &commonPasswords{list: make([]string, 0)}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return s, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s.list = append(s.list, scanner.Text())
	}

	return s, nil
}

func (s *commonPasswords) sortCommon() *commonPasswords {
	sort.Strings(s.list)
	return s
}

func (s *commonPasswords) isCommon(pass string) bool {
	// make sure all the strings are sorted
	if !sort.StringsAreSorted(s.list) {
		s = s.sortCommon()
	}
	idx := sort.SearchStrings(s.list, pass)
	// since SearchStrings returns the idx of where the string exists or can be inserted into
	// lets make sure there are no out of bound errors when finally comparing the results to the password arg
	if len(s.list) > idx && s.list[idx] == pass {
		return true
	}

	return false
}

// Return the original password in string form, with error.
func (s *commonPasswords) IsValid(pass []byte) (string, error) {
	// new rune slice will contain any unprintable chars replaced with *
	isInvalidASCII, runes := s.containsInvalidASCII(pass)
	passStr := string(runes)
	if isInvalidASCII {
		return passStr, errors.New("Error: Invalid Charaters")
	}
	if s.isOverMaximumLength(pass) {
		return passStr, errors.New("Error: Too Long")
	}
	if s.isUnderMinimumLength(pass) {
		return passStr, errors.New("Error: Too Short")
	}
	if s.isCommon(passStr) {
		return passStr, errors.New("Error: Too Common")
	}
	return passStr, nil
}

func (s *commonPasswords) isUnderMinimumLength(pass []byte) bool {
	return utf8.RuneCount(pass) < minLength
}

func (s *commonPasswords) isOverMaximumLength(pass []byte) bool {
	return utf8.RuneCount(pass) > maxLength
}

// If the byte array contains an invalid character (non-inclusive ascii code 32 to 126),
// return true and the password with any invalid chars replaced with "*""
func (s *commonPasswords) containsInvalidASCII(pass []byte) (bool, []rune) {
	var (
		tempArr = pass
		result  = false
		runes   = make([]rune, 0)
	)

	for len(tempArr) > 0 {
		r, size := utf8.DecodeRune(tempArr)
		// just handling ASCII chars as below for now
		if r < 32 || r > 126 {
			result = true
			runes = append(runes, '*')
		} else {
			runes = append(runes, r)
		}
		tempArr = tempArr[size:]
	}
	return result, runes
}
