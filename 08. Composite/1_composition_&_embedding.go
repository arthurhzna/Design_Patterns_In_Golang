package main

import "fmt"

type Engine struct{}

func (e Engine) Start() {
	fmt.Println("Engine Start (value)")
}

func (e *Engine) StartPtr() {
	fmt.Println("Engine Start (pointer)")
}

func main() {

	// ==================================================
	// 1. Named field
	// ==================================================
	type Car1 struct {
		Engine Engine
	}

	car1 := Car1{} // Engine is automatically initialized with zero value: Engine{}
	// actually:
	// car1 := Car1{Engine: Engine{}}

	car1.Engine.Start()

	// ==================================================
	// 2. Embedded field
	// ==================================================
	type Car2 struct {
		Engine
	}

	car2 := Car2{} // Engine is automatically initialized with zero value: Engine{}
	// actually:
	// car2 := Car2{Engine: Engine{}}

	car2.Start() // same as car2.Engine.Start()

	// ==================================================
	// 3. Named pointer field
	// ==================================================
	type Car3 struct {
		Engine *Engine
	}

	car3 := Car3{} // Engine == nil (NOT initialized yet)

	// must initialize manually:
	car3.Engine = &Engine{}

	car3.Engine.Start()

	// ==================================================
	// 4. Embedded pointer field
	// ==================================================
	type Car4 struct {
		*Engine
	}

	car4 := Car4{} // Engine == nil (NOT initialized yet)

	// must initialize manually:
	car4.Engine = &Engine{}

	car4.Start() // same as car4.Engine.Start()
}
