/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/test-infra/prow/config/org"
	"k8s.io/test-infra/prow/github"

	"github.com/ghodss/yaml"
)

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
func testTeamMembers(teams map[string]org.Team, admins sets.String, orgMembers sets.String, orgName string) []error {
	var errs []error
	for teamName, team := range teams {
		teamMaintainers := sets.NewString(team.Maintainers...)
		teamMembers := sets.NewString(team.Members...)

		teamMaintainers = normalize(teamMaintainers)
		teamMembers = normalize(teamMembers)

		// check for non-admins in maintainers list
		if nonAdminMaintainers := teamMaintainers.Difference(admins); len(nonAdminMaintainers) > 0 {
			errs = append(errs, fmt.Errorf("the team %s in org %s has non-admins listed as maintainers; these users should be in the members list instead: %s", teamName, orgName, strings.Join(nonAdminMaintainers.List(), ",")))
		}

		// check for users in both maintainers and members
		if both := teamMaintainers.Intersection(teamMembers); len(both) > 0 {
			errs = append(errs, fmt.Errorf("the team %s in org %s has users in both maintainer admin and member roles: %s", teamName, orgName, strings.Join(both.List(), ", ")))
		}

		// check for duplicates
		if err := testDuplicates(teamMaintainers); err != nil {
			errs = append(errs, fmt.Errorf("the team %s in org %s has duplicate maintainers: %v", teamName, orgName, err))
		}
		if err := testDuplicates(teamMembers); err != nil {
			errs = append(errs, fmt.Errorf("the team %s in org %s has duplicate members: %v", teamMembers, orgName, err))
		}

		// check if all are org members
		if missing := teamMembers.Difference(orgMembers); len(missing) > 0 {
			errs = append(errs, fmt.Errorf("the following members of team %s are not %s org members: %s", teamName, orgName, strings.Join(missing.List(), ", ")))
		}

		// check if admins are a regular member of team
		if adminTeamMembers := teamMembers.Intersection(admins); len(adminTeamMembers) > 0 {
			errs = append(errs, fmt.Errorf("the team %s in org %s has org admins listed as members; these users should be in the maintainers list instead, and cannot be on the members list: %s", teamName, orgName, strings.Join(adminTeamMembers.List(), ", ")))
		}

		// check if lists are sorted
		if !isSorted(team.Maintainers) {
			errs = append(errs, fmt.Errorf("the team %s in org %s has an unsorted list of maintainers", teamName, orgName))
		}
		if !isSorted(team.Members) {
			errs = append(errs, fmt.Errorf("the team %s in org %s has an unsorted list of members", teamName, orgName))
		}

		if team.Children != nil {
			errs = append(errs, testTeamMembers(team.Children, admins, orgMembers, orgName)...)
		}
	}
	return errs
}

func TestIstioOrg(t *testing.T) {
	raw, err := ioutil.ReadFile("./istio.yaml")
	if err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}
	var org org.Config
	if err := yaml.Unmarshal(raw, &org); err != nil {
		t.Fatalf("cannot read istio.yaml: %v", err)
	}
	members := normalize(sets.NewString(org.Members...))
	admins := normalize(sets.NewString(org.Admins...))
	allOrgMembers := members.Union(admins)

	if both := admins.Intersection(members); len(both) > 0 {
		t.Errorf("users in both org admin and member roles for org '%s': %s", *org.Name, strings.Join(both.List(), ", "))
	}

	requiredRobots := sets.NewString("istio-testing", "google-admin", "googlebot")
	if !admins.IsSuperset(requiredRobots) {
		t.Errorf("Missing required robots as admins: %v", requiredRobots.List())
	}

	if org.BillingEmail != nil {
		t.Errorf("billing_email must be unset")
	}

	if err := testDuplicates(admins); err != nil {
		t.Errorf("duplicate admins: %v", err)
	}
	if err := testDuplicates(allOrgMembers); err != nil {
		t.Errorf("duplicate members: %v", err)
	}
	if !isSorted(org.Admins) {
		t.Errorf("admins for %s org are unsorted", *org.Name)
	}
	if !isSorted(org.Members) {
		t.Errorf("members for %s org are unsorted", *org.Name)
	}

	if errs := testTeamMembers(org.Teams, admins, allOrgMembers, *org.Name); errs != nil {
		for _, err := range errs {
			t.Error(err)
		}
	}

}
