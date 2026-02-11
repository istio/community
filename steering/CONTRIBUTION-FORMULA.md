# Contribution seat allocation formula

The [steering committee charter](CHARTER.md) states "Each year, when
the Steering Committee term begins, the Steering Committee will vote on an exact
formula and procedure for determining allocation for the following year."

For the February 2026 to February 2027 term, the agreed contribution formula was
**[contributions][contrib]** (being a review, commit, discussion, issue or PR on GitHub) attributed to a Contributing Company,
as measured by [DevStats](https://istio.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=istio&var-country_name=All&var-companies=All),
filtered to the `istio` repository group,
and the [D'Hondt method](https://en.wikipedia.org/wiki/D%27Hondt_method) of allocation of seats.

Attribution is done using the [cncf/gitdm](https://github.com/cncf/gitdm) developer affiliation files.
Changes to attribution can be made by Pull Request to that repository.

[The method and results for prior terms are published here](https://drive.google.com/drive/folders/1sRkDGN591L_bau5RCw7ZFBhdTO4pJj-F).

   [contrib]: https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-github-profile/managing-contribution-settings-on-your-profile/viewing-contributions-on-your-profile#what-counts-as-a-contribution
