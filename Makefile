img := gcr.io/istio-testing/build-tools:2019-10-24T14-05-17
docker := docker run -e -t -i --sig-proxy=true --rm -v $(shell pwd):/foo -w /foo $(img)

lint:
	@$(docker) prow/community-lint.sh

test:
	go test ./...
