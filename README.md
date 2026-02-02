# Go Embedding Tutorial

Code examples from [Go Lesson 18: Embedding](https://www.youtube.com/watch?v=YOUR_VIDEO_ID).

## Examples

### Embedding Basics (`embed_basics.go`)
Learn how to embed structs and access promoted fields.

```go
type Animal struct {
    Name string
    Age  int
}

type Dog struct {
    Animal  // Embedded - no field name
    Breed string
}

dog := Dog{Animal: Animal{Name: "Buddy", Age: 3}, Breed: "Golden Retriever"}
fmt.Println(dog.Name)  // Promoted field - access directly!
```

### Promoted Methods (`embed_methods.go`)
Methods are promoted and can be overridden.

```go
func (a Animal) Speak() { fmt.Println(a.Name, "makes a sound") }
func (a Animal) Sleep() { fmt.Println(a.Name, "is sleeping...") }

// Dog overrides Speak
func (d Dog) Speak() { fmt.Println(d.Name, "barks!") }

dog.Sleep()  // Uses Animal's Sleep (promoted)
dog.Speak()  // Uses Dog's Speak (overridden)
```

### Multiple Embedding (`embed_multiple.go`)
Embed multiple structs for composition.

```go
type Person struct {
    Name string
    Address      // Embedded
    ContactInfo  // Embedded
}

person.City   // From Address
person.Email  // From ContactInfo
```

### Embedding with Interfaces (`embed_interface.go`)
Embedded types satisfy interfaces, and you can embed interfaces in structs.

```go
type Robot struct {
    Name    string
    Speaker  // Embedded interface
}

robot := Robot{Name: "RoboDog", Speaker: dog}
robot.Speak()  // Calls dog's Speak
```

## Running the Examples

```bash
go run embed_basics.go
go run embed_methods.go
go run embed_multiple.go
go run embed_interface.go
```

## Key Takeaways

- **Embedding** includes a struct without a field name
- **Promoted fields** are accessed directly on the outer struct
- **Promoted methods** work the same way
- **Override** methods by defining them on the outer type
- Go favors **composition over inheritance**

## Subscribe

For more Go tutorials, subscribe to [CelesteAI](https://www.youtube.com/@GoCelesteAI)!
