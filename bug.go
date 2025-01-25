```go
func main() {
    var m = make(map[int]int)
    m[1] = 1
    fmt.Println(m[2]) // This will print 0, not an error
    delete(m, 2)
    fmt.Println(m[2]) // This will also print 0, not an error
    
    var m2 = make(map[string]int)
    m2["a"] = 1
    fmt.Println(m2["b"]) // This will print 0, not an error
    delete(m2, "b")
    fmt.Println(m2["b"]) // This will also print 0, not an error
}
```