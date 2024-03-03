package main

import (
	"fmt"
	"strings"
)

func main() {

	studenRoleNumbers := map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Println(studenRoleNumbers)

	m := map[[2]int]string{}
	fmt.Println(m)

	// var x map[string]int
	// x["key"] = 10
	// fmt.Println(x)

	y := map[string]int{}
	y["key"] = 10
	fmt.Println(y)

	makeMap := make(map[string]int)

	makeMap = map[string]int{
		"a": 1,
		"c": 232,
	}

	makeMap["b"] = 10

	fmt.Println(makeMap, len(makeMap))

	for key, val := range makeMap {

		fmt.Println(key, val)
	}

	fmt.Println(strings.Repeat("===", 30))
	delete(makeMap, "b")
	fmt.Println(makeMap)
	fmt.Println(strings.Repeat("===", 30))
	fmt.Println(makeMap["b"])

	b, ok := makeMap["b"]
	fmt.Println(b, ok)
	c, ok := makeMap["c"]
	fmt.Println(c, ok)

	mp1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	mp2 := mp1

	delete(mp1, 1)

	fmt.Println(mp1)
	fmt.Println(mp2)

	// NESTING MAPS

	// https://www.boot.dev/assignments/94309851-18a5-433e-91db-788597fda6c1

	delete(mp1, 23344)

	names := []string{"billy", "bon", "shankha", "sinha", "joe"}

	names_map := getNameCounts(names)

	// Print the nested map
	for firstChar, namesMap := range names_map {
		fmt.Printf("%c: {\n", firstChar)
		for name, count := range namesMap {
			fmt.Printf("    %s: %d,\n", name, count)
		}
		fmt.Println("}")
	}

}

func getNameCounts(names []string) map[rune]map[string]int {

	// use the rune
	// because , in GO names[0] is a byte and it can be represented as RUNE !
	counts := make(map[rune]map[string]int)

	for _, name := range names {
		if len(name) != 0 {

			firstChar := rune(name[0])

			_, ok := counts[firstChar]

			if !ok {
				counts[firstChar] = make(map[string]int)
			}

			counts[firstChar][name]++

		}

	}

	return counts
}
