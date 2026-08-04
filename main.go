package main
import ("fmt"; "net/http")
func main() {
    http.HandleFunc("/", home)
    fmt.Println("Server ok")
}