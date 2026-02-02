package main

import "fmt"

// Multiple structs to embed
type Address struct {
  Street  string
  City    string
  Country string
}

func (a Address) FullAddress() string {
  return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

type ContactInfo struct {
  Email string
  Phone string
}

func (c ContactInfo) SendEmail(message string) {
  fmt.Printf("Sending to %s: %s\n", c.Email, message)
}

// Person embeds multiple structs
type Person struct {
  Name string
  Age  int
  Address
  ContactInfo
}

// Employee embeds Person (nested embedding)
type Employee struct {
  Person
  JobTitle   string
  Department string
}

func main() {
  fmt.Println("=== Multiple Embedding ===")

  person := Person{
    Name: "Alice",
    Age:  30,
    Address: Address{
      Street:  "123 Main St",
      City:    "New York",
      Country: "USA",
    },
    ContactInfo: ContactInfo{
      Email: "alice@example.com",
      Phone: "555-1234",
    },
  }

  // All fields are promoted
  fmt.Println("Name:", person.Name)
  fmt.Println("City:", person.City)     // From Address
  fmt.Println("Email:", person.Email)   // From ContactInfo
  fmt.Println("Country:", person.Country)

  // Methods are also promoted
  fmt.Println("\nFull address:", person.FullAddress())
  person.SendEmail("Hello!")

  fmt.Println("\n=== Nested Embedding ===")

  emp := Employee{
    Person: Person{
      Name: "Bob",
      Age:  35,
      Address: Address{
        City: "Boston",
        Country: "USA",
      },
      ContactInfo: ContactInfo{
        Email: "bob@company.com",
      },
    },
    JobTitle:   "Engineer",
    Department: "Development",
  }

  // Fields promoted through nested embedding
  fmt.Println("Name:", emp.Name)       // Person.Name
  fmt.Println("City:", emp.City)       // Person.Address.City
  fmt.Println("Email:", emp.Email)     // Person.ContactInfo.Email
  fmt.Println("Job:", emp.JobTitle)

  // Methods work too
  fmt.Println("\nEmployee address:", emp.FullAddress())
}
