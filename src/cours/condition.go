package cours

import "fmt"

func If(nb int) {

	if nb == 0 {
		fmt.Println(nb)
	}

}

func IfElse(nb int) bool {
	if nb == 0 {
		return true
	}

	return false
}

func SwitchCase(name string) {

	switch name {
	case "Florent":
		fmt.Println("Un admin")
	default:
		fmt.Println("un user")
	}

}
