package main
import "fmt"

func main() {
    // Création du slice
    nombres := []int{10, 20, 30, 40}

    // Ajout d'un élément (append retourne un nouveau slice)
    nombres = append(nombres, 50)

    fmt.Println("Contenu :", nombres)
    fmt.Println("Longueur :", len(nombres))
    fmt.Println("Capacité :", cap(nombres))
}
