+++
title = "Atelier"
weight = 3
+++

# Atelier Golang : Premiers pas et utilisation des génériques

Bienvenue dans cet atelier d'introduction au langage **Go** (Golang).
Cet atelier permet aux étudiants de découvrir le langage **Golang**, ses concepts de base et ses nouveautés récentes comme les **génériques** et les **goroutines**. L’objectif est de créer de petits programmes fonctionnels pour se familiariser avec le langage.

---

## 1) Objectifs

À la fin de ce laboratoire, l’étudiant sera capable de :

1. Installer Go et configurer l’environnement de développement.
2. Créer des programmes simples avec variables, boucles, fonctions et slices.
3. Utiliser les génériques pour factoriser du code.
4. Travailler avec maps et structs.
5. Expérimenter les goroutines et channels pour la concurrence.
6. Comprendre un mini-projet combinant ces concepts.

---

## 2) Prérequis

- VS Code ou un autre éditeur.
- Connaissance de base en programmation (variables, boucles, fonctions).
- Docker (optionnel mais recommandé pour un environnement isolé).

---

## 3) Installation et configuration

Choisissez **une** des deux méthodes ci-dessous.

### Option A : Installation Locale (Recommandée pour VS Code)

 1. Télécharger Go : [https://go.dev/dl/](https://go.dev/dl/).
 2. Ouvrez votre terminal et vérifiez l'installation :
```bash
    go version
```
 3. Installez l'extension **"Go"** officielle dans VS Code.
 4. Créez un dossier nommé `atelier-go` sur votre bureau.

### Option B : Environnement Docker (Si vous ne voulez rien installer)

Si vous préférez ne pas installer Go sur votre machine, nous allons utiliser un conteneur.

 1. Créez un dossier `atelier-go`.
 2. À l'intérieur, créez un fichier nommé `Dockerfile` :

```dockerfile
FROM golang:1.23-alpine
WORKDIR /app
# On installe nano pour pouvoir éditer dans le conteneur si besoin, 
# mais l'idéal est d'éditer en local via le volume.
RUN apk add --no-cache nano git
CMD ["sh"]
``` 

 * Construisez l'image et lancez le conteneur en liant votre dossier actuel :

```bash
# 1. Construction
docker build -t atelier-go .

# 2. Lancement (Windows PowerShell)
docker run --rm -it -v ${PWD}:/app atelier-go
```

### Initialisation du projet

Dans votre terminal (ou celui de Docker), initialisez le module Go. C'est la carte d'identité de votre projet.

```bash
go mod init atelier
```
Cela va créer un fichier `go.mod`.

## 4) Exercices

### Partie 1 : Les Fondamentaux
#### Exercice 1 : Hello World
 1. Créez un fichier main.go.
 2. Définissez le package main.
 3. Dans la fonction main, affichez "Bonjour le monde !".
 4. Lancez le programme avec go run main.go.

<details> <summary>=> <strong>Voir la solution</strong></summary>


```go
package main

import "fmt"

func main() {
    fmt.Println("Bonjour le monde !")
}
```

</details>

#### Exercice 2 : Variables et types

 1. Déclarez une variable `nom (string)` initialisée à votre prénom.
 2. Déclarez une variable `age (int)` initialisée à votre âge.
 3. Déclarez une variable `taille (float64)` avec la syntaxe courte `:=`.
 4. Affichez une phrase complète combinant ces variables.
    **Astuce :** Utilisez `:=` uniquement à l'intérieur des fonctions.

<details> <summary>=> <strong>Voir la solution</strong></summary>

```go
package main

import "fmt"

func main() {
    var nom string = "Alice"
    var age int = 25
    taille := 1.75 // Go devine que c'est un float64

    fmt.Printf("Je m'appelle %s, j'ai %d ans et je mesure %.2fm.\n", nom, age, taille)
}
```

</details>


#### Exercice 3 : Logique de contrôle

En Go, il n'y a que `for` (pas de `while`).
 1. Créez une boucle qui itère de 1 à 10.
 2. Si le nombre est pair, affichez "Pair".
 3. Sinon, affichez "Impair".

<details> <summary>=> <strong>Voir la solution</strong></summary>

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Printf("%d est Pair\n", i)
        } else {
            fmt.Printf("%d est Impair\n", i)
        }
    }
}
```

</details>

### Partie 2 : Structures de données
#### Exercice 4 : Slices et tableaux

Les tableaux ont une taille fixe, les slices sont dynamiques. C'est ce qu'on utilise 99% du temps.
 1. Créez un slice d'entiers contenant `[10, 20, 30, 40]`.
 2. Ajoutez la valeur 50 à ce slice (utilisez append).
 3. Affichez la longueur du slice (`len`).
 4. Affichez le slice complet.

<details> <summary>=> <strong>Voir la solution</strong></summary>

```go

