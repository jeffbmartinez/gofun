/*
Using interfaces and an anonymous field in a struct types to
simulate inheritance and polymorphism.
*/

package main

import (
  "fmt"
)

type Displayer interface {
  ShowInfo()
}

/*
The "parent" type, which implements the Displayer interface 
*/
type Shape struct {
  Color string
}

func (s Shape) ShowColor() {
  fmt.Printf("Color: %v\n", s.Color)
}

func (s Shape) ShowInfo() {
  s.ShowColor()
}

/*
The "child" type, which uses an anonymous field of type Shape.
Shape's fields and methods are now accessible as if they were
Circle's fields and methods.
*/
type Circle struct {
  Shape

  Radius int
}

func (c Circle) ShowRadius() {
  fmt.Printf("Radius: %v\n", c.Radius)
}

/*
This overrides the ShowInfo() method defined above with Shape as the receiver.
*/
func (c Circle) ShowInfo() {
  c.ShowColor()
  c.ShowRadius()
}


func displayThing(thing Displayer) {
  fmt.Println("/-------------")
  thing.ShowInfo()
  fmt.Println("\\-------------")
}

func main() {
  shape := &Shape{"red"}
  circle := &Circle{ *&Shape{"blue"}, 10}

  displayThing(shape)
  displayThing(circle)

  // We can still access the Shape's original overriden method, even though
  // this variable is of type Circle
  fmt.Println("")
  circle.Shape.ShowInfo()
  fmt.Println("")
}
