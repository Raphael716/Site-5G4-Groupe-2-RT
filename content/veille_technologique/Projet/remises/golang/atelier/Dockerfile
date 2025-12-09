# Image de base Go
FROM golang:1.23-alpine

# Installer outils utiles
RUN apk add --no-cache git nano bash

# Créer un dossier de travail
WORKDIR /app

# Copier le projet
COPY go.mod ./

RUN go mod download

COPY . .

# Compiler le programme principal (si tu veux lancer un binaire par défaut)
RUN echo "Bienvenue dans l'atelier Go ! Tapez 'go run exercises/ex1_hello/main.go' pour tester." > /etc/motd

# Commande par défaut : shell interactif
CMD ["/bin/bash"]
