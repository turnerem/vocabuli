package stringprocess

import (
  "regexp"
  "sort"
  "strings"
  "errors"
)

// Words is a list of words extracted from the string
//type Words map[string]string

// ProcessString transforms a block of text into a list of words
func ProcessString(str string) ([]string, error) {
  lastWord := ""
  newSl := []string{}

  if len(str) == 0 {
    return nil, errors.New("Empty string cannot be processed")
  }

  // convert to lower case
  str = strings.ToLower(str)

  r, _ := regexp.Compile("[a-zA-Z]+")

  // extract words from tet into slice
  sl := r.FindAllString(str, -1)

  // sort
  sort.Strings(sl)

  // remove duplicated words
  for _, word := range(sl) {
    if word != lastWord {
      newSl = append(newSl, word)
      lastWord = word
    }
  }

  return newSl, nil

}
