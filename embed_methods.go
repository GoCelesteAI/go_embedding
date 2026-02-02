package main

import "fmt"

// Base struct with methods
type Animal struct {
  Name string
}

func (a Animal) Speak() {
  fmt.Println(a.Name, "makes a sound")
}

func (a Animal) Sleep() {
  fmt.Println(a.Name, "is sleeping...")
}

// Dog embeds Animal and adds its own methods
type Dog struct {
  Animal
  Breed string
}

// Override the Speak method
func (d Dog) Speak() {
  fmt.Println(d.Name, "barks! Woof!")
}

// Add new method
func (d Dog) Fetch() {
  fmt.Println(d.Name, "fetches the ball")
}

// Cat embeds Animal
type Cat struct {
  Animal
}

// Override Speak
func (c Cat) Speak() {
  fmt.Println(c.Name, "meows! Meow!")
}

func main() {
  fmt.Println("=== Promoted Methods ===")

  dog := Dog{
    Animal: Animal{Name: "Buddy"},
    Breed:  "Labrador",
  }

  // Promoted method from Animal
  dog.Sleep() // Uses Animal's Sleep method

  // Overridden method
  dog.Speak() // Uses Dog's Speak method

  // Dog's own method
  dog.Fetch()

  // Access original method through embedded struct
  fmt.Println("\nCalling Animal.Speak() directly:")
  dog.Animal.Speak()

  fmt.Println("\n=== Cat Example ===")
  cat := Cat{
    Animal: Animal{Name: "Whiskers"},
  }

  cat.Sleep() // Promoted from Animal
  cat.Speak() // Cat's overridden method

  fmt.Println("\n=== Method Resolution ===")
  fmt.Println("Dog.Speak() -> Dog's method (overrides Animal)")
  fmt.Println("Dog.Sleep() -> Animal's method (promoted)")
  fmt.Println("Dog.Fetch() -> Dog's method (new)")
}
