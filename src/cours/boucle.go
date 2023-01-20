package cours

import "fmt"

func While() {
	i := 0
	for {
		if i > 5 {
			break
		}

		// algo ici
		i++
	}
	fmt.Println(i)
}

func DoWhile() {
	i := 0
	for {
		// algo ici
		if i > 5 {
			break
		}
		i++
	}
	fmt.Println(i)
}

func For() {
	for i := 0; i < 5; i++ {
		// algo ici
	}
}

func ForEach(array []string, args ...string) {
	for index := range array {
		fmt.Println(index)
	}

	for index := range args {
		fmt.Println(index)
	}

	for _, value := range array {
		fmt.Println(value)
	}

	for _, value := range args {
		fmt.Println(value)
	}

	for index, value := range array {
		fmt.Println(index, value)
	}

	for index, value := range args {
		fmt.Println(index, value)
	}
}
