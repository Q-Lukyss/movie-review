# Application Movie Review 

DISCLAIMER: En Developpement

## Contexte

Le but de cette petite application est de permettre aux utilisateurs de créer des critiques de film et de les partager avec la communauté

## Stack 

[Docker](https://www.docker.com/), [Traefik](https://doc.traefik.io/traefik/), [Nuxt](https://nuxt.com/) et [TypeScript](https://www.typescriptlang.org/), [Go](https://go.dev/) et [Gin](https://gin-gonic.com/), [Tailwind](https://tailwindcss.com/)

## Prérequis

[Docker](https://www.docker.com/) installé

## Configuration

Récupérer via git ou .zip   
Renommer le fichier env en .env   
Pour le développement local vous pouvez laisser les champs DOMAIN et LETSENCRYPT_EMAIL par défaut, à changer en prod

## Developpement Local

A la racine du projet executez:
```
docker compose -f docker-compose.local.yml up -d --build
```
Vous pouvez ensuite accéder au dashboard traefik à `http://localhost:<TRAEFIK_DASHBOARD_PORT>/dashboard/`   
PS: le dernier `/` après dashboard est important   
Au frontend Nuxt à `http://localhost:<TRAEFIK_HTTP_PORT>/`   
Et à l'api Go Gin à `http://localhost:<TRAEFIK_HTTP_PORT>/api/`   

## Production

A venir