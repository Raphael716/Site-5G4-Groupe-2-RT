# Atelier Golang - Premiers pas

Ce dossier contient un laboratoire prêt à l'emploi pour apprendre les bases de Go, les génériques et la concurrence.

### Objectifs
- Installation & prise en main (local / Docker)
- Exercices progressifs avec corrigés
- Bonnes pratiques pour tests et exécution

### Prérequis
- Docker (optionnel si vous installez Go localement)
- (Optionnel) VS Code et l'extension Go

### Structure
- `Dockerfile` : image pour travailler sans installer Go localement
- `go.mod` : module pour l'atelier
- `exercises/` : exercices (fichiers prêts à lancer)
- `solutions/` : corrigés pour chaque exercice

### Utilisation rapide (local)
```bash
# depuis ce dossier
cd atelier
# exécuter un exercice, par ex. exercice 1
go run ./exercises/ex1_hello/main.go
```

Utilisation via Docker
```bash
# construire l'image (depuis la racine du dossier atelier)
docker build -t atelier-go .

# lancer le conteneur
# Windows (PowerShell) :
docker run --rm -it -v ${PWD}:/workspace -w /workspace atelier-go

# dans le conteneur, exécutez par exemple
# go run ./exercises/ex1_hello/main.go
```

Conseils pédagogiques
- Encourager l'utilisation de `go run` puis `go test` pour vérifier le comportement
- Montrer `go fmt` / `gofmt` et `go vet` rapidement

Images
- Vous pouvez utiliser l'image du projet située dans `static/images/logo5G4.png` pour illustrer la page de l'atelier.

---

Exercices inclus
1. `ex1_hello` - Hello World - go run ./exercises/ex1_hello/main.go
2. `ex2_vars` - Variables et types - go run ./exercises/ex2_vars/main.go
3. `ex3_for` - Boucles et condition - go run ./exercises/ex3_for/main.go
4. `ex4_slices` - Slices - go run ./exercises/ex4_slices/main.go
5. `ex5_add_max` - fonctions et génériques - go run ./exercises/ex5_add_max/main.go
6. `ex6_person_map` - structs et maps - go run ./exercises/ex6_person_map/main.go
7. `ex7_worker` - goroutines & channels - go run ./exercises/ex7_worker/main.go
8. `ex8_mini` - mini-projet combiné - go run ./exercises/ex8_mini/main.go
