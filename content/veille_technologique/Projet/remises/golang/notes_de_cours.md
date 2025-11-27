+++
title = "Notes de cours"
weight = 2
+++

## Introduction

Ces notes de cours sur Golang (Go) s’appuient sur la documentation officielle et sur des tutoriels fiables. L’objectif est de comprendre les bases tout en gardant en tête les bonnes pratiques recommandées par la communauté. Go mise sur la simplicité, la compilation rapide et une concurrence efficace. Comme ce langage est très utilisé pour bâtir des services web, des outils réseau ou des pipelines de données, ces notes servent de référence pour les exercices, l’atelier et la veille technologique.

## Installer Go et organiser son environnement

Go est fourni sous forme d’un SDK qui inclut le compilateur, le linker et les outils nécessaires. Sur Linux, la documentation recommande d’utiliser soit le paquet officiel, soit l’archive tar.gz disponible sur go.dev (Source 1). L’installation standard configure un `GOROOT` propre, puis il suffit d’ajouter `GOPATH` au profil utilisateur pour stocker le cache et les exécutables. À partir de Go 1.20, il est possible de travailler uniquement avec les modules, mais garder `GOPATH` reste utile pour la compatibilité. La commande `go env` permet de vérifier rapidement les variables d’environnement en cas de doute. VS Code avec l’extension Go (Source 2) fournit l’autocomplétion, la navigation et la gestion des tests directement dans l’éditeur. Cette extension installe aussi des outils comme `gopls` et `dlv`, devenus indispensables dans un workflow moderne.

## Structure typique d’un projet

Un projet Go moderne repose sur les modules. À la racine, on place un fichier `go.mod` généré par la commande `go mod init monmodule`. Ce fichier indique le nom du module, la version de Go ciblée et les dépendances. La documentation souligne que chaque dossier peut devenir un module autonome, mais pour un travail scolaire il est conseillé d’adopter une structure simple avec un dossier `cmd/` pour les exécutables et un dossier `pkg/` pour le code réutilisable (Source 3). Les tests se placent dans les mêmes dossiers que le code, avec le suffixe `_test.go`. Go n’impose pas de hiérarchie stricte, mais il faut garder en tête qu’un package correspond à un dossier. Les fichiers `main.go` qui contiennent la fonction `main()` appartiennent au package `main`. Cette structure facilite la création de sous-commandes pour un outil CLI tout en séparant les fonctions réutilisables dans différents exercices.

## Typage statique et déclarations

Go est statiquement typé tout en restant ergonomique. Une variable peut être déclarée avec `var` en précisant le type, ou en laissant le compilateur l’inférer avec `:=`. Exemple : `var compteur int = 0` ou `compteur := 0`. Les types de base incluent `int`, `float64`, `bool`, `string`, mais aussi des types agrégés comme les tableaux fixes `[5]int` et les slices dynamiques `[]int`. La documentation rappelle que les slices sont des vues sur un tableau sous-jacent, donc ils partagent la même mémoire tant que la capacité n’est pas dépassée (Source 4). Les maps se déclarent avec `map[cleType]valType`. Lorsqu’une map est lue, `val, ok := maMap[cle]` permet de vérifier si la clé existe. Il est courant d’utiliser des alias de types (`type Celsius float64`) pour clarifier le code lorsqu’on manipule des unités ou des identifiants. 

## Constantes, iota et packages standards

Les constantes déclarées avec `const` sont évaluées au moment de la compilation. `iota` est pratique pour créer des enum simples. Par exemple :

```go
type Niveau int

const (
	Debutant Niveau = iota
	Intermediaire
	Avance
)
```

La documentation montre que `iota` se remet à zéro à chaque bloc `const`, ce qui évite d’écrire les valeurs à la main. Parmi les packages standards les plus utilisés, on retrouve `fmt` pour afficher, `os` pour lire les arguments ou l’environnement, `net/http` pour monter des serveurs web et `encoding/json` pour manipuler des données JSON (Source 5). Ces packages sont inclus de base, donc il n’est pas nécessaire d’ajouter des dépendances externes pour démarrer.

## Fonctions et multiples valeurs de retour

Les fonctions se déclarent avec le mot-clé `func`. Go autorise le retour de plusieurs valeurs, ce qui sert beaucoup pour l’erreur et le résultat. Exemple typique :

