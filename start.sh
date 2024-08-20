#! /bin/bash

go build -a -o go-admin-api main.go

./go-admin-api server -c=config/settings.dev.yml