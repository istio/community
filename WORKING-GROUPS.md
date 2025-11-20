# Istio Working Groups

Most community activity is organized into *working groups*.

- [Istio Working Groups](#istio-working-groups)
  - [Working group meetings](#working-group-meetings)
  - [Working group leads](#working-group-leads)
  - [Getting in touch](#getting-in-touch)

Working groups follow the [contributing](CONTRIBUTING.md) guidelines although each of these groups may operate a little differently depending on
their needs and workflow.

When the need arises, a new working group can be created. Please ask in #toc-steering-questions in the [Istio Slack](https://slack.istio.io/) if you think a new group is necessary.

The working groups generate design docs which are kept in a shared Google Drive.
Anybody can access the drive for reading and commenting. To get access simply join the
[istio-team-drive-access@](https://groups.google.com/forum/#!forum/istio-team-drive-access) group.
Once you've done that, head to the [team drive](https://drive.google.com/drive/folders/1l_zqgBq_yfc1PfbJiWsFubXBtAz22sau) and
behold all the docs.

The below links to Slack require you to first be logged into Istio's Slack workspace. Please use [this link](https://slack.istio.io/) if you need to join.

The current working groups are:

| Group | Design Docs | Slack Channel | Description
|-------|-------------|---------------|------------
| Docs | [Folder](https://drive.google.com/drive/folders/16Alb6m30ypLbz3iVv2OTjpqUHO8NJIUv) | [#docs](https://istio.slack.com/messages/C50V5EATT/) | User docs, information architecture, istio.io infrastructure
| Environments | [Folder](https://drive.google.com/drive/folders/16P5sCnSzEk204LHBKYTkC9jEmgzBJy6x) | [#environments](https://istio.slack.com/messages/C6KA8TTSS/) |  Raw VM support, Hybrid Mesh, Mac/Windows support, Cloud Foundry integration
| Networking | [Folder](https://drive.google.com/drive/folders/16RDUAJj_LnJM83weyk0WSCeqgtTw8Ogz) | [#networking](https://istio.slack.com/messages/C38CF1PEC/) | Traffic Management, TCP Support, Additional L7 protocols, Proxy injection
| Extensions and Telemetry | [Folder](https://drive.google.com/drive/folders/16Jl2ASNQ226h5XmXSAtxx78cs1iSEAYQ) | [#extensions-telemetry](https://istio.slack.com/messages/C382V8Q92/) | WebAssembly based extensibility, Istio extensions for features such as Rate Limiting, Tracing, Monitoring, Logging
| Product Security | [Folder](https://drive.google.com/drive/folders/16gnkbCgO18zkMIk_IHrE_QV2EFqFWe4R) | [Report a vulnerability](https://istio.io/about/security-vulnerabilities/) | Product Security: Vulnerability, security guidelines, threats
| Security | [Folder](https://drive.google.com/drive/folders/16eF1aoVnknX1vY853gEIPfoO2-i-FDcU) | [#security](https://istio.slack.com/messages/C3TEGNZ7W/) | Service-to-service Auth, Identity/CA/SecretStore plugins, Identity Federation, End User Auth, Authority Delegation, Auditing
| Test and Release |[Folder](https://drive.google.com/drive/folders/16M0Ba8uT6-1F71NARhddXDdnk-EE4jEn) | [#test-and-release](https://istio.slack.com/messages/C6FCV6WN4/) | Build, test, release
| User Experience | [UX](https://drive.google.com/drive/folders/16CplcqFZTT-5WXDkpi443-ly2RnnA3C5) | [#user-experience](https://istio.slack.com/messages/CFTRP8NTW/) | User experience across Istio, API and CLI guidelines and support

## Working group meetings

Currently all working groups share a weekly meeting, on Wednesdays at 9:00 AM PST. The meetings are open to anyone interested,
so please join us. The meeting can be joined with [this](https://meet.google.com/qza-pfbq-wne) link, and the notes can be found
[here](https://docs.google.com/document/d/1wsa06GGiq1LEGwhkiPP0FKIZJqdAiue-VeBonWAzAyk/edit).

There's a number of ways to track the
meeting schedule:

* You can view the schedule of upcoming working group meetings in this [calendar](https://calendar.google.com/calendar/u/1?cid=Y19mZTljZDhkMDkxZDBkNGQyODE4ZDAxYTkzMGRjOTI1ZjAwZDRmOTc5OTVlZmU1MGE4ZDcyNTEyYjI0MTU2OGY1QGdyb3VwLmNhbGVuZGFyLmdvb2dsZS5jb20).

* You can download [an ICS file](https://calendar.google.com/calendar/ical/c_fe9cd8d091d0d4d2818d01a930dc925f00d4f97995efe50a8d72512b241568f5%40group.calendar.google.com/public/basic.ics)
to import the working group meetings into your favorite iCal client.

If you have a Google account, you can also:

* Link the [working group event calendar](https://calendar.google.com/calendar/u/1?cid=Y19mZTljZDhkMDkxZDBkNGQyODE4ZDAxYTkzMGRjOTI1ZjAwZDRmOTc5OTVlZmU1MGE4ZDcyNTEyYjI0MTU2OGY1QGdyb3VwLmNhbGVuZGFyLmdvb2dsZS5jb20)
directly to your account.

* Add yourself to the [istio-working-group-calendar-invites@](https://groups.google.com/forum/#!forum/istio-working-group-calendar-invites) group. This will
automatically invite you to all of the working group meetings on your primary Google calendar.

## Working group leads

Each working group has one or more leads which coordinate the activities of the group:

&nbsp;                                                                               | Name                                               | Company    | Groups
-------------------------------------------------------------------------------------|----------------------------------------------------|------------|-------
&nbsp;                                                                               | [Rama Chavali](https://github.com/ramaraochavali)  | Salesforce | Networking - Control Plane
<img width="30px" src="https://avatars3.githubusercontent.com/u/821270?s=80&v=4">    | [Mitch Connors](https://github.com/therealmitchconnors) | Microsoft   | User Experience
<img width="30px" src="https://avatars.githubusercontent.com/u/1330566?v=4">         | [Daniel Hawton](https://github.com/dhawton)        | Solo.io    | Product Security
&nbsp;                                                                               | [Jianpeng He](https://github.com/zirain)           | Tetrate    | Extensions and Telemetry
<img width="30px" src="https://avatars1.githubusercontent.com/u/623453?s=400&v=4">   | [John Howard](https://github.com/howardjohn)       | Solo.io    | Networking
&nbsp;                                                                               | [Steven Landow](https://github.com/stevenctl)      | Solo.io    | Networking - Control Plane
<img width="30px" src="https://avatars.githubusercontent.com/u/64559656?v=4">        | [Jackie Maertens (Elliot)](https://github.com/jaellio)        | Microsoft  | Product Security
&nbsp;                                                                               | [Costin Manolache](https://github.com/costinm)     | Google     | Environments
&nbsp;                                                                               | [Sam Naser](https://github.com/monkeyanator)       | Aviatrix   | Environments
<img width="30px" src="https://avatars3.githubusercontent.com/u/3237651?s=400&v=4">  | [Ed Snible](https://github.com/esnible)            | IBM        | User Experience
&nbsp;                                                                               | [Lei Tang](https://github.com/lei-tang)            | Google     | Security, Extensions and Telemetry
<img width="30px" src="https://avatars1.githubusercontent.com/u/10537847?s=400&v=4"> | [Eric Van Norman](https://github.com/ericvn)       | IBM        | Test and Release
<img width="30px" src="https://avatars0.githubusercontent.com/u/1016047?s=400&v=4">  | [Lizan Zhou](https://github.com/lizan)             | Aviatrix   | Networking - Data Plane, Security
