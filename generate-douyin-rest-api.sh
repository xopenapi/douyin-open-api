#!/bin/sh
docker run --rm -v "${PWD}:/local" xopenapi/openapi-generator-cli generate \
    -i /local/douyin-open-api.yaml \
    --git-repo-id douyin-open-api-go \
    --git-user-id xopenapi \
    --package-name douyin \
    -g go \
    -o /local/out/douyin-open-api-go

docker run --rm -v "${PWD}:/local" xopenapi/openapi-generator-cli generate \
    -i /local/douyin-open-api.yaml \
    --git-repo-id douyin-open-api-java \
    --git-user-id xopenapi \
    --package-name com.xopenapis.douyin \
    -g java \
    -o /local/out/douyin-open-api-java
