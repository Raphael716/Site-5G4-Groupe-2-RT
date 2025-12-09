package main

import (
    "fmt"
    "time"
    "sync" // Bonus pour faire propre
)

type Personne struct {
    Nom string
    Age int
}

// Fonction générique pour trouver le max dans une slice
// On passe une fonction "selector" pour savoir quoi comparer
func FindMax[T any](items []T, getValue func(T) int) T {
    maxItem := items[0]
    for _, item := range items {
        if getValue(item) > getValue(maxItem) {
            maxItem = item
        }
    }
    return maxItem
}

func main() {
    personnes := []Personne{
        {"Alice", 25},
        {"Bob", 62},
        {"Charlie", 22},
        {"Diana", 45},
    }

    // 1. Gestion de la concurrence (WaitGroup permet d'attendre proprement)
    var wg sync.WaitGroup

    fmt.Println("--- Lancement des traitements ---")
    for _, p := range personnes {
        wg.Add(1) // On dit qu'on ajoute une tâche
        
        // Fonction anonyme lancée en goroutine
        go func(perso Personne) {
            defer wg.Done() // On dit qu'on a fini à la fin de la fonction
            time.Sleep(500 * time.Millisecond)
            fmt.Printf("- Email envoyé à %s\n", perso.Nom)
        }(p)
    }

    // 2. Calcul du max pendant que les emails partent
    doyen := FindMax(personnes, func(p Personne) int {
        return p.Age
    })

    // Attendre que tous les emails soient envoyés
    wg.Wait()
    
    fmt.Println("--- Résultat ---")
    fmt.Printf("La personne la plus âgée est %s (%d ans).\n", doyen.Nom, doyen.Age)
}
