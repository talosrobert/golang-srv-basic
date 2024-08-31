# golang-srv-basic

The goal of this project is for me to learn about `golang` by creating a JSON API server, mimicking an auction site.
It is using the native `net/http` library for routing and PostgreSQL as the database for persistently storing all the data.

## environment setup
For creating the test environment I've been using `podman` to run both the backend application and the database instance. Rather then building HTTP rate-limiting and timeouts into the application, my plan is to run HAProxy in front of the backend application and do all the traffic-policing in the load-balancer configuration. 

Pull each necessary container image, create a separate network and pod, where each container will run.
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
$ podman container run -d -t --pod appl --network appl --name appl -e APPL_POSTGRES_DSN="postgres://appl_role:${APPL_ROLE_POSTGRES_PWD}@db/appl?search_path=appl" -p 4000:4000 localhost/appl:latest
~~~

## load-balancer configuration

TODO
