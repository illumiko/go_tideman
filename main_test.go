package main

import (
	"reflect"
	"testing"
)

// # Helper
func expections(t testing.TB, expect, got interface{}, deep bool) {
	t.Helper()
	if !deep {
		if got != expect {
			t.Errorf("Got: %v \nWant: %v", got, expect)
		}
	} else {
		if reflect.DeepEqual(expect, got) {
			t.Errorf("Got: %v \nWant: %v", got, expect)
		}
	}
}

func Test_record_preferences_and_vote(t *testing.T) { // {{{
	voter1 := voters{"John", voter_ranks{}}
	err := voter1.vote(voter_ranks{"Miko", "Inari", "Luk"})
	expections(t, nil, err, false)
	voter2 := voters{"Bob", voter_ranks{}}
	err = voter2.vote(voter_ranks{"Inari", "Miko", "Luk"})
	expections(t, nil, err, false)
	voter3 := voters{"Doe", voter_ranks{}}
	err = voter3.vote(voter_ranks{"Miko", "Luk", "Inari"})
	expections(t, nil, err, false)
	expect := map[voter_name]voter_ranks{
		"John": {"Miko", "Inari", "Luk"},
		"Bob":  {"Inari", "Miko", "Luk"},
		"Doe":  {"Miko", "Luk", "Inari"},
	}
	expections(t, expect, global_ranking, true)
} // }}}
func Test_make_pairs(t *testing.T) {
	s := []string{"a", "b", "c"}
	expect := [][]string{{"a", "b"}, {"a", "c"}, {"b", "c"}}
	get := make_pairs(s)
	expections(t, expect, get, true)
}
