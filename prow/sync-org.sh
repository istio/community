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

WD=$(dirname "$0")
WD=$(cd "$WD"; pwd)

set -eu

OUT_DIR=$(mktemp -d)

# Sync the Github organization according to config in org/
# Must be run by a Github admin (generally in CI)
echo "Writing generated org file to ${OUT_DIR}"

go run org/gen.go --output "${OUT_DIR}/istio.yaml"

echo "Generated configuration: $(cat "${OUT_DIR}/istio.yaml")"

peribolos --fix-org --fix-org-members --fix-teams --fix-team-members --fix-team-repos \
	--config-path "${OUT_DIR}/istio.yaml" --github-token-path /etc/github-token/oauth --confirm
