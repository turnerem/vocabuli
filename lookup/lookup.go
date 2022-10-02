package lookup

import (
  "fmt"
)

// URLType is just a string
type URLType string

// DictionarySearcher is a function type
type DictionarySearcher func(u URLType) string

// GetURL returns the URLType that we want to lookup
type GetURL func(word string) URLType

type result struct {
  word string
  translation string
}

// Lookup translates the slice of words
func Lookup(d DictionarySearcher, u GetURL, words []string) map[string]string {
  // TODO: language is configurable

  results := make(map[string]string)
  resultChannel := make(chan result)

  // loop words
  // lookup
  // pop into map
  for _, word := range(words) {
    go func(w string) {
      URLType := u(w)
      resultChannel <- result{w, d(URLType)}
    }(word)
  }

  for range(words) {
    r := <-resultChannel
    results[r.word] = r.translation
  }

  return results
}

func getURLGerman(word string) URLType {
  format := getURLFormatForLanguage("german")

  return URLType(fmt.Sprintf(format, word))
}

func getURLFormatForLanguage(language string) string {
  switch(language) {
  case "german":
    return "https://german-english-dictionary-api.uc.r.appspot.com/translate?term=%s&limit=1"
  }

  return "fakeURLType_%s"
}
