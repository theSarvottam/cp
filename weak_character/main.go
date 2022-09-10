package main

import (
	"fmt"
	"sort"
)

func main() {

	properties := [][]int{
		{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5},
	}
	count := numberOfWeakCharactersWithSorting(properties)
	fmt.Println(count)
}
func numberOfWeakCharacters(properties [][]int) int {
	result := 0
	weak_mapping := make([]int, len(properties))
	for i := 0; i < len(properties); i++ {
		if weak_mapping[i] != 1 {
			for j := i + 1; j < len(properties); j++ {
				if properties[i][0] == properties[j][0] || properties[i][1] == properties[j][1] {
					continue
				} else if weak_mapping[i] != 1 && properties[i][0] < properties[j][0] && properties[i][1] < properties[j][1] {
					weak_mapping[i] = 1
					result++
				} else if weak_mapping[j] != 1 && properties[i][0] > properties[j][0] && properties[i][1] > properties[j][1] {
					weak_mapping[j] = 1
					result++
				}
			}
		}
	}
	return result

}

func numberOfWeakCharactersWithSorting(properties [][]int) int {
	result := 0
	currentMaxDefence := 0

	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] > properties[j][1]

		}
		return properties[i][0] < properties[j][0]
	})

	for i := len(properties) - 1; i >= 0; i-- {
		if properties[i][1] < currentMaxDefence {
			result++
		} else if properties[i][1] > currentMaxDefence {
			currentMaxDefence = properties[i][1]
		}

	}
	return result

}
