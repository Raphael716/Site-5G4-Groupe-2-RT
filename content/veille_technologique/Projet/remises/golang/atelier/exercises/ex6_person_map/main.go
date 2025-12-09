package main
import "fmt"

// Fonction classique
func Addition(a int, b int) int {
    return a + b
}

// Fonction générique avec contrainte d'union
// T peut être int OU float64
func Max[T int | float64](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println("Addition :", Addition(5, 7))
    
    fmt.Println("Max Int :", Max(10, 5))
    fmt.Println("Max Float :", Max(10.5, 5.2))
}
