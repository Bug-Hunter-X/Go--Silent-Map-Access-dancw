```go
func main() {
    var m map[string]int
    value, ok := m["a"]
    if ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found")
    }
}
```