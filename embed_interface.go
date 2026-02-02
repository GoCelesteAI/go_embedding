package main

import "fmt"

// Interface
type Speaker interface {
  Speak()
}

// Base with method
type Animal struct {
  Name string
}

func (a Animal) Speak() {
  fmt.Println(a.Name, "makes a sound")
}

// Dog embeds Animal
type Dog struct {
  Animal
}

func (d Dog) Speak() {
  fmt.Println(d.Name, "barks!")
}

// Cat embeds Animal
type Cat struct {
  Animal
}

func (c Cat) Speak() {
  fmt.Println(c.Name, "meows!")
}

// Function accepting interface
func MakeNoise(s Speaker) {
  s.Speak()
}

// Struct embedding an interface
type Robot struct {
  Name    string
  Speaker // Embedded interface
}

func main() {
  fmt.Println("=== Embedding and Interfaces ===")

  dog := Dog{Animal{Name: "Buddy"}}
  cat := Cat{Animal{Name: "Whiskers"}}

  // Both satisfy Speaker interface
  MakeNoise(dog)
  MakeNoise(cat)

  // Slice of interfaces
  fmt.Println("\n=== Polymorphism ===")
  speakers := []Speaker{dog, cat}
  for _, s := range speakers {
    s.Speak()
  }

  fmt.Println("\n=== Embedding Interface in Struct ===")

  // Robot embeds the Speaker interface
  robot := Robot{
    Name:    "RoboDog",
    Speaker: dog,
  }

  fmt.Println("Robot name:", robot.Name)
  robot.Speak() // Calls the embedded Speaker

  // Change the embedded speaker
  robot.Speaker = cat
  fmt.Println("\nAfter changing speaker:")
  robot.Speak() // Now calls cat's Speak

  fmt.Println("\n=== Composition over Inheritance ===")
  fmt.Println("Go uses embedding for composition")
  fmt.Println("No class hierarchy, just struct composition")
  fmt.Println("More flexible than traditional inheritance")
}
