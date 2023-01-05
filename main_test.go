package main

import (
	"testing"
)

func Test_record_preferences(t *testing.T) {
	voter1 := voters{"john", voter_ranks{}}
	err := voter1.vote(voter_ranks{"Miko", "Inari", "Luk"})
	if err == nil {
		t.Error(err_candidate_not_found)
	}
	// voter2 := voters{"Bob", voter_ranks{}}
	// voter3 := voters{"Doe", voter_ranks{}}
	// voter4 := voters{"Doe", voter_ranks{}}
	// voter5 := voters{"Doe", voter_ranks{}}
	//
}
