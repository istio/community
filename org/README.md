# Istio Organization

The folder contains configuration for the Istio Github Organization. Data here is automatically synced with Github using [peribolos](https://github.com/kubernetes/test-infra/tree/master/prow/cmd/peribolos).

## Making Changes

To modify the org, simply change the config file and submit a PR. Once the PR is merged, the org will be updated.

Changes can be tested with

```bash
make test
```

To add yourself to the Org, please add your name in alphabetical order under `members.yaml` and fill out the form in the PR description.

Teams, such as working groups and release managers, are controlled by the `teams.yaml` file.
