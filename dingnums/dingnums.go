package main

import "fmt"

func main() {
    for char := '➊'; char < '➊' + 10; char++ {
        fmt.Printf("%c\n", char)
    }
}
