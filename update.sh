#!/bin/bash

VERSION=v1.4.1
curl https://raw.githubusercontent.com/alephium/alephium/${VERSION}/api/src/main/resources/openapi.json -o openapi.json
openapi-generator generate -i ./openapi.json -g go -o ./ --skip-validate-spec --package-name alephium --git-repo-id alephium-go-sdk --git-user-id lbqds
rm openapi.json
rm git_push.sh

git restore model_val.go model_tx_status.go

go mod tidy
go build ./
go test ./

