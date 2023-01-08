package main

import (
	"errors"
	"log"
)

const (
	err_candidate_not_found = "Candidate not found"
)

type candidate struct {
	name     string
	strength int
	source   bool
}

func make_pairs(unpaired_slice []string) (pairs [][]string) {
	for i, v := range unpaired_slice {
		pairs_slice := unpaired_slice[i+1:]
		if len(pairs_slice) == 0 {
			break
		}
		for _, x := range pairs_slice {
			pairs = append(pairs, []string{v, x})
		}
	}

	return pairs
}

type global_ranks map[voter_name]voter_ranks

var global_ranking = global_ranks{}

//	var candidates = [3]candidate{
//		{name: "Miko", strength: 0, source: false},
//		{name: "Luk", strength: 0, source: false},
//		{name: "Inari", strength: 0, source: false},
//	}
var candidates = [...]string{"Miko", "Luk", "Inari"}

type voter_ranks []string // {{{
type voter_name string

type voters struct {
	name  voter_name
	votes voter_ranks
}

// Voters struct has access to all the candiates voter voted for, it checks if their is an entry for
// for the voter in the global_rankings, if not - then adds the voters list to the global ranking.
// Otherwise returns an error
func (v voters) record_preferences() error {
	_, voted := global_ranking[v.name]
	if voted {
		return errors.New("Voter voted")
	}
	global_ranking[v.name] = v.votes
	return nil
}

// Takes a [3]string each value corresponding to a candidate running for election,
// checks if the voted candidate is running in the election, if not - returns an error,
// otherwise adds voted to v.voters and records the it a global map to keep strack of it
func (v *voters) vote(voted voter_ranks) error {
	for _, candidate := range candidates {
		true_candidate_check(candidate, voted)
	}
	v.votes = voted
	v.record_preferences()
	return nil
}

// }}}

//Helper functions

// Loops through vouter_ranks and checks if all the votes are for true candidates
func true_candidate_check(find string, slice voter_ranks) (truthy bool) {
	maps := make(map[string]bool)
	for _, value := range slice {
		maps[value] = false
	}
	_, truthy = maps[find]
	if truthy == false {
		log.Fatalln(err_candidate_not_found, find)
	}
	return truthy
}

func main() {
}

/*

xx//Complete the vote function.

The function takes arguments rank, name, and ranks. If name is a match for the name of a
valid candidate, then you should update the ranks array to indicate that the voter has
the candidate as their rank preference (where 0 is
the first preference, 1 is the second preference, etc.)
Recall that ranks[i] here represents the user’s ith preference. The function should return true if
the rank was successfully recorded, and false otherwise (if, for instance, name is not the name of one
of the candidates).
You may assume that no two candidates will have the same name.

xx//Complete the record_preferences function.

The function is called once for each voter, and takes as argument the ranks array,
(recall that ranks[i] is the voter’s ith preference, where ranks[0] is the first preference).
The function should update the global preferences array to add the current voter’s preferences. Recall that preferences[i][j] should represent the number of voters who prefer candidate i over candidate j.
You may assume that every voter will rank each of the candidates.

//Complete the add_pairs function.

The function should add all pairs of candidates where one candidate is preferred to the pairs array.
A pair of candidates who are tied (one is not preferred over the other) should not be added to the array.
The function should update the global variable pair_count to be the number of pairs of candidates.
(The pairs should thus all be stored between pairs[0] and pairs[pair_count - 1], inclusive).

//Complete the sort_pairs function.

The function should sort the pairs array in decreasing order of strength of victory, where strength of
victory is defined to be the number of voters who prefer the preferred candidate. If multiple pairs have
the same strength of victory, you may assume that the order does not matter.

//Complete the lock_pairs function.

The function should create the locked graph, adding all edges in decreasing order of victory strength
so long as the edge would not create a cycle.

//Complete the print_winner function.

The function should print out the name of the candidate who is the source of the graph. You may
assume there will not be more than one source.
You should not modify anything else in tideman.c other than the implementations of the vote,
record_preferences, add_pairs, sort_pairs, lock_pairs, and print_winner functions (and the inclusion of
additional header files, if you’d like). You are permitted to add additional functions to tideman.c, so
long as you do not change the declarations of any of the existing functions.
*/
