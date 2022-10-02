package lookup_test

import (
  "testing"
)

func TestLookup(t *testing.T) {
  t.Run("takes slice of words and returns map of words and translations", func(t *testing.T) {
    data := []string{"ein", "zwei"}

    got := lookup.Lookup(data)
    want := map[string]string{
      "ein": "one",
      "zwei": "two",
    }

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v got %v", got, want)
    }
  })
}
