package main

import "fmt"

// Base struct
type Animal struct {
  Name string
  Age  int
}

// Dog embeds Animal
type Dog struct {
  Animal // Embedded struct (no field name)
  Breed  string
}

// Cat also embeds Animal
type Cat struct {
  Animal
  Indoor bool
}

func main() {
  fmt.Println("=== Embedding Basics ===")

  // Create a Dog with embedded Animal
  dog := Dog{
    Animal: Animal{Name: "Buddy", Age: 3},
    Breed:  "Golden Retriever",
  }

  // Promoted fields - access directly!
  fmt.Println("\nDog info:")
  fmt.Println("  Name:", dog.Name)   // Promoted from Animal
  fmt.Println("  Age:", dog.Age)     // Promoted from Animal
  fmt.Println("  Breed:", dog.Breed) // Dog's own field

  // Can also access through embedded struct
  fmt.Println("\nAccess via Animal:")
  fmt.Println("  dog.Animal.Name:", dog.Animal.Name)

  // Create a Cat
  cat := Cat{
    Animal: Animal{Name: "Whiskers", Age: 5},
    Indoor: true,
  }

  fmt.Println("\nCat info:")
  fmt.Println("  Name:", cat.Name)
  fmt.Println("  Age:", cat.Age)
  fmt.Println("  Indoor:", cat.Indoor)

  // Modify promoted fields
  fmt.Println("\n=== Modifying Fields ===")
  dog.Name = "Max"
  dog.Age = 4
  fmt.Println("Updated dog:", dog.Name, "age", dog.Age)
}
