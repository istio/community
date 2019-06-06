# Istio Community Roles

This document describes the set of roles individuals may have within the Istio community, the requirements of each role, and the privileges that each role grants.

* [Role summary](#role-summary)
* [Collaborator](#collaborator)
* [Member](#member)
* [Maintainer](#maintainer)
* [Triager](#triager)
* [Lead](#lead)
* [Administrator](#administrator)

Transient roles

* [Release manager](#release-manager)
* [Sponsor](#sponsor)

## Role summary

Here is the set of roles we use within the Istio community, the general responsibilities expected by individuals in each role,
the requirements necessary to join or stay in a given role, and the concrete manifestation of the role in terms of permissions and
privileges.

<table>
  <tr>
    <td>Role</td>
    <td>Responsibilities</td>
    <td>Requirements</td>
    <td>Privileges</td>
  </tr>

  <tr>
    <td><a href="#collaborator">Collaborator</a></td>
    <td>Casual contributor to the project</td>
    <td>n/a</td>
    <td>
        <p>Outside collaborator of the GitHub Istio organization</p>
        <p>Can submit PRs</p>
        <p>Read and commenting permission on the Istio Team drive</p>
    </td>
  </tr>

  <tr>
    <td><a href="#member">Member</a></td>
    <td>Regular active contributor in the community</td>
    <td>
        <p>Has pushed at least one PR to an Istio repository</p>
    </td>
    <td>
        <p>Member of the GitHub Istio organization</p>
        <p>Edit permission on the Istio Team drive</p>
        <p>Triage permission on the Istio repos, allowing issues to be manipulated.</p>
    </td>
  </tr>

  <tr>
    <td><a href="#maintainer">Maintainer</a></td>
    <td>Approve contributions from other members</td>
    <td>Highly experienced and active reviewer and contributor to an area</td>
    <td>Like a member, plus:
        <p>Able to approve code changes in GitHub</p>
        <p>Voting rights in the context of working group decision-making</p>
    </td>
  </tr>

  <tr>
    <td><a href="#trigger">Triager</a></td>
    <td>
        <p>Triage incoming issues, set milestones, set labels</p>
    </td>
    <td>Appointed by the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a></td>
    <td>Triage permissions on the Istio repos, allowing issues to be manipulated.</td>
  </tr>

  <tr>
    <td><a href="#lead">Lead</a></td>
    <td>
        <p>Set priorities for a functional area and approve proposals</p>
        <p>Triage incoming issues, set milestones, set labels</p>
        <p>Run their working group</p>
        <p>Write permission on the Istio repos.</p>
    </td>
    <td>Appointed by the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a> as documented in <a href="./WORKING-GROUP-PROCESSES.md">Istio Working Group Processes</a></td>
    <td>Like a maintainer
  </tr>

  <tr>
    <td><a href="#administrator">Administrator</a></td>
    <td>Manage & control permissions</td>
    <td>Appointed by the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a></td>
    <td>
        <p>Admin privileges on varous Istio-related resources, as defined in <a href="./ADMINS-FOR-ISTIO.md">ADMINS-FOR-ISTIO</a></p>
    </td>
  </tr>

  <tr>
    <td><a href="#release-manager">Release manager</a></td>
    <td>Sheperd a release through to general availability</td>
    <td>Appointed by the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a></td>
    <td>Admin privilege over the GitHub repos</td>
  </tr>

  <tr>
    <td><a href="#sponsor">Sponsor</a></td>
    <td>Vouch for another GitHub user in order to let the user get more privilege within the project</td>
    <td>
        <p>Must have close interactions with the user</p>
    </td>
    <td>n/a</td>
  </tr>

</table>

## Collaborator

Individuals may be added as an outside collaborator (with READ access) to a repo in the Istio GitHub organization without becoming a member.
This allows them to be assigned issues and PRs until they become a member, but will not allow tests to be run against their PRs automatically
nor allow them to interact with the PR bot.

### Requirements

* Working on some contribution to the project that would benefit from the ability to have PRs or Issues to be assigned to the contributor

* <a href="#sponsor">Sponsored</a> by 1 member

## Member

Established community members are expected to demonstrate their adherence to the principles in this document, familiarity with project
organization, roles, policies, procedures, conventions, etc., and technical and/or writing ability.

Members are continuously active contributors in the community. They can have issues and PRs assigned to them, participate in working group
meetings, and pre-submit tests are automatically run for their PRs. Members are expected to remain active contributors to the community.

All members are encouraged to help with the code review burden, although each PR must be reviewed by one or more official maintainers for
the area before being accepted into the source base.

### Requirements

* Has pushed at least one PR to the Istio repositories within the last 6 months.

* Subscribed to [contributors](https://discuss.istio.io/c/contributors)

* Actively contributing to 1 or more areas.

### Responsibilities and privileges

* Responsive to issues and PRs assigned to them

* Active owner of code they have contributed (unless ownership is explicitly transferred)

    * Code is well tested

    * Tests consistently pass

    * Addresses bugs or issues discovered after code is accepted

Members who frequently contribute code are expected to proactively perform code reviews for the area that they are active in.

## Maintainer

Maintainers review and approve code contributions. While code review is focused on code quality and correctness,
approval is focused on holistic acceptance of a contribution including: backwards / forwards compatibility, adhering to API and flag
conventions, subtle performance and correctness issues, interactions with other parts of the system, etc. Maintainer status is scoped to a
part of the codebase and is reflected in a GitHub CODEOWNERS file or a prow OWNERS file.

### Requirements

The following apply to the part of the codebase for which one would be a maintainer:

* Member for at least 3 months

* Contributed at least 30 substantial PRs to the codebase

* Must remain an active participant in the community by contributing
code, performing reviews, triaging issues, etc.

* Knowledgeable about the codebase

* Sponsored by a working group lead with no objections from other leads

### Responsibilities and privileges

The following apply to the part of the codebase for which one would be a maintainer:

* Maintainer status may be a precondition to accepting large code contributions

* Demonstrates sound technical judgement

* Responsible for project quality control via [code reviews](https://github.com/istio/istio/wiki/Reviewing-Pull-Requests)

    * Focus on code quality and correctness, including testing and factoring

    * Focus on holistic acceptance of contribution such as dependencies with other features, backwards / forwards compatibility, API and flag definitions, etc

* Expected to be responsive to review requests as per [community expectations](https://github.com/istio/istio/wiki/Reviewing-Pull-Requests);

* May approve code contributions for acceptance

* Maintainers in an area get a stronger voice when a working group needs to make decisions.

### Emeritus status

If a maintainer becomes inactive in the project for an extended period of time, the individual will transition to being a
emeritus maintainer. Emeritus maintainers lose their ability to approve code contributions, but retain their voting rights
for up to one year. After one year, emeritus maintainers revert back to being normal members. 

A maintainer becomes an emeritus maintainer if the individual hasn't contributed a PR to the project in 3 months. An emeritus 
maintainer can regain full maintainer status by submitting 5 PRs to the project.

## Triager

Triagers are members that help triage issues. They can set priority, assignment, milestones, and labels.

### Requirements

Triagers are appointed by the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a>

### Responsibilities and privileges

The following apply to the area / component for which one would be an owner.

* Perform issue triage on GitHub

* Apply/remove/create/delete GitHub labels and milestones

## Lead

Working group leads, or just ‘leads’, are maintainers of an entire area that have demonstrated good judgement and responsibility.
Leads accept design proposals and approve design decisions for their area of ownership, triage issues, and run regular working group
meetings.

### Requirements

Getting to be a lead of an existing working group:

* Recognized as having expertise in the group’s subject matter

* Maintainer for some part of the codebase for at least 3 months

* Member for at least 1 year

* Primary reviewer for 20 substantial PRs

* Reviewed or merged at least 50 PRs

* Appointed by the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a>

Establishing the leads for a new working group:

* Originally authored or contributed major functionality to an area

* A maintainer in the top-level OWNERS file for the group’s code

* Maintainer for some part of the codebase for at least 3 months

* Member for at least 1 year

* Primary reviewer for 20 substantial PRs

* Reviewed or merged at least 50 PRs

* Appointed by the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a>

### Responsibilities and privileges

The following apply to the area / component for which one would be an owner.

* Run their working group as outlined in the [Working Group Processes](WORKING-GROUP-PROCESSES.md) document.

* Design/proposal approval authority over the area / component, though escalation to the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a> is possible

* Perform issue triage on GitHub

* Apply/remove/create/delete GitHub labels and milestones

* Expected to work to holistically maintain the health of the project through:

    * Reviewing PRs

    * Fixing bugs

    * Mentoring and guiding maintainers, members, and contributors

## Administrator

Administrators are responsible for the bureaucratic aspects of the project.

### Requirements

* Appointed by the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a>

### Responsibilities and privileges

* Manage a variety of infrastructure support for the Istio project as described in [ADMINS-FOR-ISTIO](./ADMINS-FOR-ISTIO.md)

* Although admins may have the authority to override any policy and cut corners, we expect admins to generally abide
by the overall rules of the project. For example, unless strictly necessary, admins should not approve and/or commit
PRs they aren't entitled to if they were not admins.

## Release manager

Individual responsible for a particular release of the product. This person takes ownership of the
release processes and knocks down obstacles in order to drive to a successful timely general availability of
the release.

### Requirements

* Appointed by the <a href="./TECH-OVERSIGHT-COMMITTEE.md">technical oversight committee</a>

### Responsibilities and privileges

* Get a release out the door.

* Admin privilege over the GitHub repos.

## Sponsor

A sponsor is a member that wishes to vouch for another user to support a user being assigned a specific role.

### Requirements

* Sponsors must have close interactions with the user - e.g. code/design/proposal review, coordinating on issues, etc.

* Sponsors are responsible for assessing whether a new potential user behaves in accordance to our <a href="./CONTRIBUTING.md">contribution guidelines</a>.
