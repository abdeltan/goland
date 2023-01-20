package cours

import (
	"errors"
	"fmt"
)

func Division(nb1 int, nb2 int) (*int, error) {
	if nb2 == 0 {
		return nil, errors.New("Division par 0")
	}

	var calcul = nb1 / nb2

	return &calcul, nil
}

func ArgList(items ...string) {
	for id, item := range items {
		fmt.Println(id, item)
	}
}
