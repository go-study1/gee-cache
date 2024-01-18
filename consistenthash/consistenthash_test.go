package consistenthash

import (
	"strconv"
	"testing"
)

func TestHash(t *testing.T) {
	hash := New(3, func(Data []byte) uint32 {
		i, _ := strconv.Atoi(string(Data))
		return uint32(i)
	})
	hash.Add("6", "4", "2")
	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"25": "6",
		"27": "2",
	}

	for key, v := range testCases {
		if hash.Get(key) != v {
			t.Errorf("Asking for %s, should have yielded %s", key, v)
		}
	}
}
