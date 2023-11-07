package main

import "fmt"

type Instance struct {
	name       string
	isLaunched bool
}

func Launch() *Instance {
	newInstance := new(Instance)
	newInstance.name = "My cool instance"
	newInstance.isLaunched = true
	return newInstance
}

func Stop(instancePointer *Instance) {
	instancePointer.isLaunched = false
}

func main() {
	instance := Launch()
	fmt.Println("Not launchedinstance.isLaunched: ", instance.isLaunched)

	Stop(instance)
	fmt.Println("Stopped! instance.isLaunched: ", instance.isLaunched)
}
