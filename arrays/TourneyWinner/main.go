package main

import (
	"fmt"
	"sort"
)

type Competitor struct {
	Name  string
	Score int
}

// given an array of pairs of team that competed against each other,
// and the array containing results,
// write a function that returns the winner of the tourney.
func TourneyWinner(competitions [][]string, results []int) string {
	competitors := make(map[string]int)

	var innerJoin []string
	for i := range competitions {
		innerSlice := competitions[i]
		innerJoin = append(innerJoin, innerSlice...)
		copy(innerJoin, innerSlice[:])
	}
	for _, v := range innerJoin {
		competitors[v] = 0
	}

	for j := range competitions {
		if results[j] == 0 {
			competitor := competitions[j][1]
			competitors[competitor] += 1
		}
		if results[j] == 1 {
			competitor := competitions[j][0]
			competitors[competitor] += 1
		}
	}

	var tourneyArr []Competitor
	for k, v := range competitors {
		tourneyArr = append(tourneyArr, Competitor{Name: k, Score: v})
	}
	sort.Slice(tourneyArr, func(i, j int) bool {
		return tourneyArr[i].Score > tourneyArr[j].Score
	})

	return tourneyArr[0].Name
}

func main() {
	var competitions [][]string

	htmCSharp := []string{"HTML", "C#"}
	csharpPython := []string{"C#", "Python"}
	pyHtml := []string{"Python", "HTML"}

	results := []int{0, 0, 1}

	competitions = append(competitions, htmCSharp, csharpPython, pyHtml)

	winner := TourneyWinner(competitions, results)
	fmt.Println(winner)
}
