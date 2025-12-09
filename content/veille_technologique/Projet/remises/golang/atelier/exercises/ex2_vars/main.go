package main

import "fmt"

func main() {
    var nom string = "Alice"
    var age int = 25
    taille := 1.75 // Go devine que c'est un float64

    fmt.Printf("Je m'appelle %s, j'ai %d ans et je mesure %.2fm.\n", nom, age, taille)
}
