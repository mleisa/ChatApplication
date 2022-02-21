package test

import "fmt"

type Animaler interface {
	Eat()
	Speak()
	Error()
}

type SuperAnimals struct {
	locomotion string
}

type Animals struct {
	SuperAnimals
	food  string
	sound string
}

func (x Animals) Eat() {
	fmt.Println(x.food)
}

func (x Animals) Speak() {
	fmt.Println(x.sound)
}

func (x Animals) Error() {
	fmt.Println("Invalid query entered!")
}

func main() {
	m := map[string]Animals{
		"cow":   Animals{SuperAnimals{"walk"}, "grass", "moo"},
		"Cow":   Animals{SuperAnimals{"walk"}, "grass", "moo"},
		"Bird":  Animals{SuperAnimals{"fly"}, "worms", "peep"},
		"bird":  Animals{SuperAnimals{"fly"}, "worms", "peep"},
		"Snake": Animals{SuperAnimals{"slither"}, "mice", "hsss"},
		"snake": Animals{SuperAnimals{"slither"}, "mice", "hsss"},
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Enter animal name & query (eat / speak): ")
		fmt.Print(">")
		var animal, op string
		fmt.Scan(&animal)
		fmt.Print(">")
		fmt.Scan(&op)

		if op == "eat" {
			m[animal].Eat()
		} else if op == "speak" {
			m[animal].Speak()
		} else {
			m[animal].Error()
		}
	}

}
