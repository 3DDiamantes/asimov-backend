#!/usr/bin/bash
go test -coverpkg=./... -coverprofile=cov.out -covermode=atomic ./... && go tool cover -html=cov.out && rm cov.out