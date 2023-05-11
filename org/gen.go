// Copyright Istio Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"k8s.io/test-infra/prow/config/org"
	"k8s.io/test-infra/prow/github"
	"sigs.k8s.io/yaml"
)

type Organization struct {
	Admins  []string            `json:"admins,omitempty"`
	Members []string            `json:"members,omitempty"`
	Repos   map[string]org.Repo `json:"repos,omitempty"`
	Teams   map[string]org.Team `json:"teams,omitempty"`
}

var (
	input  = flag.String("input", "org", "folder to read config from")
	output = flag.String("output", "", "file to write config to")
)

func main() {
	flag.Parse()
	cfg, err := readConfig(*input)
	if err != nil {
		exit(err)
	}
	org := convertConfig(cfg)
	if err := writeConfig(org, *output); err != nil {
		exit(err)
	}
}

func strptr(s string) *string {
	return &s
}

// Remove all elements in drop from list
func drop(list []string, drop []string) []string {
	dd := map[string]struct{}{}
	for _, d := range drop {
		dd[d] = struct{}{}
	}
	res := []string{}
	for _, x := range list {
		if _, f := dd[x]; !f {
			res = append(res, x)
		}
	}
	return res
}

// Remove all elements in not in keep from list
func keep(list []string, keep []string) []string {
	dd := map[string]struct{}{}
	for _, d := range keep {
		dd[d] = struct{}{}
	}
	res := []string{}
	for _, x := range list {
		if _, f := dd[x]; f {
			res = append(res, x)
		}
	}
	return res
}

const defaultRepo = ".default"

func defaultRepos(cur map[string]github.RepoPermissionLevel) map[string]github.RepoPermissionLevel {
	def, df := cur[defaultRepo]
	if !df {
		// Default to none
		def = github.None
	}
	res := map[string]github.RepoPermissionLevel{
		"api":             def,
		"bots":            def,
		"client-go":       def,
		"cni":             def,
		"common-files":    def,
		"community":       def,
		"cri":             def,
		"enhancements":    def,
		"envoy":           def,
		"gogo-genproto":   def,
		"installer":       def,
		"istio":           def,
		"istio.io":        def,
		"operator":        def,
		"pkg":             def,
		"proxy":           def,
		"release-builder": def,
		"test-infra":      def,
		"tools":           def,
		"ztunnel":         def,
	}
	// Allow per-repo overrides
	for k, v := range cur {
		if k == defaultRepo {
			continue
		}
		res[k] = v
	}
	// Remove "none" entries; we will just not give any permission for this case
	for k, v := range res {
		if v == github.None {
			delete(res, k)
		}
	}
	return res
}

func strPointer(s string) *string {
	return &s
}

func convertConfig(cfg Organization) org.FullConfig {
	allMembers := cfg.Members
	sort.Slice(allMembers, func(i, j int) bool { return strings.ToLower(allMembers[i]) < strings.ToLower(allMembers[j]) })

	// Insert the members team, which is handled separately
	closed := org.Closed
	cfg.Teams["Members"] = org.Team{
		TeamMetadata: org.TeamMetadata{
			Description: strPointer("Folks actively working on Istio."),
			Privacy:     &closed,
		},
		Members: cfg.Members,
		Repos: map[string]github.RepoPermissionLevel{
			defaultRepo: github.Triage,
		},
	}

	istio := org.Config{
		Metadata: org.Metadata{
			Name:        strptr("Istio"),
			Description: strptr("Connect, secure, control, and observe services."),
		},
		Teams: cfg.Teams,
		// Members list shouldn't contain admins
		Members: drop(allMembers, cfg.Admins),
		Admins:  cfg.Admins,
		Repos:   cfg.Repos,
	}
	normalizeTeams(istio, cfg.Admins)
	return org.FullConfig{
		Orgs: map[string]org.Config{
			"istio": istio,
		},
	}
}

func normalizeTeam(t org.Team, admins []string) org.Team {
	closed := org.Closed
	t.Maintainers = keep(t.Members, admins)
	t.Members = drop(t.Members, admins)
	t.Repos = defaultRepos(t.Repos)
	if t.Privacy == nil {
		t.Privacy = &closed
	}
	for k, child := range t.Children {
		t.Children[k] = normalizeTeam(child, admins)
	}
	return t
}

func normalizeTeams(istio org.Config, admins []string) {
	for k, t := range istio.Teams {
		istio.Teams[k] = normalizeTeam(t, admins)
	}
}

func writeConfig(cfg org.FullConfig, output string) error {
	yml, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(output, yml, 0o750)
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "failed to generate org: %v\n", err)
	os.Exit(-1)
}

func readConfig(input string) (Organization, error) {
	dir, err := os.ReadDir(input)
	if err != nil {
		return Organization{}, fmt.Errorf("failed to read %v: %v", input, err)
	}
	cfg := Organization{}
	for _, f := range dir {
		if strings.HasSuffix(f.Name(), ".yaml") {
			contents, err := os.ReadFile(filepath.Join(input, f.Name()))
			if err != nil {
				return cfg, fmt.Errorf("failed to read file %v: %v", f.Name(), err)
			}
			if err := yaml.Unmarshal(contents, &cfg); err != nil {
				return cfg, fmt.Errorf("invalid yaml: %v", err)
			}
		}
	}
	return cfg, nil
}
