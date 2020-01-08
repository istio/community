# Istio Organization

The [`istio.yaml`](istio.yaml) file contains configuration for the Istio Github Organization. Data here is automatically synced with Github using [peribolos](https://github.com/kubernetes/test-infra/tree/master/prow/cmd/peribolos).

## Making Changes

To modify the org, simply change the config file and submit a PR. Once the PR is merged, the org will be updated.

Changes can be tested with `make test`.

To add yourself to the Org, you just need to add your name to the `members` list.
