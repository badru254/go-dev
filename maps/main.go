package main

import (
	"fmt"
	"sort"
)

func main() {

	states := make(map[string]string)

	fmt.Println(states)

	states["MOM"] = "Mombasa"
	states["KTL"] = "Kitale"
	states["ELD"] = "Eldoret"
	states["NBI"] = "Nairobi"

	fmt.Println(states)

	eldoret := states["ELD"]

	fmt.Println("Selected State is", eldoret)

	delete(states, "ELD")
	states["KSM"] = "Kisumu"

	fmt.Println(states)

	for k, v := range states {

		fmt.Printf("Code = %v, Name = %v\n", k, v)

	}

	keys := make([]string, len(states))

	i := 0

	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		fmt.Printf("Sorted : Code = %v, Name = %v\n", key, states[key])
	}

}
