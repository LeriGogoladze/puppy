package puppy

import (
	"fmt"

	"github.com/LeriGogoladze/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof Woof!!"
}

func BigBark() string {
	return dog.WhenGrowingUp("Bark()")
}

func BigBarks() string {
	return dog.WhenGrowingUp("Barks()")
}

func from13() {
	fmt.Println("Am from version 1.3.0")
}