```go
func division(num, denom float64) (float64, error) {
	if denom == 0 {
		return 0, fmt.Errorf("division par zéro")
	}
	return num / denom, nil
}
```

Dans ce pattern, l’appelant vérifie `if err != nil`. Les fonctions peuvent aussi avoir des paramètres variadiques comme `func somme(nums ...int) int`. Go passe les arguments par valeur, mais comme les slices et maps contiennent des pointeurs internes, les mutations se propagent. Effective Go insiste pour garder les signatures simples et retourner tôt lorsqu’une erreur arrive (Source 6). Les fonctions anonymes (`func() {}`) sont utilisées pour les goroutines ou les closures.

## Structures, méthodes et pointeurs

Les `struct` permettent de créer des types composés. Pour ajouter des méthodes, on attache une fonction à un type via un receiver. Le receiver peut être une valeur ou un pointeur, mais la documentation conseille d’utiliser un pointeur dès qu’on modifie l’état ou si la struct est lourde à copier (Source 7). Exemple :

```go
type Compte struct {
	titulaire string
	solde     float64
}

func (c *Compte) Depot(montant float64) {
	c.solde += montant
}

func (c Compte) Solde() float64 {
	return c.solde
}
```

Les pointeurs en Go sont simples : on utilise `&` pour obtenir l’adresse et `*` pour dereferencer. Il n’y a pas d’arithmétique de pointeurs comme en C, ce qui évite plusieurs catégories de bugs. `new(Type)` renvoie un pointeur vers une instance zéro-initialisée, alors que `&Type{}` reste souvent plus lisible. Les structures zéro-valeurs sont généralement prêtes à l’emploi (par exemple un `bytes.Buffer` vide fonctionne sans initialisation supplémentaire), ce qui encourage un style sans constructeur.

## Interfaces et composition

Les interfaces décrivent un comportement en listant les signatures des méthodes attendues. L’implémentation est implicite : si un type possède toutes les méthodes, il satisfait l’interface sans mot-clé spécial. Exemple :

```go
type Lecteur interface {
	Lire() (string, error)
}

type Fichier struct { /* champs */ }

func (f *Fichier) Lire() (string, error) {
	// implémentation
}
```

Cette approche favorise la composition. On peut combiner plusieurs interfaces comme `type ReadWriteCloser interface { io.Reader; io.Writer; io.Closer }`. La documentation recommande de concevoir de petites interfaces ciblées (Source 6). Les interfaces vides `interface{}` restent parfois utiles, mais Go 1.18 a introduit le mot-clé `any`, plus lisible. Il faut également faire attention aux `nil` : une interface peut être non nulle mais contenir une valeur nulle, ce qui peut provoquer des surprises lors des comparaisons.

## Contrôle de flux et erreurs

Go propose des boucles `for` qui remplacent à la fois `while` et `do-while`. La forme la plus simple est `for condition {}`, mais il est aussi possible d’utiliser `for range`. Les `switch` n’ont pas besoin de `break`, et permettent de tester plusieurs valeurs ou expressions. Pour la gestion des erreurs, Go reste fidèle au pattern `if err != nil`. La documentation rappelle que les `panic` doivent être exceptionnels, plutôt réservés aux erreurs irréversibles (Source 1). Il existe aussi `defer`, qui exécute une fonction à la sortie du scope. Exemple classique : `defer fichier.Close()`. Les `defer` sont empilés et exécutés dans l’ordre inverse d’enregistrement.

## Concurrence avec goroutines

La particularité la plus marquante de Go, c’est la concurrence via les goroutines. Lorsqu’une instruction `go fonction()` est exécutée, Go démarre une goroutine légère. Pour synchroniser, on utilise les channels. Un channel typé `chan int` permet d’envoyer et de recevoir des entiers. Exemple :

```go
func produire(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func consommer(ch <-chan int) {
	for valeur := range ch {
		fmt.Println(valeur)
	}
}

func main() {
	ch := make(chan int)
	go produire(ch)
	consommer(ch)
}
```

Le canal est créé avec `make`. Un buffer (`make(chan int, 3)`) peut être spécifié pour réduire le blocage. Le mot-clé `select` permet d’attendre sur plusieurs canaux. La documentation officielle insiste sur le slogan « ne communiquez pas en partageant la mémoire, partagez la mémoire en communiquant » (Source 8). En pratique, certains cas nécessitent encore `sync.Mutex` ou `sync.WaitGroup`. Les `context.Context` sont aussi essentiels pour propager des délais ou des annulations.

