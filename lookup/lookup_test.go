package lookup_test

import (
  "testing"
  "reflect"
  lookup "github.com/turnerem/vocabuli/lookup"
)

func TestLookup(t *testing.T) {
  t.Run("takes slice of words and returns map of words and translations", func(t *testing.T) {
    data := []string{"ein", "zwei"}

    got := lookup.Lookup(mockDictSearch, mockGetURL, data)
    want := map[string]string{
      "ein": "one",
      "zwei": "two",
    }

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v got %v", got, want)
    }
  })
}

func mockDictSearch(url lookup.URLType) string {
  switch(url) {
  case lookup.URLType("ein"):
    return "one"
  case lookup.URLType("zwei"):
    return "two"
  }

  return ""
}

func mockGetURL(word string) lookup.URLType {
  return lookup.URLType(word)
}