package main
import "fmt"

func main() {
    // Création du slice
    nombres := []int{10, 20, 30, 40}

    // Ajout d'un élément (append retourne un nouveau slice)
    nombres = append(nombres, 50)

    fmt.Println("Contenu :", nombres)
    fmt.Println("Longueur :", len(nombres))
}
```

</details>

#### Exercice 5 : Structs et Maps

 1. Définissez une structure `Personne` avec les champs `Nom` (string) et `Age` (int).
 2. Dans le `main`, créez 3 variables de type `Personne`.
 3. Stockez-les dans une `map` où la **clé** est le nom (string) et la **valeur** est l'âge (int).
 4. Affichez la map.

 <details> <summary>=> <strong>Voir la solution</strong></summary>

```go

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
```

</details>

### Partie 3 : Fonctions avancées & Génériques
#### Exercice 6 : Introduction aux Génériques

 1. Écrivez une fonction classique `Addition(a, b int) int`.
 2. Écrivez une fonction générique `Max[T comparable](a, b T) T`.
    - Note : `comparable` est une contrainte qui permet d'utiliser `==` ou `!=`. Pour utiliser `>`, nous avons besoin de `any` et d'une logique spécifique ou du package constraints
    - **Simplification pour l'exercice :** Créez `Max[T int | float64](a, b T) T`. Cette syntaxe dit "T peut être un int OU un float64".
 3. Testez votre fonction `Max` avec des entiers puis avec des flottants.

<details> <summary>=> <strong>Voir la solution</strong></summary>

```go

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
```

</details>

### Partie 4 : Concurrence
#### Exercice 7 : Goroutines et channels

 1. Créez une fonction `worker(c chan string)` qui :
    - Attend 500ms (`time.Sleep`).
    - Envoie le message "Travail terminé" dans le canal `c`.
    - Répète cela 3 fois.
 2. Dans le `main` :
    - Créez un channel de string : `c := make(chan string)`.
    - Lancez la goroutine : `go worker(c)`.
    - Récupérez et affichez les messages reçus du canal.

<details> <summary>=> <strong>Voir la solution</strong></summary>

```go

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
```

</details>

#### Exercice 8 : Mini-projet combiné

**Objectif :**
 * Avoir une liste de `Personne` (Nom, Age).
 * Utiliser une **fonction générique** pour trouver la personne la plus âgée.
 * Lancer une **goroutine** pour chaque personne qui simule un "traitement" (ex: "Envoi d'email à Bob...") avec un petit délai.

**Instructions :**
 1. Définissez la struct `Personne`.
 2. Créez une slice avec 3-4 personnes.
 3. Lancez les goroutines d'envoi d'email (pas besoin de channel complexe, juste un `go func()` avec un `Printf` et un `Sleep` suffira, mais n'oubliez pas d'attendre la fin dans le main, par exemple avec un gros `time.Sleep` ou un `WaitGroup` si vous voulez).
 4. Trouvez et affichez le doyen de la liste.

<details> <summary>=> <strong>Voir le corrigé complet</strong></summary>

```go

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
            fmt.Printf("📧 Email envoyé à %s\n", perso.Nom)
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
```

</details>