## Génériques et fonctions paramétriques

Depuis Go 1.18, les génériques sont disponibles. On déclare un paramètre de type entre crochets. Exemple simple :

```go
func Map[T any](slice []T, fn func(T) T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}
```

Les contraintes se définissent en listant les opérations permises sur `T`. Il existe des contraintes prédéfinies dans `constraints` (Source 9). La documentation recommande de ne pas abuser des génériques afin de conserver la lisibilité. Ils sont particulièrement utiles pour écrire des helpers réutilisables ou pour éviter de dupliquer du code sur des slices de types différents.

## Gestion des dépendances avec Go Modules

Avec Go Modules, tout projet peut utiliser des dépendances versionnées. `go mod tidy` nettoie et ajoute automatiquement ce qui manque. Les dépendances sont stockées dans le cache global. La documentation mentionne aussi la directive `replace`, utile pour pointer vers un module local pendant le développement (Source 3). Pour vérifier l’intégrité, Go utilise des fichiers `go.sum` contenant les checksums. Lorsqu’un module est publié, il faut taguer une version (respect du semver) afin que les utilisateurs puissent choisir `module@v1.2.3`. Il est recommandé d’éviter d’imposer une version trop récente de Go tant que ce n’est pas nécessaire, pour faciliter la compilation sur différentes machines.

## Tests, benchmarks et exemples

Tester en Go consiste à écrire des fonctions qui commencent par `Test` et qui prennent `*testing.T`. Toute la suite de tests se lance avec `go test ./...`. Exemple :

```go
func TestDivision(t *testing.T) {
	res, err := division(6, 2)
	if err != nil {
		t.Fatalf("division a échoué: %v", err)
	}
	if res != 3 {
		t.Fatalf("résultat attendu 3, obtenu %f", res)
	}
}
```

Les benchmarks commencent par `Benchmark` et utilisent `*testing.B`. Les exemples commencent par `Example` et servent aussi à générer la documentation. La documentation conseille de garder les tests dans le même package afin d’accéder aux fonctions non exportées lorsqu’un contrôle supplémentaire est requis (Source 1). Pour consulter la couverture, la commande `go test -cover` est disponible. Avec VS Code, l’extension affiche la couverture directement dans les fichiers sources.

## Documentation et go doc

Commenter le code en Go suit un format précis. Les commentaires de documentation doivent commencer par le nom de l’élément commenté, comme `// Compte représente un compte bancaire`. En exécutant `go doc` ou en publiant sur pkg.go.dev, ces commentaires deviennent la référence officielle. La documentation recommande d’utiliser des phrases complètes et d’éviter de documenter les éléments qui sont déjà évidents par leur nom (Source 6). Les exemples inclus dans les fichiers `_test.go` s’affichent également dans la documentation générée, ce qui encourage la rédaction de snippets pertinents pour les futurs lecteurs.

## Gestion des entrées/sorties et sérialisation

`io` est le package de base pour l’input/output. Tout ce qui implémente `io.Reader` ou `io.Writer` fonctionne avec les fonctions génériques comme `io.Copy`. Pour le JSON, `encoding/json` fournit `Marshal` et `Unmarshal`. Il faut structurer ses structs avec des tags :

```go
type Utilisateur struct {
	Nom    string `json:"nom"`
	Courriel string `json:"courriel"`
}
```

La documentation insiste sur le fait que seules les valeurs exportées (nom en majuscule initiale) sont sérialisées (Source 1). Pour manipuler des fichiers, `os.Open` et `os.Create` renvoient un `*os.File` qui implémente aussi `io.Reader` et `io.Writer`. Le package `bufio` reste pratique pour lire ligne par ligne ou bufferiser les écritures.

## Contexte réseau et services web

Go est très utilisé côté serveur. Le package `net/http` permet de définir un gestionnaire avec `http.HandleFunc` ou d’implémenter l’interface `http.Handler`. Exemple minimal :

