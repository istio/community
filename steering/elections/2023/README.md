# 2023 Istio Steering Committee election guide

The Steering Committee is composed of nine (9) Contribution Seats and four (4) Community Seats. The role of this election is to elect the four Community Seats to the Istio Steering Committee for a one-year term.

## Background

The Istio Steering Committee oversees the administrative aspects of the project, including governance, marketing and community health. Steering exists to allow the Technical Oversight Committee to exclusively focus on the technical aspects of the project.

Some direct responsibilities of steering members to consider as you are deciding whether to run or who to vote for:

- Establishing and enforcing principles to guide the Istio project.
- Fostering an environment for a healthy and happy community of developers, contributors, vendors, and users.
- Defining, evolving, and defending a Code of Conduct.
- Setting marketing and advocacy direction for the project; establishing a publishing schedule and vetting content, encouraging and assisting community members in creating content for conferences, fostering an ecosystem of vendors.
- Providing neutral mediation as appropriate to try to resolve non-technical disputes that arise as part of the project.

## Schedule

| Date               | Event                                                               |
|--------------------|---------------------------------------------------------------------|
| July 18            | Announcement of Election and publication of voters.yaml             |
| July 30            | Candidate bios and voting exception forms due by 2359 UTC (5pm PDT) |
| July 31            | Election begins via Elekto UI                                       |
| August 13          | Election closes at 2359 UTC (5pm PDT)                               |
| week of August 14  | Winners announced                                                   |

## Candidacy process

Eligibility for candidacy is defined in the [Steering Committee charter](../../CHARTER.md) as a [project member](../../../ROLES.md#member) who does not work for a Company that holds a Contribution Seat.

Members from **Google**, **IBM**, **Red Hat** and **Huawei** are therefore **ineligible** to stand in this election.

If you want to stand for election, open a PR against the [istio/community repository](https://github.com/istio/community) to include
your candidate profile in the `/steering/elections/2023` folder, with the file named in the format `candidate-yourname.md`. We have included [a template file](nomination-template.md) as an example. This profile should include:

- Your name
- Your GitHub ID
- Your company affiliation (employer or otherwise)
- Your contributions to Istio
- Why you are running

Once you have created the PR, you may encourage endorsements as comments on it. After a candidate has met all election requirements, the Election Officers will merge the profile PR.

If you want to nominate someone else, you may do so, but please confirm their interest first.

## Voting process

Eligibility to vote is defined in the [Steering Committee charter](../../CHARTER.md) as either:

- a project member who has had at least one Pull Request merged in the past 12 months, or
- someone who has submitted the [voting exception form](https://forms.gle/2WmEsUbAQMz7QnwR6) and has been accepted by the steering committee as having standing in the community.

Elections will be held using [Elekto](https://elekto.io/), an online voting tool
created for the CNCF. Elekto has one critical advantage over CIVS, which is that
it uses GitHub OAuth to log you in to vote, instead of relying on email.

Thus, when you go to Elekto, you will be prompted to log in your GitHub account.
You will then be able to click on "Explore Election" to look at the list of
elections, and select the "2023 Istio Steering Committee Election."

The election page will, among other things, confirm if you are eligible to vote,
via a button at the top of the screen.

As candidates file their candidate statements in the community repo, they will
become visible in the Elekto UI.  You may click through to any candidate to see
their profile.

Once the vote begins, you will be able to rank the candidates in the order of
your preference, and submit your ballot.  When you submit, you will be offered
a chance to set a password, which is required if you want the ability to return
and re-cast your ballot before August 13th.

All data is stored on cloud infrastructure maintained by the election officers,
and is not shared with third parties. Individual ballot data is encrypted, and
not retrievable by anyone except in aggregate form.

The top four candidates who are employed by a unique Company (as defined in the
Charter) shall be deemed to be elected.

### Election officers

- [Cameron Etezadi](https://github.com/cetezadi), Google
- [Ram Vennam](https://github.com/rvennam), Solo.io
