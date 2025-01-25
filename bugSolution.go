```go
func main() {
    var m = make(map[int]int)
    m[1] = 1
    
    value, ok := m[2] 
    if ok {
        fmt.Println("Value found:", value)
    } else {
        fmt.Println("Key not found")
    }
    
    delete(m,2)
    value, ok = m[2]
    if ok {
        fmt.Println("Value found:", value)
    } else {
        fmt.Println("Key not found")
    }
    
    var m2 = make(map[string]int)
    m2["a"] = 1
    
    value2, ok2 := m2["b"]
    if ok2 {
        fmt.Println("Value found:", value2)
    } else {
        fmt.Println("Key not found")
    }
    
    delete(m2, "b")
    value2, ok2 = m2["b"]
    if ok2 {
        fmt.Println("Value found:", value2)
    } else {
        fmt.Println("Key not found")
    }
}
```