```go
func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Pour un projet plus poussé, des frameworks comme `chi` ou `gin` apportent des fonctionnalités supplémentaires, mais l’API standard reste très solide. Les middlewares se composent facilement. La documentation explique comment utiliser les contexts pour propager les délais ou les annulations (Source 10). Pour les requêtes sortantes, `http.Client` est configurable et il est recommandé de réutiliser le même client afin de profiter du pooling des connexions. La gestion des JSON se combine naturellement avec les handlers.

## Outils de formatage, lint et analyse

`gofmt` est l’outil officiel pour formatter. Aucun débat sur le style, puisque tout le monde se fie au format imposé (Source 6). `goimports` ajoute ou retire les imports automatiquement. Pour l’analyse statique, `go vet` détecte des patterns suspects, et des outils comme `staticcheck` vont plus loin. Dans VS Code, il est possible de configurer l’extension afin que ces outils se lancent à l’enregistrement. Beaucoup de développeurs exécutent `gofmt` et `go test` avant chaque commit pour éviter les mauvaises surprises. Le langage encourage la simplicité, donc lorsqu’un pattern devient trop compliqué, mieux vaut vérifier dans Effective Go si la convention est respectée.

## Déploiement et exécutions multi-plateformes

Compiler un binaire Go se fait avec `go build`. Les variables `GOOS` et `GOARCH` permettent de produire des binaires pour d’autres plateformes, par exemple `GOOS=windows GOARCH=amd64 go build`. La documentation précise que Go inclut un runtime léger avec un garbage collector, donc il n’est pas nécessaire d’installer une machine virtuelle sur la cible (Source 1). Pour gérer les variables d’environnement et la configuration, le package `os` peut être utilisé ou un fichier `env` peut être importé. Dans un contexte d’atelier, générer un binaire statique facilite le partage sans obliger les autres à installer Go.

## Étude de cas : mini API REST

Pour illustrer ces notions, un exemple classique consiste à développer une mini API REST en Go. Une telle API expose des routes pour créer et lister des tâches. La struct `Task` est sérialisée en JSON et les handlers utilisent `context.Context` pour gérer les délais. Une goroutine peut nettoyer les tâches expirées, tandis que `sync.RWMutex` protège les données partagées. Les tests couvrent les handlers en utilisant `httptest.NewRecorder`. Ce type de projet sert d’atelier pratique et permet de manipuler tout le workflow : `go mod tidy`, `go test`, `gofmt`, génération de documentation, etc.

## Conclusion

Go reste un langage facile à prendre en main tout en proposant des fonctionnalités puissantes comme les goroutines et les modules. Le contexte d’études collégiales en informatique met en valeur un écosystème qui pousse à écrire du code propre et structuré rapidement. La documentation officielle étant très claire, il devient plus simple de rester aligné avec les pratiques recommandées. Ces notes servent pour les travaux pratiques et pour la veille technologique portant sur les usages de Go dans l’industrie (microservices, cloud, outils DevOps). En s’appuyant sur des sources fiables comme celles listées ci-dessous, il est possible de construire une base solide pour progresser.

## Démonstration

Voici un jeux de blackjack créer pour la démonstration de ce qui est possible de faire avec Go Lang.
[Lien vers le repo](https://github.com/Raphael716/blackjack)

## Sources

1. [Go Documentation – Getting Started](https://go.dev/doc/) (consulté le 27 novembre 2025)
2. [Go for Visual Studio Code – VS Code Marketplace](https://marketplace.visualstudio.com/items?itemName=golang.go) (consulté le 27 novembre 2025)
3. [Go Modules Reference](https://go.dev/doc/modules/managing-dependencies) (consulté le 27 novembre 2025)
4. [Go Slices: Usage and Internals](https://go.dev/blog/slices-intro) (consulté le 27 novembre 2025)
5. [Standard Library Overview](https://pkg.go.dev/std) (consulté le 27 novembre 2025)
6. [Effective Go](https://go.dev/doc/effective_go) (consulté le 27 novembre 2025)
7. [Go FAQ – Methods and Interfaces](https://go.dev/doc/faq#methods) (consulté le 27 novembre 2025)
8. [Share Memory By Communicating – Go Codewalk](https://go.dev/doc/codewalk/sharemem/) (consulté le 27 novembre 2025)
9. [Type Parameters Proposal and Guide](https://go.dev/doc/tutorial/generics) (consulté le 27 novembre 2025)
10. [Go net/http Package Documentation](https://pkg.go.dev/net/http) (consulté le 27 novembre 2025)

