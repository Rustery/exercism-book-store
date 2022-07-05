package bookstore

import (
	"sort"

	"golang.org/x/exp/maps"
)

type Group map[int]int

const price int = 800

func getCostByCount(count int) int {
	d := Group{
		2: 5,
		3: 10,
		4: 20,
		5: 25,
	}
	return count * (price - price*d[count]/100)
}

func Cost(books []int) int {
	mainGroup := Group{}
	for _, book := range books {
		mainGroup[book]++
	}
	keys := maps.Keys(mainGroup)
	sort.Slice(keys, func(i, j int) bool {
		return mainGroup[keys[i]] > mainGroup[keys[j]]
	})
	groups := [][]Group{}
	for i := len(mainGroup); i > 0; i-- {
		group := []Group{}
		for _, book := range keys {
			qty := mainGroup[book]
			for ii := 0; ii < qty; ii++ {
				bookAdded := false
				for {
					for _, subGroup := range group {
						if len(subGroup) < i {
							_, ok := subGroup[book]
							if !ok {
								subGroup[book] = 1
								bookAdded = true
								break
							}
						}
					}
					if bookAdded {
						break
					}
					subGroup := Group{}
					group = append(group, subGroup)
				}
			}
		}
		groups = append(groups, group)
	}

	cost := 0
	for _, group := range groups {
		groupCost := 0
		for _, subGroup := range group {
			groupCost += getCostByCount(len(subGroup))
		}
		if groupCost < cost || cost == 0 {
			cost = groupCost
		}
	}
	return cost
}
