package lib

import (
	"reflect"
	"testing"
)

func TestCache_WriteToState(t *testing.T) {
	var (
		state         = make(map[string]*CacheData)
		unixInt int64 = 1507351311
	)

	key := "1"
	testData := "test"

	WriteToState(key, testData, unixInt, state)

	expected := NewCacheData(testData, unixInt)
	actual := state[key]

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unexpected: %v - actual is: %v", expected, actual)
	}
}
