package stringprocess

import (
  "testing"
  "reflect"
)

func TestProcessString(t *testing.T) {
  t.Run("converts text into slice of words", func(t *testing.T) {
    text := "ein, zwei"

    got, _ := ProcessString(text)
    want := []string{"ein", "zwei"}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("Converts text to lower case", func(t *testing.T) {
    text := "Danke"

    got, _ := ProcessString(text)
    want := []string{"danke"}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("Sorts words alphabetically", func(t *testing.T) {
    text := "Ich warte hier."

    got, _ := ProcessString(text)
    want := []string{"hier", "ich", "warte"}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %q, want %q", got, want)
    }
  })

  t.Run("returns empty slice if passed empty string", func(t *testing.T) {
    text := ""

    _, err := ProcessString(text)

    if err == nil {
      t.Error("expected an error")
    }
  })

  t.Run("remove word duplication", func(t *testing.T) {
    text := "ein, zwei, ein, ein ein: zwei"

    got, _ := ProcessString(text)
    want := []string{"ein", "zwei"}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })
}
