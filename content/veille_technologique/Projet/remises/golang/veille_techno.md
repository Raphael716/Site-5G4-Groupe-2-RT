+++
title = "Veille technologique"
weight = 1
+++

Cette page est un modèle et des instructions pour réaliser votre veille technologique sur Golang.

 
## 1) Premier prompt

Indiquez ici le prompt exact que vous avez utilisé avec le LLM, et précisez le modèle/version.

Prompt à envoyer au LLM :

```text
Quelles fonctionnalités récentes ont fondamentalement changé la façon d'écrire du code Go, et cela a-t-il rendu le langage plus complexe ?
```

Demandé à ChatGPT, modèle GPT-5.1 (dernière mise à jour publique du 12 novembre 2025)

### Réponse du LLM

Voici la réponse qui m'a été donné par ChatGPT :

Super bonne question — il y a eu plusieurs évolutions récentes du langage Go qui ont vraiment changé la façon d’écrire du code. Voici les plus importantes, ce qu’elles apportent concrètement, et si — globalement — elles rendent Go plus complexe.

#### 1) Les génériques (Go 1.18)

Ce changement est probablement le plus fondamental des dernières années : Go a ajouté les type parameters (génériques) qui permettent d’écrire des fonctions et types réutilisables pour plusieurs types concrets, ainsi que la notion d’interface comme ensemble de types (type sets) et l’inférence de paramètres.
Impact : on peut factoriser énormément de code répétitif (collections, wrappers, fonctions utilitaires) sans recourir à `interface{}` + assertions. Mais il faut apprendre la syntaxe des paramètres de type, les contraintes (`constraints`) et les subtilités (ex. méthodes sur types paramétrés). 
Go
+1

Ex. très court (illustratif) :

```go
// fonction générique qui trouve le maximum selon < comparator >
func Max[T any](a, b T, cmp func(T,T) bool) T {
    if cmp(a,b) { return a }
    return b
}
```

#### 2) Fuzzing natif dans l’outil de test (Go 1.18)
Go a intégré le fuzz testing directement dans `go test`. Ça change la façon d’écrire des tests : on peut maintenant ajouter des fuzz tests qui cherchent automatiquement des cas limites pour vos fonctions. Utile pour la robustesse et la sécurité. Go

#### 3) Améliorations de l’outillage / runtime (p.ex. Go 1.20 / 1.21)
Les versions récentes ont apporté plusieurs améliorations à la toolchain : PGO (profile-guided optimizations) expérimental, optimisations du compilateur, gains de performance, et amélioration du support de rétrocompatibilité dans le go tool. Ces changements influencent surtout comment on compile, profile et déploie — pas la syntaxe, mais le workflow dev → build → prod.

#### 4) `go.work`, modules et flux de travail (workspace)
Le support go.work (et la stabilisation de modules) a clarifié les layouts multi-module et le développement multi-repo. Résultat : workflows monorepo/multi-module plus simples, meilleure ergonomie pour gros projets. Cela change comment on organise le code et les dépendances

#### 5) Propositions / débats sur la gestion des erreurs (et statut)

Il y a eu beaucoup de discussions (et propositions) pour simplifier la gestion d’erreurs (ex. try/handle), mais l’équipe Go a récemment déclaré qu’elle n’avancerait pas de changements syntactiques pour l’error handling pour l’instant. Autrement dit : on reste avec le modèle d’erreurs-valeures (if err != nil) mais le débat continue et certaines améliorations non-syntaxiques (librairies, patterns) se répandent.

---

##### Est-ce que ça rend Go plus complexe ?
Court : oui et non.

 -	Plus d’outils / concepts = plus de choses à apprendre. Les génériques introduisent de nouveaux concepts (type parameters, contraintes, type sets) qui augmentent la surface cognitive surtout pour les débutants. Le fuzzing et les options de toolchain demandent aussi un peu d’apprentissage pour être utilisés efficacement.

 - Mais le but était la puissance sans complexifier la surface d’usage quotidienne. Les auteurs et la communauté encouragent à ne pas abuser des génériques — utiliser des fonctions/tp génériques pour dédupliquer du boilerplate plutôt que pour tout abstraire. Beaucoup de code Go existant reste inchangé et lisible. En pratique : pour du code « application », la complexité supplémentaire est gérable ; pour des bibliothèques génériques, la capacité nouvellement offerte est puissante mais demande rigueur.

 Règles pratiques

 - N’utiliser les génériques que pour réduire du boilerplate ou quand vous gagnez en sécurité de type.
 - Préférer des API simples et compréhensibles — la lisibilité est une priorité en Go.
 - Profiter du fuzzing pour les fonctions critiques (parsing, sérialisation, validation).
 - Suivre les releases (`go1.x` release notes) pour les changements d’outillage/perf.

 ---

