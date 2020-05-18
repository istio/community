# Istio Organization

The folder contains configuration for the Istio Github Organization. Data here is automatically synced with Github using [peribolos](https://github.com/kubernetes/test-infra/tree/master/prow/cmd/peribolos).

## Making Changes

To modify the org, simply change the config file and submit a PR. Once the PR is merged, the org will be updated.

Changes can be tested with `make test`.

To add yourself to the Org, please add your name under `members.yaml`. 

To add yourself as a developer, which is given additional access to repos, add yourself under `developers.yaml` instead.

Teams, such as working groups and release managers, are controlled by the `teams.yaml` file.
