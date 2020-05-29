img := gcr.io/istio-testing/build-tools:2019-10-24T14-05-17
docker := docker run -e -t -i --sig-proxy=true --rm -v $(shell pwd):/foo -w /foo $(img)

lint:
	@$(docker) prow/community-lint.sh

test:
	go test ./...

# Sync the Github organization according to config in org/
# Must be run by a Github admin (generally in CI)
sync-org:
	/app/prow/cmd/peribolos/app.binary --fix-org --fix-org-members --fix-teams --fix-team-members \
		--config-path org/istio.yaml --github-token-path /etc/github-token/oauth --confirm