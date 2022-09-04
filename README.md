# kuro-movies-api
REST and GraphQL API of movies and actors

WIth REST you can:

- Basic crud operations with movies and actors
- Register and login with JWT

With GraphQL you can:

- Query movies and actors

THIS API DON'T USE REAL INFORMATION.

This API was crated for learning purposes only.

See the [DOCS](https://documenter.getpostman.com/view/20195671/UzBpLRz8)

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/20195671-8e639575-089c-415a-b082-f2a4d23f0469?action=collection%2Ffork&collection-url=entityId%3D20195671-8e639575-089c-415a-b082-f2a4d23f0469%26entityType%3Dcollection%26workspaceId%3D340d12f8-bfd8-4f84-8bc7-f3b080c24682)

## Docker image

You can run this project with the [docker image](https://hub.docker.com/r/kurovale/kuro-movies) I made.

### Quick Setup

1. ```git clone https://github.com/kuro-vale/kuro-movies-api.git```

2. Set enviroment variables
    - DATABASE_UR = The url of your postgres database with sslmode disabled example: ```postgres://user:password@localhost:5432/kuro-movies?sslmode=disable```

    - SECRET_KEY = A secret key to sign JWT tokens
3. Meet go.mod dependencies
4. Run ```go run ./main.go```
