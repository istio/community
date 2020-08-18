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

set -e

# Sync the Github organization according to config in org/
# Must be run by a Github admin (generally in CI)

PERIBOLOS="${PERIBOLOS:-peribolos}"
# Fallback to legacy peribolos binary if not found
if ! command -v "${PERIBOLOS}" &> /dev/null; then
	PERIBOLOS=/app/prow/cmd/peribolos/app.binary
fi

"${PERIBOLOS}" --fix-org --fix-org-members --fix-teams --fix-team-members \
	--config-path "${WD}"/../org/istio.yaml --github-token-path /etc/github-token/oauth --confirm
