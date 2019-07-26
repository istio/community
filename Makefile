img := gcr.io/istio-testing/website-tools:2019-07-25
docker := docker run -e -t -i --sig-proxy=true --rm -v $(shell pwd):/foo -w /foo $(img)

lint:
	@$(docker) prow/community-lint.sh
