package main
import "fmt"

type Personne struct {
    Nom string
    Age int
}

func main() {
    p1 := Personne{"Alice", 25}
    p2 := Personne{Nom: "Bob", Age: 30}
    p3 := Personne{"Charlie", 22}

    // Création de la map
    annuaire := make(map[string]int)
    
    // Remplissage
    annuaire[p1.Nom] = p1.Age
    annuaire[p2.Nom] = p2.Age
    annuaire[p3.Nom] = p3.Age

    fmt.Println("Annuaire :", annuaire)
}