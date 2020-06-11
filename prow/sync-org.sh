#!/bin/bash

# Copyright 2020 Istio Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

# Sync the Github organization according to config in org/
# Must be run by a Github admin (generally in CI)

/app/prow/cmd/peribolos/app.binary --fix-org --fix-org-members --fix-teams --fix-team-members \
		--config-path ../org/istio.yaml --github-token-path /etc/github-token/oauth --confirm
