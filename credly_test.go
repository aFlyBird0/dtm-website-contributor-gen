package main

import (
	"testing"
)

func TestHandleCredlyLink(t *testing.T) {

	const (
		originConst   = "https://www.credly.com/mgmt/organizations/19f0ed81-a1a5-4df4-bd76-e3d40e23c328/badges/earners/59c69146-58a7-471d-a908-bc5a0b7f5f6f/details"
		expectedConst = "https://www.credly.com/badges/59c69146-58a7-471d-a908-bc5a0b7f5f6f"
	)

	tables := []struct {
		origin   string
		expected string
	}{
		{
			origin:   originConst,
			expected: expectedConst,
		},
		{
			origin:   expectedConst,
			expected: expectedConst,
		},
		{
			origin:   originConst + "/",
			expected: expectedConst,
		},
		{
			origin:   "abc",
			expected: defaultBadgeLink,
		},
		{
			origin:   "",
			expected: defaultBadgeLink,
		},
	}

	for _, row := range tables {
		actual := handleCredlyLink(row.origin)
		if row.expected != actual {
			t.Errorf("origin: %s\nexpected:%s\nactual:%s\n", row.origin, row.expected, actual)
		}
	}
}
