package puppy

import (
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
