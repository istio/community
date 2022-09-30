# Contributing to Istio

So, you want to hack on Istio? Yay!

The following sections outline the process all changes to the Istio
repositories go through.  All changes, regardless of whether they are from
newcomers to the community or from the core team follow the
same process and are given the same level of review.

- [Working groups](#working-groups)
- [Code of conduct](#code-of-conduct)
- [Team values](#team-values)
- [Contributor license agreements](#contributor-license-agreements)
- [Design documents](#design-documents)
- [Contributing a feature](#contributing-a-feature)
- [Setting up to contribute to Istio](#setting-up-to-contribute-to-istio)
- [Pull requests](#pull-requests)
- [Issues](#issues)
- [Promote your company on istio.io](#promote-your-company-on-istioio)
- [Tell the world you're using Istio](#tell-the-world-youre-using-istio)

You can find additional contributor material on the [Istio Wiki](https://github.com/istio/istio/wiki).

## Working groups

The Istio community is organized into a set of [working groups](WORKING-GROUPS.md).
Any contribution to Istio should be started by first engaging with the
appropriate working group.

## Code of conduct

All members of the Istio community must abide by the [Code of Conduct](CODE-OF-CONDUCT.md).

## Team values

We promote and encourage a set of [shared values](VALUES.md) to improve our
productivity and inter-personal interactions.

## Contributor license agreements

We'd love to accept your contributions! But before we can take them, you will
have
to fill out the [Google CLA](https://cla.developers.google.com).

Once you are CLA'ed, we'll be able to accept your pull requests. This is
necessary because you own the copyright to your changes, even after your
contribution becomes part of this project. So this agreement simply gives Google
permission to use and redistribute your contributions as part of the project.

## Design documents

Any substantial design deserves a design document. Design documents are written with Google Docs and
should be shared with the community by adding the doc to our [shared Drive](https://drive.google.com/corp/drive/folders/0ADmbrU7ueGOUUk9PVA)
and sending a note to the appropriate working group to let people know the doc is there. To get write access
to the drive, you'll need to be a [member](ROLES.md#member) of the Istio organization.

Anybody can access the shared Drive for reading. To get access to comment, join the
[istio-team-drive-access@](https://groups.google.com/forum/#!forum/istio-team-drive-access) group.
Once you've done that, head to the [shared Drive](https://drive.google.com/corp/drive/folders/0ADmbrU7ueGOUUk9PVA) and
behold all the docs.

When documenting a new design, we recommend a 2-step approach:

1. Use the short-form
[RFC template](https://docs.google.com/document/d/1ewJoCcw5-04crH-M0xw4zFxz1cfwVCPnNyW4K3m4Yyc/template/preview)
to outline your ideas and get early feedback.
1. Once you have received sufficient feedback and consensus, you may use the longer-form
[design doc template](https://docs.google.com/document/d/16FLQK8uhhic1ovKnnOG3OXJjFKs2aHnSmbximidpKwM/template/preview)
to specify and discuss your design in more details.

To use either template, open the template and select "Use Template" in order to bootstrap your document.

## Contributing a feature

In order to contribute a feature to Istio you'll need to go through the following steps:

- Discuss your idea with the appropriate [working groups](WORKING-GROUPS.md) on the working
group's Slack channel.

- Once there is general agreement that the feature is useful, create a GitHub issue to track the discussion. The issue should include information
about the requirements and use cases that it is trying to address. Include a discussion of the proposed design and technical details of the
implementation in the issue.

- If the feature is substantial enough:

  - Working group leads will ask for a design document as outlined in the previous section.
  Create the design document and add a link to it in the GitHub issue. Don't forget to send a Slack to the
  working group to let everyone know your document is ready for review.

  - Depending of the breath of the design and how contentious it is, the working group leads may decide
  the feature needs to be discussed in one or more working group meetings before being approved.

  - Once the major technical issues are resolved and agreed upon, post a note to the working group's Slack
  with the design decision and the general execution plan.

- Submit PRs to [istio/istio](https://github.com/istio/istio) with your code changes.

- Submit PRs to [istio/istio.io](https://github.com/istio/istio.io) with
documentation for your feature, including usage examples when possible. See [here](https://istio.io/about/contribute/)
to learn how to write docs for istio.io.

> Note that we prefer bite-sized PRs instead of giant monster PRs. It's therefore preferable if you
can introduce large features in smaller reviewable changes that build on top of one another.

If you would like to skip the process of submitting an issue and
instead would prefer to just submit a pull request with your desired
code changes then that's fine. But keep in mind that there is no guarantee
of it being accepted and so it is usually best to get agreement on the
idea/design before time is spent coding it. However, sometimes seeing the
exact code change can help focus discussions, so the choice is up to you.

## Setting up to contribute to Istio

Check out this [README](https://github.com/istio/istio/blob/master/README.md) to learn about
the Istio source base and setting up your [development environment](https://github.com/istio/istio/wiki/Preparing-for-Development).

## Pull requests

If you're working on an existing issue, simply respond to the issue and express
interest in working on it. This helps other people know that the issue is
active, and hopefully prevents duplicated efforts.

To submit a proposed change:

- Fork the affected repository.

- Create a new branch for your changes.

- Develop the code/fix.

- Add new test cases. In the case of a bug fix, the tests should fail
  without your code changes. For new features try to cover as many
  variants as reasonably possible.

- Modify the documentation as necessary.

- Verify the entire CI process (building and testing) works.

While there may be exceptions, the general rule is that all PRs should
be 100% complete - meaning they should include all test cases and documentation
changes related to the change.

When ready, if you have not already done so, sign a
[contributor license agreements](#contributor-license-agreements) and submit
the PR.

See [Writing Good Pull Requests](https://github.com/istio/istio/wiki/Writing-Good-Pull-Requests) for guidance on creating
pull requests, and [Reviewing Pull Requests](https://github.com/istio/istio/wiki/Reviewing-Pull-Requests) for the PR review
we use.

## Issues

[GitHub issues](https://github.com/istio/istio/issues/new/choose) can be used to report bugs or submit feature requests.

When reporting a bug please include the following key pieces of information:

- The version of the project you were using (e.g. version number,
  or git commit)

- Operating system you are using.

- The exact, minimal, steps needed to reproduce the issue.
  Submitting a 5 line script will get a much faster response from the team
  than one that's hundreds of lines long.

## Promote your company on istio.io

If your company supports Istio, you can list it on our [Ecosystem page](https://istio.io/latest/about/ecosystem/).
We have categories for *providers* (who offer hosted or managed Istio services
to their customers), *professional services* (who offer implementation support
or consulting for Istio) and *integrations* (commercial or open source products
that work with Istio).

To add an entry to this page, edit [data/companies.yml](https://github.com/istio/istio.io/blob/master/data/companies.yml)
and add your information in the `providers`, `pro_services` or `integrations`
section, as appropriate. Please add your company logo, preferably in SVG
format, to [static/logos](https://github.com/istio/istio.io/blob/master/static/logos).

The Istio steering committee reserves the right to remove logos of companies that
are not demeed to be in good standing in the community.

## Tell the world you're using Istio

Are you running Istio in production?  List yourself on our [Case Studies](https://istio.io/latest/about/case-studies/) page.

If you want to add your logo to our wall of users, add an entry to the `users`
section in [data/companies.yml](https://github.com/istio/istio.io/blob/master/data/companies.yml), and add your logo,
preferably in SVG format, to [static/logos](https://github.com/istio/istio.io/blob/master/static/logos).

If you'd like to be interviewed for a full case study, please [create an issue](https://github.com/istio/istio.io/issues/new),
and we'll be in touch!