##### En résumé (TL;DR)

 - Les génériques (Go 1.18) ont le plus fort impact : plus d’expressivité, possibilité de réduire le code répétitif, mais une nouvelle surface à maîtriser. 
 - Le fuzzing natif et les améliorations toolchain changent le workflow tests/build. 
 - L’équipe Go a refusé d’aller de l’avant avec des changements syntaxiques d’erreurs récemment, le style traditionnel if err != nil reste central, du moins pour l’instant. 
Si tu veux, je peux :
 - te montrer 3 exemples concrets (avant/après) d’un petit module transformé avec génériques,
 - ou te fournir un guide court (« quand utiliser les génériques ») adapté à ton niveau. Lequel tu préfères ?

## 2) Discussion de la réponse du LLM

### 2.1 La réponse est-elle complète ?

La réponse capture bien les changements majeurs des changements dans le language GoLang, mais elle manque certaines nouveautés très récentes (2024-2025), notamment :

 - **Go 1.22 (février 2024)** : correction de la capture de variables dans les boucles `for`
 - **Go 1.23 (août 2024)** : itérateurs (`range-over-func` : itérateurs personnalisés permettant d'itérer sur des séquences définies par une fonction), télémétrie, améliorations diverses

#### Points forts :
   - Bonne mise en contexte
   - Explications claires
   - Mention du fuzzing intégré (tests qui génèrent automatiquement des entrées pour trouver des cas limites)
   - Discussion équilibrée sur la complexité

#### Faiblesses importantes :
   * **Absence des versions récentes :** Aucune mention de Go 1.22 (février 2024) et Go 1.23 (août 2024)
   * **Changements syntaxiques critiques :** La correction de la capture des variables dans les boucles for (Go 1.22)
   * **Manque d'actualité sur l'écosystème :** Évolutions récentes des outils (gopls, govulncheck)
   * **Perspective d'apprentissage absente :** Comment ces changements affectent l'apprentissage et le codage

Donc, il manque une vision complète des changements de 2024-2025.

### 2.2 Comparaison : réponse de ChatGPT vs réponse de Github Copilot

J'ai posé la même question à GitHub Copilot et voici mon analyse comparative

GitHub Copilot fournit une réponse plus structurée et pratique que ChatGPT. Son approche est davantage orientée "développement" avec des recommandations immédiatement applicables, tandis que ChatGPT reste plus explicatif et théorique.

**Tableau comparatif :**

|**Aspect**|**ChatGPT (GPT-5.1)**|**GitHub Copilot**|
|---|---|---|
|Actualité|S'arrête vers 2023|Inclut Go 1.22+|
|Structure|Narrative, conversationnelle|Sections claires, format consultatif|
|Orientation|Explications théoriques|Recommandations pratiques|
|Exemples|Généraux|Plus concrets et courts|
|Pédagogie|Un peu théorique|Plus adapté à la pratique|
|Complexité|Discussion nuancée|Approche pragmatique|

**Conclusion :** Pour un travail scolaire, Copilot donne quelque chose de plus exploitable directement, car il est plus structuré et plus proche de l'expérience de developpement.


### 2.3 LLM comparé aux moteurs de recherche

Pour évaluer la qualité de la veille réalisée avec les LLM, j'ai comparé leurs réponses avec les résultats de 2 recherches Google ciblée sur les mêmes thématiques.

Voici les requêtes qui ont été effectuées pour valider et pour comparer avec les LLM :

 - **Requête 1** : "new functionality that changed golang"
 - **Requête 2** : "Is Golang losing its simplicity in recent years"

Et ensuite fait une analyse des 5 premiers résutats de chaque requête. 

Les résultats montrent une différence d'efficacité. Les IA ont réussi à résumé l'information immédiatement. Elles ont bien expliqué les changements techniques en filtrant les détails inutiles. À l'inverse, les résultats de Google sont très mélangés, on y trouve bien la documentation officielle, mais elle est parfois perdue au milieu d'articles de blog aux titres exagérés ou négatifs (ex: « Est-ce que Go est mort ? »), ce qui rend l'analyse plus difficile.

Le tableau ci-dessous résume les points forts et faibles observés pour chaque méthode.

|**Critère**|**IA (ChatGPT / Copilot)**|**Recherche Google**|
|---|---|---|
|Synthèse|Très bonne. Résume tout de suite les points clés.|Faible. Oblige à lire et trier plusieurs sites soi-même.|
|Qualité du code|Utile. Donne des exemples "avant/après" clairs.|Variable. Dépend du site (très bon sur le site officiel, aléatoire sur les blogs).|
|Neutralité|Équilibrée. Explique le pour et le contre calmement.|Biaisée. Affiche souvent des articles "clickbait" en premier.|
|Rapidité|Rapide. Une réponse claire en moins de 2 minutes.|Lente. Demande 15 à 20 minutes de lecture.|


**En conclusion,** les LLM sont excellents pour commencer une veille, mais Google reste indispensable pour vérifier les informations sur les sources officielles.


### 2.4 Y a-t-il une page web qui contient toutes ces informations ou suffisamment proche ?

En recherchant, **non**, aucune page web unique ne rassemble à la fois l'historique technique complet, les exemples de code comparatifs et l'analyse critique sur la complexité avec la même concision que le résumé généré par l'IA. L'information est fragmentée.
Cependant, nous avons identifié deux sources officielles qui, mises bout à bout, contiennent la totalité de la vérité technique et factuelle citée précédemment.

#### 1. La source technique officielle
**Source** : Go 1.23 Release Notes (Août 2024)
 - **Lien** : https://go.dev/doc/go1.23
 - **Contenu** : C'est le document de référence pour l'année écoulée. Il officialise l'introduction des itérateurs (`range-over-func` itérateurs personnalisés) et l'activation de la télémétrie.
 - **Qualité de la source** : Maximale. Cette page liste exactement ce qui est implémenté dans le compilateur. Elle prouve que le langage continue d'évoluer vers plus d'expressivité tout en complexifiant légèrement sa syntaxe interne, validant l'analyse de l'IA.

#### 2. La source explicative (La justification du design)
**Source** : The Go Blog - "Range Over Function Types"
 - **Lien** : https://go.dev/blog/range-functions
 - **Contenu** : Cet article officiel de l'équipe Go explique le pourquoi des nouvelles boucles for. Il démontre comment standardiser les séquences personnalisées.
 - **Qualité de la source** : Très Haute. Elle est essentielle pour comprendre l'intention des créateurs. Là où l'IA résume le "comment", cette source explique la philosophie derrière le changement, permettant de nuancer le débat sur la "perte de simplicité" du langage.

#### Conclusion sur les sources
L'IA a correctement expliquer des informations éparpillées sur le site go.dev (entre les Release Notes 1.22/1.23 et le Go Blog). Bien qu'il n'existe pas de page unique "résumé 2025", la combinaison de ces deux liens officiels suffit à reconstituer l'intégralité des faits techniques avancés dans cette veille, avec une fiabilité supérieure à n'importe quel résultat d'autres recherches.


### 2.5) Autres sources pertinentes trouvées

Auusi, il est important de voir comment les développeurs utilisent vraiment les nouveautés. J'ai sélectionné deux vidéos clés qui montrent bien l'évolution actuelle de Go.

- **YouTube / Vidéo** : 
	* [Go 1.23 - Custom Iterators Explained](https://www.youtube.com/watch?v=iurUVx0Nquc)
  		Cette vidéo explique en profondeur l’ajout majeur de Go 1.23 : le range-over-func (itérateurs personnalisés).
		Elle montre des exemples concrets et démontre pourquoi cette nouveauté change réellement la manière d’écrire des boucles en Go.
	* [Golang 1.23 - What’s New?](https://www.youtube.com/watch?v=EL4hg73mT2A)
		Cette vidéo offre une vue d’ensemble des nouveautés de Go 1.23, avec un bon contexte sur les raisons du changement et les implications pratiques.
		Une excellente source complémentaire qui montre comment la communauté perçoit cette mise à jour.

- **Réseaux sociaux / communautés** :  
    - [Reddit r/golang :](https://www.reddit.com/r/golang/)
		Discussions fréquentes sur les nouveautés du langage, dont plusieurs fils expliquant les cas d’usage des itérateurs.
	- [Twitter/X :](https://x.com/golang) golang


### 3) Sources choisies pour les notes de cours

Pour construire mes notes de cours sur Golang, nous avons sélectionné des sources officielles, stables et reconnues par la communauté Go. Le but était d’avoir des informations fiables, à jour, et suffisamment pédagogiques pour un usage scolaire. Voici la liste des sources retenues et ce qu’elles m’ont apporté dans mes notes.

* [Go Documentation – Getting Started](https://go.dev/doc/) 

**Utilisé pour :**
   - Comprendre **l’installation** du SDK
   - Décrire la **structure d’un environnement** Go (`GOROOT`, `GOPATH`)
   - Présenter les **commandes de base** du workflow (`go build`, `go run`, `go test`)
Cette source a été essentielle pour établir la base théorique et pratique du langage.

* [Go for Visual Studio Code – VS Code Marketplace](https://marketplace.visualstudio.com/items?itemName=golang.go) 

**Utilisé pour :**
   - Configuration de l’IDE
   - Utilisation de `gopls` (autocomplétion, diagnostics)
   - Intégration des **tests**, du debugging (`dlv`) et du **formattage auto**
Cette ressource aide à montrer un environnement de développement moderne et efficace pour les étudiants.

* [Go Modules Reference](https://go.dev/doc/modules/managing-dependencies) 

**Utilisé pour :**
   - Expliquer `go.mod`, `go.sum`
   - Montrer l’usage de `go mod init`, `go mod tidy`, `replace`
   - Comprendre la **gestion des versions** et des **dépendances**
Indispensable pour introduire les projets structurés en Go, puisque tout le cours repose sur les modules.

* [Go Slices: Usage and Internals](https://go.dev/blog/slices-intro) 

**Utilisé pour :**
   - Expliquer le fonctionnement interne des **slices**
   - Clarifier **longueur, capacité et partage mémoire**
   - Illustrer les implications pratiques de `append`
Ce document m’a permis de rédiger une section claire et intuitive sur un concept souvent mal compris.

* [Standard Library Overview](https://pkg.go.dev/std) 

**Utilisé pour :**
   - Introduire les **packages** les plus utilisés (`fmt`, `os`, `net/http`, `encoding/json`, etc.)
   - Montrer les **fonctions** essentielles avec exemples simples
Base indispensable pour développer des exercices concrets.

* [Effective Go](https://go.dev/doc/effective_go) 

**Utilisé pour :**
   - Présenter les **conventions** Go
   - Bonnes **pratiques** de style, erreurs, documentation
   - Recommandations pour écrire du **code idiomatique**
Cette source améliore la qualité du code présenté dans les notes et enseigne la philosophie Go.

* [Go FAQ – Methods and Interfaces](https://go.dev/doc/faq#methods) 

**Utilisé pour :**
   - Clarifier la différence entre **receivers par valeur / pointeur**
   - Expliquer le fonctionnement particulier des **interfaces**
   - Donner des guidelines simples pour leur **utilisation**
Utile pour éviter les erreurs typiques des débutants.

* [Share Memory By Communicating – Go Codewalk](https://go.dev/doc/codewalk/sharemem/) 

**Utilisé pour :**
   - Présenter les goroutines
   - Expliquer les channels avec des exemples pratiques
   - Introduire l’approche "share memory by communicating"
Parfait pour aborder la concurrence propre à Go.

* [Type Parameters Proposal and Guide](https://go.dev/doc/tutorial/generics) 

**Utilisé pour :**
   - Introduction progressive à la **syntaxe**
   - Exemple de fonctions paramétrées
   - Conseils pour éviter la complexité excessive
Essentiel pour une compréhension moderne du langage depuis Go 1.18.

* [Go net/http Package Documentation](https://pkg.go.dev/net/http) 

**Utilisé pour :**
   - Créer des serveurs HTTP simples
   - Expliquer `ServeMux`, `Handler`, `http.Client`
   - Montrer comment tester un handler avec `httptest`
Cette source m’a permis d’inclure une section complète sur la programmation Web en Go.

Chaque source est citée dans [Notes de cours - Sources](./notes_de_cours/#sources) avec le lien et la date de consultation.


## 4) Conclusion personnelle

Voici ce que je retiens personnellement du LLM vs les sources officielles et comment j’organiserai mon apprentissage:

- **LLM (ChatGPT)** : Ils sont parfaits pour expliquer un concept complexe en quelques secondes ou pour générer un squelette de code. C'est un bon point de départ pour "comprendre l'idée".
- **Sources officielles** : plus précises et fiables pour les détails techniques et les bonnes pratiques, mais parfois plus difficiles à lire et dispersées. Donc ellent restent quand indispensables pour la précision.
- **Organisation de l’apprentissage** :
    * Lire attentivement les sources officielles pour chaque notion importante.
    * Tester immédiatement les exemples de code pour bien comprendre leur fonctionnement.
    * Noter les points difficiles ou les astuces pour s’y référer plus tard.
    * Passer progressivement à des concepts plus avancés une fois les bases maîtrisées.
