package cours

import "fmt"

func worker(channel chan string, callback func() string) {
	channel <- callback()
}

func Master(callbacks ...func() string) {
	channel := make(chan string)
	for _, callback := range callbacks {
		go worker(channel, callback)
	}

	i := 0
	for result := range channel {
		fmt.Println(result)
		i++
		if i >= len(callbacks) {
			break
		}
	}

}
