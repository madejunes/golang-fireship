package main

import "fmt"

func main() {
	firstReleases := map[string]int{
		"C": 1972,
		"C++": 1985,
		"Java": 1996,
		"Python": 1991,
		"Javascript": 1996,
		"Go": 2012,
	}

	for key, value := range firstReleases {
		fmt.Printf("%s first release in %d \n", key, value)
	}
}