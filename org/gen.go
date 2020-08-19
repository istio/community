package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ghodss/yaml"
	"k8s.io/test-infra/prow/config/org"
	"k8s.io/test-infra/prow/github"
)

type Organization struct {
	Admins     []string            `json:"admins,omitempty"`
	Members    []string            `json:"members,omitempty"`
	Developers []string            `json:"developers,omitempty"`
	Repos      map[string]org.Repo `json:"repos,omitempty"`
	Teams      map[string]org.Team `json:"teams,omitempty"`
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

var repos = map[string]github.RepoPermissionLevel{
	"api":             github.Write,
	"bots":            github.Write,
	"client-go":       github.Write,
	"cni":             github.Write,
	"common-files":    github.Write,
	"cri":             github.Write,
	"envoy":           github.Write,
	"gogo-genproto":   github.Write,
	"installer":       github.Write,
	"istio":           github.Write,
	"istio.io":        github.Write,
	"operator":        github.Write,
	"pkg":             github.Write,
	"proxy":           github.Write,
	"release-builder": github.Write,
	"test-infra":      github.Write,
	"tools":           github.Write,
}

func strPointer(s string) *string {
	return &s
}

func convertConfig(cfg Organization) org.FullConfig {
	allMembers := cfg.Members[:]
	allMembers = append(allMembers, cfg.Developers...)
	sort.Slice(allMembers, func(i, j int) bool { return strings.ToLower(allMembers[i]) < strings.ToLower(allMembers[j]) })

	// Insert the developers team, which is handled separately
	closed := org.Closed
	cfg.Teams["Developers"] = org.Team{
		TeamMetadata: org.TeamMetadata{
			Description: strPointer("Folks actively working on the Istio code base."),
			Privacy:     &closed,
		},
		Members: cfg.Developers,
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
	if t.Repos == nil {
		t.Repos = repos
	}
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
	return ioutil.WriteFile(output, yml, 0750)
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "failed to generate org: %v\n", err)
	os.Exit(-1)
}

func readConfig(input string) (Organization, error) {
	dir, err := ioutil.ReadDir(input)
	if err != nil {
		return Organization{}, fmt.Errorf("failed to read %v: %v", input, err)
	}
	cfg := Organization{}
	for _, f := range dir {
		if strings.HasSuffix(f.Name(), ".yaml") {
			contents, err := ioutil.ReadFile(filepath.Join(input, f.Name()))
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
