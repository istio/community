img := gcr.io/istio-testing/build-tools:2019-08-16
docker := docker run -e -t -i --sig-proxy=true --rm -v $(shell pwd):/foo -w /foo $(img)

lint:
	@$(docker) prow/community-lint.sh
