#!/usr/bin/env bash

# Copyright Istio Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Regenerates .project/maintainers.yaml from istio/community's org/teams.yaml.
#
# The maintainer roster is the union of the "Maintainers" and "Steering
# Committee" teams, so Steering Committee members always keep maintainer
# access in .project even when they are not separately listed on the
# "Maintainers" team.
#
# Runs with the current working directory set to a clone of istio/.project
# (see prow/config/jobs/community.yaml in istio/test-infra).

set -euo pipefail

TMP=$(mktemp -d)
trap 'rm -rf "$TMP"' EXIT

BRANCH=${AUTOMATOR_SRC_BRANCH:-master}
git clone -q --depth 1 --single-branch --branch "$BRANCH" \
  https://github.com/istio/community "$TMP/community"

TEAMS_FILE="$TMP/community/org/teams.yaml"

maintainers=$(yq eval '.teams.Maintainers.members[]' "$TEAMS_FILE")
steering=$(yq eval '.teams["Steering Committee"].members[]' "$TEAMS_FILE")

printf '%s\n%s\n' "$maintainers" "$steering" | sort -u -f > "$TMP/members.txt"

if [[ ! -s maintainers.yaml ]]; then
  echo "error: maintainers.yaml not found (or empty) in $(pwd)" >&2
  exit 1
fi

# Replace only the (first) "members:" list in-place, preserving indentation,
# comments, and formatting of the rest of the file, instead of rewriting the
# whole document.
awk -v members_file="$TMP/members.txt" '
  /^[[:space:]]*members:[[:space:]]*$/ && !done {
    match($0, /^[[:space:]]*/)
    item_indent = substr($0, RSTART, RLENGTH) "  "
    print
    while ((getline name < members_file) > 0) {
      print item_indent "- " name
    }
    close(members_file)
    in_list = 1
    done = 1
    next
  }
  in_list && $0 ~ "^" item_indent "- " { next }
  { in_list = 0; print }
  END { if (!done) { print "error: no members: list found in maintainers.yaml" > "/dev/stderr"; exit 1 } }
' maintainers.yaml > "$TMP/maintainers.yaml.new"
mv "$TMP/maintainers.yaml.new" maintainers.yaml
