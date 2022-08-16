package main

import "fmt"

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

	fmt.Println(competitors)

	return ""
}

func main() {
	var competitions [][]string

	htmCSharp := []string{"HTML", "C#"}
	csharpPython := []string{"C#", "Python"}
	pyHtml := []string{"Python", "HTML"}

	results := []int{0, 0, 1}

	competitions = append(competitions, htmCSharp, csharpPython, pyHtml)

	TourneyWinner(competitions, results)
}
