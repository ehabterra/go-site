#!/bin/bash

if [ ! "$(docker ps -q -f name=dev-postgres)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=dev-postgres)" ]; then
        # cleanup
        docker rm dev-postgres
    fi
    # run your container
    docker run -d \
        --name dev-postgres \
        -e POSTGRES_PASSWORD=secret \
        -e PGDATA=/var/lib/postgresql/data/pgdata \
        -v /${HOME}/postgresql/data/:/var/lib/postgresql/data \
        -p 5432:5432 \
        postgres:13-alpine
fi



if [ ! "$(docker ps -q -f name=dev-pgadmin)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=dev-pgadmin)" ]; then
        # cleanup
        docker rm dev-pgadmin
    fi
    # run your container
    docker run -d \
        --name dev-pgadmin \
        -p 810:80 \
        -e 'PGADMIN_DEFAULT_EMAIL=user@example.com' \
        -e 'PGADMIN_DEFAULT_PASSWORD=secret' \
        -d dpage/pgadmin4
fi

