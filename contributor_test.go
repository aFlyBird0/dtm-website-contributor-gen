package main

import "testing"

func TestRemoveBracket(t *testing.T) {
	if removeBracket("wzymumon(mumon)") != "wzymumon" {
		t.Error("wzymumon(mumon)")
	}
	if removeBracket("wzymumon（mumon）") != "wzymumon" {
		t.Error("wzymumon（mumon）")
	}
	if removeBracket("wzymumon") != "wzymumon" {
		t.Error("wzymumon")
	}
}
