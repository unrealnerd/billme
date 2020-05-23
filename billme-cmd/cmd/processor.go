package cmd

import "fmt"

//Command ... 
type Command struct {
	Name    string
	Options []string
}

//Process ... 
func (c *Command) Process() {
	fmt.Println("processing", c.Name)
}
