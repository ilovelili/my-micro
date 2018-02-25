#!/bin/sh

# Go build
go build -o greeter-srv

# docker build
docker build .

# docker push
# docker tag imageID ilovelili/microkubesrv:firsttry

# docker push ilovelili/microkubesrv

kubectl create -f k8sgreeterdeploy.yml


