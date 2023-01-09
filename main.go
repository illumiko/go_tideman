package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

const (
	err_candidate_not_found = "Candidate not found"
	err_record_preference   = "Preference not recorded"
)

type candidate struct {
	name     string
	strength int
	source   bool
}

type global_ranks map[voter_name]voter_ranks

var global_ranking = global_ranks{}

var candidates = []string{"Miko", "Luk", "Inari"}
var pair_candidates = [][]string{}

func (gl global_ranks) sort_pairs() (margins map[string]int) {
	//compares each pairs with all votes of the candidates to find the margin
	for _, pairs := range pair_candidates {
		for _, v := range gl {
			fmt.Println(v, pairs, determine_margin(pairs, v))

		}
	}

	return
}

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
		true_candidate_check(voted, candidate)
	}
	v.votes = voted
	err := v.record_preferences()
	if err != nil {
		log.Fatalln(err, err_record_preference)
	}
	return nil
}

// }}}

//Helper functions

// Compares Voters list with pair_candidates to determine margin
func determine_margin(pairs []string, votes voter_ranks) (margin []int) {
	for _, candidate := range pairs {
		margin = append(margin, indexof(votes, candidate))
	}

	return margin
}

// Sorts the slice, and creates pairs from the passed slice without affecting the passed slice
func make_pairs(unpaired_slice []string) (pairs [][]string) {
	//Prevents changing the original slice
	copy_slice := make([]string, len(unpaired_slice))
	copy(copy_slice, unpaired_slice)

	//sorts
	sort.Strings(copy_slice)

	//makes pairs
	for i, v := range copy_slice {
		pairs_slice := copy_slice[i+1:]
		if len(pairs_slice) == 0 {
			break
		}
		for _, x := range pairs_slice {
			pairs = append(pairs, []string{v, x})
		}
	}

	return pairs
}

// Loops through vouter_ranks and checks if all the votes are for true candidates
func true_candidate_check(slice voter_ranks, find string) (truthy bool) {
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
func indexof(slice voter_ranks, find string) (index int) {
	maps := make(map[string]int)
	for i, value := range slice {
		maps[value] = i
	}
	index, truthy := maps[find]
	if truthy == false {
		log.Fatalln(err_candidate_not_found, find)
	}
	return index
}

func init() {
	voter1 := voters{"John", voter_ranks{}}
	err := voter1.vote(voter_ranks{"Luk", "Inari", "Miko"})
	if err != nil {
		log.Fatalln(err_candidate_not_found)
	}
	voter2 := voters{"Bob", voter_ranks{}}
	err = voter2.vote(voter_ranks{"Inari", "Miko", "Luk"})
	if err != nil {
		log.Fatalln(err_candidate_not_found)
	}

	voter3 := voters{"Doe", voter_ranks{}}
	err = voter3.vote(voter_ranks{"Inari", "Luk", "Miko"})
	if err != nil {
		log.Fatalln(err_candidate_not_found)
	}

	pair_candidates = make_pairs(candidates)

	global_ranking.sort_pairs()
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
