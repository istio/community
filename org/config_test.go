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
	"fmt"
	"sort"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/test-infra/prow/config/org"
	"k8s.io/test-infra/prow/github"
)

// nolint:unused
func testDuplicates(list sets.String) error {
	found := sets.String{}
	dups := sets.String{}
	all := list.List()
	for _, i := range all {
		if found.Has(i) {
			dups.Insert(i)
		}
		found.Insert(i)
	}
	if n := len(dups); n > 0 {
		return fmt.Errorf("%d duplicate names: %s", n, strings.Join(dups.List(), ", "))
	}
	return nil
}

func isSorted(list []string) bool {
	items := make([]string, len(list))
	for _, l := range list {
		items = append(items, strings.ToLower(l))
	}

	return sort.StringsAreSorted(items)
}

func normalize(s sets.String) sets.String {
	out := sets.String{}
	for i := range s {
		out.Insert(github.NormLogin(i))
	}
	return out
}

// testTeamMembers ensures that a user is not a maintainer and member at the same time,
// there are no duplicate names in the list and all users are org members.
func testTeamMembers(teams map[string]org.Team, admins sets.String, orgMembers sets.String) []error { //nolint:unparam // admins
	var errs []error
	for teamName, team := range teams {
		teamMaintainers := normalize(sets.NewString(team.Maintainers...))
		if len(teamMaintainers.List()) > 0 {
			errs = append(errs, fmt.Errorf("all users should be under members in %v", teamName))
		}
		teamMembers := normalize(sets.NewString(team.Members...))

		// check if all are org members
		if missing := teamMembers.Difference(orgMembers); len(missing) > 0 {
			errs = append(errs, fmt.Errorf("the following members of team %s are not parent team members: %s", teamName, strings.Join(missing.List(), ", ")))
		}

		// check if lists are sorted
		if !isSorted(team.Maintainers) {
			errs = append(errs, fmt.Errorf("the team %s has an unsorted list of maintainers", teamName))
		}
		if !isSorted(team.Members) {
			errs = append(errs, fmt.Errorf("the team %s has an unsorted list of members", teamName))
		}

		if team.Children != nil {
			errs = append(errs, testTeamMembers(team.Children, admins, teamMembers)...)
		}
	}
	return errs
}

func TestConvertedOrg(t *testing.T) {
	cfg, err := readConfig(".")
	if err != nil {
		t.Fatal(err)
	}
	org := convertConfig(cfg).Orgs["istio"]
	maintainers := org.Teams["Maintainers"]
	leads := org.Teams["Working Group Leads"]
	allLeads := sets.NewString(leads.Members...)
	allMaintainers := sets.NewString(maintainers.Members...)
	childMaintainers := sets.NewString()
	for _, ct := range maintainers.Children {
		childMaintainers.Insert(ct.Members...)
	}
	if diff := allMaintainers.Difference(childMaintainers); diff.Len() != 0 {
		t.Errorf("top level maintainer not in a team: %v", diff.UnsortedList())
	}
	if diff := childMaintainers.Difference(allMaintainers); diff.Len() != 0 {
		t.Errorf("team maintainer not in top level: %v", diff.UnsortedList())
	}
	if diff := allLeads.Difference(allMaintainers); diff.Len() != 0 {
		t.Errorf("leads not in maintainers: %v", diff.UnsortedList())
	}
}

func TestIstioOrg(t *testing.T) {
	cfg, err := readConfig(".")
	if err != nil {
		t.Fatal(err)
	}
	members := normalize(sets.NewString(cfg.Members...))
	admins := normalize(sets.NewString(cfg.Admins...))
	allOrgMembers := members.Union(admins).Union(members)

	requiredRobots := sets.NewString("istio-testing", "google-admin", "googlebot")
	if !admins.IsSuperset(requiredRobots) {
		t.Errorf("Missing required robots as admins: %v", requiredRobots.List())
	}

	if !isSorted(cfg.Members) {
		t.Errorf("members unsorted")
	}
	if !isSorted(cfg.Admins) {
		t.Errorf("admins unsorted")
	}

	if errs := testTeamMembers(cfg.Teams, admins, allOrgMembers); errs != nil {
		for _, err := range errs {
			t.Error(err)
		}
	}
}
