package collection

import (
	"testing"
)

func TestSetOnlyIncludesAddedItems(t *testing.T) {
  set := NewSet()
  set.Add("item 1")
  set.Add("item 2")
  if set.Contains("notAddedItem") {
    t.Error("Expected set to not contain notAddedItem")
  }
  if set.Contains("") {
    t.Error("Expected set to not contain not added empty string")
  }
  if !set.Contains("item 1") {
    t.Error("Expected set to contain item 1")
  }
  if !set.Contains("item 2") {
    t.Error("Expected set to contain item 2")
  }
}
