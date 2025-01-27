# Projet Devops S9

Ce projet met en œuvre un pipeline CI/CD pour une application écrite en Go. Le pipeline inclut des étapes pour compiler le code, créer une image Docker, et la pousser vers le registre GitHub Container Registry (GHCR).

## Table des Matières
- [Fonctionnalités](#fonctionnalités)
- [Prérequis](#prérequis)
- [Installation et Utilisation](#installation-et-utilisation)
- [Pipeline CI/CD](#pipeline-cicd)
- [Branche `docker-with-buildpack`](#branche-docker-with-buildpack)
- [Structure du Projet](#structure-du-projet)


## Fonctionnalités

- **Développement en Go** : Une API web basique avec des endpoints `/`, `/aboutme`, et `/whoami`.
- **Pipeline CI/CD** : Utilisation de GitHub Actions pour automatiser la compilation et la création d'images Docker.
- **Déploiement Docker** : Création et publication d'images Docker sur GHCR.

## Prérequis

- **Go** : Version 1.21.5 ou supérieure
- **Docker** : Installé et configuré
- **Compte GitHub** : Avec les permissions nécessaires pour pousser dans GHCR
- **GitHub Token** : Configuré comme un secret dans votre dépôt (`GITHUB_TOKEN`)

## Installation et Utilisation

### Cloner le dépôt
```bash
git clone https://github.com/Teroaz/projet-devops.git
cd projet-devops
```
Accédez à l'application sur http://localhost:8080.

Points de terminaison disponibles
- `/` : Page d'accueil
- `/aboutme` : Renvoie des informations sur le projet
- `/whoami` : Renvoie des informations sur les développeurs


## Pipeline CI/CD
Le pipeline est défini dans `.github/workflows/go-docker.yml`. 

Voici les étapes principales :
- Compilation Go, télécharge le code et configure l'environnement Go.
- Création d'image Docker, compile le binaire Go et crée une image Docker à partir du Dockerfile.
- Publication sur GHCR, publie l'image Docker avec le tag suivant :
ghcr.io/Teroaz/projet-devops:latest.


### Branche docker-with-buildpack
Les Buildpacks permettent de simplifier la construction d'images Docker en détectant automatiquement les dépendances requises pour l'application.

Pour utiliser cette branche :

```bash
git checkout docker-with-buildpack
```
Assurez-vous d'avoir les outils nécessaires comme Pack CLI.


## Structure du Projet
```bash
├── webapi/
│   ├── main.go
│   ├── go.mod        
│   ├── Dockerfile      
├── .github/
│   └── workflows/
│       └── build.yml
└── README.md
```
