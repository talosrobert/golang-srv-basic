# golang-srv-basic

golang REST server with basic authentication

## endpoints
- POST   /ad/ `create an ad, returns ID`
- GET    /ad/<adid> `returns a single ad by ID`
- GET    /ad/ `returns all ads`
- DELETE /ad/<adid> `delete an ad by ID`
- GET    /tag/<tagname> `returns list of ads with this tag`

## environment setup

~~~bash
$ podman image pull haproxy:lts-alpine
$ podman image pull postgres:16-alpine
$ podman network create appl
$ podman pod create -n appl --network appl
$ podman container run -d -t --pod appl --network appl --name db -e POSTGRES_PASSWORD=${APPL_POSTGRES_PWD} -e POSTGRES_DB=appl postgres:16-alpine
~~~

Building the container image for the application
~~~bash
$ podman image build -t appl -f scripts/Containerfile_appl .
~~~

Running the custom image file with the golang binary (with port exposed until the haproxy config is ready):
~~~bash
$ podman container run -d -t --pod appl --network appl --name appl -e APPL_POSTGRES_DSN=${APPL_POSTGRES_DSN} -p 4000:4000 localhost/appl:latest
~~~
