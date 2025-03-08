#!/bin/sh

cd back && CGO_ENABLED=0 go build && cd ..
docker build back -t eclosix/back
docker build front -t eclosix/front