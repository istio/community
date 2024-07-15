# Istio Working Group Processes

This document describes the processes we use to manage the Istio working groups. This includes how they are formed, how leads
are established, how they are run, etc.

* [Why working groups?](#why-working-groups)
* [Proposing a new working group](#proposing-a-new-working-group)
* [Setting up a working group](#setting-up-a-working-group)
  * [Dissolving a working group](#dissolving-a-working-group)
* [Running a working group](#running-a-working-group)
  * [Be open](#be-open)
  * [Making decisions](#making-decisions)
  * [Subgroups](#subgroups)
  * [Escalations](#escalations)

## Why working groups?

Istio working groups are organizations responsible for the design and implementation of large architectural aspects of the overall Istio project.
Working groups operate with a fair amount of autonomy within the broader scope of the project. They tend to be long-lived, representing major
initiatives over Istio’s lifetime.

Some working groups focus on specific technologies. For example, the Extensions and Telemetry working group focuses on WebAssembly based
extensibility and extensions for features such as Rate Limiting, Tracing, Monitoring, Logging.

The [Technical Oversight Committee](TECH-OVERSIGHT-COMMITTEE.md) (a.k.a. TOC) is responsible for the Istio project as a whole. It sets the overall direction
of the project, helps make crosscutting architectural decisions, helps establish and dissolve working groups, and helps ensure all working
groups are generally rowing in the same direction.

Although working groups are relatively lightweight structures, we want to keep the number of working groups low in order to keep things
manageable.

## Proposing a new working group

If you’ve identified a substantial architectural area which would benefit from long-lived, concerted and focused design, then you should
consider creating a new working group. To do so, you need to:

* **Create a charter**. This should be a few paragraphs explaining:

    * The mission of the working group

    * The goals of the working group (problems being solved)

    * The scope of the working group (topics, subsystems, code repos, areas of responsibility)

* **Nominate an initial set of leads**. The leads set the agenda for the working group and serve as final arbiters on any technical decision.
See [below](#running-a-working-group) for information on the responsibilities of leads and requirements for nominating them.

* **Prepare a Roadmap**. Create a preliminary 3 month roadmap for what the working group would focus on.

* **Send an Email**. Write up an email with your charter, nominated leads, and roadmap, and post it to #toc-steering-questions on the [Istio Slack](https://slack.istio.io/).
The technical oversight committee will evaluate the request and decide whether the working group should be
formed, whether it should be merely a subgroup of an existing working group, or whether it should be subsumed by an existing working group.

## Setting up a working group

Once approval has been granted by the technical oversight committee to form a working group, the working group leads need to take a few
steps to establish the working group:

* **Create a Google Drive Folder**. Create a folder to hold your working group documents within this parent
[folder](https://drive.google.com/drive/folders/15y-2OIcjP9DNF7WLvHL0Hyn_XXxaYoOm). Call your folder "GROUP_NAME".

* **Create a Roadmap Document**. Create a document in the above folder and call it "GROUP_NAME Group Roadmap". Put your initial
roadmap in the document.

* **Register the Working Group**. Go to [WORKING-GROUPS.md](https://github.com/istio/community/blob/master/WORKING-GROUPS.md) and
add your working group name, the names of the leads, the working group charter, and a link to the meeting you created.

* **Announce your Working Group**. Post a note to #contributors on the Istio Slack (https://slack.istio.io/)
to announce your new working group. Include your charter in the post.

* **Attend the Combined Working Group Meeting**. Check the meeting agenda to verify if something was added relative to your working group.
It is desirable for the leads to attend the weekly meeting in case items do come up.

Congratulations, you now have a fully formed working group!

### Dissolving a working group

Some working groups are ephemeral or naturally reach the end of their useful life. Working group leads can petition to dissolve their working groups in a TOC meeting. The
technical oversight committee also reserves the right to dissolve or recharter working groups over time as necessary.

## Leads

Each working group must have 2 or 3 leads, with at least two different organizations represented. Working group leads must meet certain requirements, including having been members of the Istio project for a year, and a maintainer for at least 3 months.

Please see the [Community Roles](https://github.com/istio/community/blob/master/ROLES.md) document for a full
description of a lead’s role and requirements.

## Running a working group

Leads are responsible for running a working group. Running the group involves a few activities:

* **Meetings**. Help with running the combined working group meetings. Be sure to attend the meeting if there is an agenda item
for your workgroup. Ensure the meetings are recorded, and properly archived on YouTube.

* **Notes**. Ensure that meeting notes are kept up to date. The lead may delegate note-taking duties.
Provide a link to the recorded meeting in the notes.

* **Roadmap**. Establish **and maintain** a roadmap for the working group outlining the areas of focus for the working group over the next
3 months.

* **Report**. Report current status to the TOC meeting every as appropriate.

### Be open

The community design process is done in the open. Working groups should communicate primarily through the public
working group Slack channel and meetings, through design documents in the working group’s folder, through GitHub issues, and GitHub PRs.
Avoid private emails and/or meetings when possible.

### Making decisions

In general, working groups operate in a highly cooperative environment. Working groups discuss designs in the open and take input
from the community at large when making technical choices. The working group leads are ultimately responsible for setting the
direction of the working group and making the tough technical choices affecting the working group.

### Subgroups

Subgroups are ad hoc subteams within a working group with a special focus on a set of problems or technologies. We don’t formalize processes
for subgroups, each working group can decide when subgroups are needed and how they operate.

### Escalations

Working groups can get blocked on specific technical disagreements. Leads are expected to generally resolve such issues and allow work
to progress.

Sometimes, different working groups can have conflicting goals or requirements. Leads from all affected working groups generally work
together and come to an agreeable conclusion.

If a community member disagrees with decisions or designs made by a working group, the first recourse is to consult the working group
leads and try to hash things out. If that's not satisfactory, then the community member is invited to prepare an alternate proposal
that can be shared with the community.

In all cases, remaining blocking issues can be raised to the [technical oversight committee](TECH-OVERSIGHT-COMMITTEE.md) to help
resolve the situation. To trigger an escalation, add an entry to the [TOC's agenda](https://docs.google.com/document/d/13lxJqtlaQhmV2EwsNnS6h-_O4pobZQZuMjrzOeMgVI0/edit#heading=h.o8pz6aqnzzgk)
and come to a TOC meeting. Come prepared with an explanation of the underlying problem as well as
cogent arguments for both sides.
