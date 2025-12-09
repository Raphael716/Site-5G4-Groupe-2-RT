package main

import (
    "fmt"
    "time"
)

func worker(c chan string) {
    for i := 1; i <= 3; i++ {
        time.Sleep(500 * time.Millisecond)
        // Envoi dans le canal
        c <- fmt.Sprintf("Tâche %d terminée", i)
    }
    // On ferme le canal pour dire qu'on a fini
    close(c)
}

func main() {
    c := make(chan string)
    
    // Lancement en parallèle
    go worker(c)

    fmt.Println("En attente du worker...")

    // Lecture du canal tant qu'il est ouvert
    for msg := range c {
        fmt.Println("Reçu :", msg)
    }
    fmt.Println("Tout est fini !")
}
