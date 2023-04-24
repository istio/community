# Istio Working Groups

Most community activity is organized into *working groups*.

* [Working group meetings](#working-group-meetings)
* [Working group leads](#working-group-leads)
* [Getting in touch](#getting-in-touch)

Working groups follow the [contributing](CONTRIBUTING.md) guidelines although each of these groups may operate a little differently depending on
their needs and workflow.

When the need arises, a new working group can be created, please post to [technical-oversight-committee](https://discuss.istio.io/c/technical-oversight-committee)
working group if you think a new group is necessary.

The working groups generate design docs which are kept in a shared Google Drive.
Anybody can access the drive for reading and commenting. To get access simply join the
[istio-team-drive-access@](https://groups.google.com/forum/#!forum/istio-team-drive-access) group.
Once you've done that, head to the [Community Drive](https://drive.google.com/drive/folders/0ADmbrU7ueGOUUk9PVA) and
behold all the docs.

The current working groups are:

| Group | Design Docs | Discussion Forum | Slack Channel | Meeting Notes | Meeting Link | Meeting Recordings | Description
|-------|-------------|-------|-------|---------------|--------------|--------------------|------------
| Docs | [Folder](https://drive.google.com/drive/folders/16Alb6m30ypLbz3iVv2OTjpqUHO8NJIUv) | [Forum](https://discuss.istio.io/c/contributors/docs) | [#docs](https://istio.slack.com/messages/C50V5EATT/) | [Notes](https://docs.google.com/document/d/1RGb0NOp0J9QSIrMrZ6wokl16RbSfKHy-6NyTpOQdPa8/edit#heading=h.xjlp01fjb1kv) | [Hangouts Meet](https://meet.google.com/tfn-mmct-ped?hs=122) | n/a | User docs, information architecture, istio.io infrastructure
| Environments | [Folder](https://drive.google.com/drive/folders/16P5sCnSzEk204LHBKYTkC9jEmgzBJy6x) | [Forum](https://discuss.istio.io/c/environment) | [#environments](https://istio.slack.com/messages/C6KA8TTSS/) | [Notes](https://docs.google.com/document/d/1Ot9AeoiNYnI3fbQrq3w_-cyGxOqS8AD0RChkQfVxyhs/edit) | [Hangouts Meet](https://meet.google.com/pzi-ctet-ckx) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfelYS1XmTO1IaX4crk79tye) |  Raw VM support, Hybrid Mesh, Mac/Windows support, Cloud Foundry integration
| Networking | [Folder](https://drive.google.com/drive/folders/16RDUAJj_LnJM83weyk0WSCeqgtTw8Ogz) | [Forum](https://discuss.istio.io/c/networking) | [#networking](https://istio.slack.com/messages/C38CF1PEC/) | [Notes](https://docs.google.com/document/d/1xHy2jQ8oiwMponMVY2zJr2eUAmHW_Hi9JK42a7cg5Pc/edit#heading=h.2ju2wl4o5jbc) | [Hangouts Meet](https://meet.google.com/xjj-ujhi-qfk?hs=122) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffpmIKb5tVthiiQLSXenCK-) | Traffic Management, TCP Support, Additional L7 protocols, Proxy injection
| Extensions and Telemetry | [Folder](https://drive.google.com/drive/folders/16Jl2ASNQ226h5XmXSAtxx78cs1iSEAYQ) | [Forum](https://discuss.istio.io/c/policies-and-telemetry) | [#extensions-telemetry](https://istio.slack.com/messages/C382V8Q92/) | [Notes](https://docs.google.com/document/d/1pn9QdRcoyT_nxOwzklsiYpt7OQraaSDfmtN14XTOrN0/edit) | [Main Group Hangouts Meet](https://meet.google.com/cmo-jgug-uhj), [Wasm SIG Hangout Meet](https://meet.google.com/yrs-qhtf-mkb) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffF-9nyaw01Ni_0GOWBzaF4) | WebAssembly based extensibility, Istio extensions for features such as Rate Limiting, Tracing, Monitoring, Logging
| Product Security | [Folder](https://drive.google.com/drive/folders/16gnkbCgO18zkMIk_IHrE_QV2EFqFWe4R) | [Report a vulnerability](https://istio.io/about/security-vulnerabilities/) |  | [Notes](https://docs.google.com/document/d/1fGi8G9wGfJhA23CT9Awp2bBdasanILh3gK9ib5hrTlE/edit#) | [Hangouts Meet](https://meet.google.com/vao-otzc-hvx) | | Product Security: Vulnerability, security guidelines, threats
| Security | [Folder](https://drive.google.com/drive/folders/16eF1aoVnknX1vY853gEIPfoO2-i-FDcU) | [Forum](https://discuss.istio.io/c/security) | [#security](https://istio.slack.com/messages/C3TEGNZ7W/) | [Notes](https://docs.google.com/document/d/12Xz8fCuql2tPL-cpdLXGGE44kYT4o9NB6mljSiS_01U/edit#heading=h.o8pz6aqnzzgk) | [Hangouts Meet](https://meet.google.com/aop-xfen-ynm) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfd8ZbUNlZe-RYQKYTxTNJTR) | Service-to-service Auth, Identity/CA/SecretStore plugins, Identity Federation, End User Auth, Authority Delegation, Auditing
| Test and Release |[Folder](https://drive.google.com/drive/folders/16M0Ba8uT6-1F71NARhddXDdnk-EE4jEn) | [Forum](https://discuss.istio.io/c/test-and-release) |[#test-and-release](https://istio.slack.com/messages/C6FCV6WN4/) | [Notes](https://docs.google.com/document/d/18QgpvBH9N8Io5xU-0piysyOYif65U03m8GabpvHb4IQ/edit) | [Hangouts Meet](https://meet.google.com/bis-vvyb-vof) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfckNxJRS1ac9rqsx5b16NU_) | Build, test, release
| User Experience | [UX](https://drive.google.com/drive/folders/16CplcqFZTT-5WXDkpi443-ly2RnnA3C5) [_Config (old)_](https://drive.google.com/drive/folders/16W9uDkzrKJM-E5R7_5TGz-PMsAkkoT1s) | [Forum](https://discuss.istio.io/c/UX) | [#user-experience](https://istio.slack.com/messages/CFTRP8NTW/) [#config](https://istio.slack.com/messages/C7KSV4AHJ/)| [Notes](https://docs.google.com/document/d/1raZOoeYz3APZdQRlJlFA-zgfmEJu2YLaGa4_3KwbNRs/edit#heading=h.i71zsi7n759v) | [WebEx](https://ibm.webex.com/meet/snible) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfcaVTlYkvV3e09Gak6zhaLj) | User experience across Istio, API and CLI guidelines and support

_NOTE: Config working group responsibilities have been merged into User Experience._

To join Istio's Slack workspace, please use [this link](https://slack.istio.io/).

## Working group meetings

There is a weekly combined workgroup meeting. The meetings are open to anyone interested, so please join us. Here are pointers to the [Google Doc](https://docs.google.com/document/d/1wsa06GGiq1LEGwhkiPP0FKIZJqdAiue-VeBonWAzAyk/edit) as well as the [Hangouts Meet](https://docs.google.com/document/d/1wsa06GGiq1LEGwhkiPP0FKIZJqdAiue-VeBonWAzAyk/edit) for the combined workgroup meeting.

There's a number of ways to track the
meeting schedule:

* You can view the schedule of upcoming working group meetings in this [calendar](https://calendar.google.com/calendar/embed?src=4uhe8fi8sf1e3tvmvh6vrq2dog%40group.calendar.google.com).

* You can download [an ICS file](https://calendar.google.com/calendar/ical/4uhe8fi8sf1e3tvmvh6vrq2dog%40group.calendar.google.com/public/basic.ics)
to import the working group meetings into your favorite iCal client.

If you have a Google account, you can also:

* Link the [working group event calendar](https://calendar.google.com/calendar?cid=NHVoZThmaThzZjFlM3R2bXZoNnZycTJkb2dAZ3JvdXAuY2FsZW5kYXIuZ29vZ2xlLmNvbQ)
directly to your account.

* Add yourself to the [istio-working-group-calendar-invites@](https://groups.google.com/forum/#!forum/istio-working-group-calendar-invites) group. This will
automatically invite you to all of the working group meetings on your primary Google calendar.

## Working group leads

Each working group has one or more leads which coordinate the activities of the group:

&nbsp;                                                                               | Name                                               | Company    | Groups
-------------------------------------------------------------------------------------|----------------------------------------------------|------------|-------
&nbsp;                                                                               | [Rama Chavali](https://github.com/ramaraochavali)  | Salesforce | Networking - Control Plane
<img width="30px" src="https://avatars3.githubusercontent.com/u/821270?s=80&v=4">    | [Mitch Connors](https://github.com/therealmitchconnors) | Aviatrix  | User Experience
<img width="30px" src="https://avatars3.githubusercontent.com/u/38300436?s=400&v=4"> | [Jacob Delgado](https://github.com/jacob-delgado)  | Aspen Mesh | Product Security, Environments
<img width="30px" src="https://avatars1.githubusercontent.com/u/623453?s=400&v=4">   | [John Howard](https://github.com/howardjohn)       | Google     | Networking
&nbsp;                                                                               | [Steven Landow](https://github.com/stevenctl)      | Solo.io    | Networking - Control Plane
&nbsp;                                                                               | [Lei Tang](https://github.com/lei-tang)            | Google     | Security, Product Security, Extensions and Telemetry
&nbsp;                                                                               | [Shankar Ganesan](https://github.com/shankgan)     | Google     | Product Security
&nbsp;                                                                               | [Costin Manolache](https://github.com/costinm)     | Google     | Environments
&nbsp;                                                                               | [Sam Naser](https://github.com/monkeyanator)       | Google     | Environments
<img width="30px" src="https://avatars3.githubusercontent.com/u/3237651?s=400&v=4">  | [Ed Snible](https://github.com/esnible)            | IBM        | User Experience
<img width="30px" src="https://avatars1.githubusercontent.com/u/10537847?s=400&v=4"> | [Eric Van Norman](https://github.com/ericvn)       | IBM        | Test and Release
&nbsp;                                                                               | [Limin Wang](https://github.com/liminw)            | Google     | Security
<img width="30px" src="https://avatars0.githubusercontent.com/u/1016047?s=400&v=4">  | [Lizan Zhou](https://github.com/lizan)             | Tetrate    | Networking - Data Plane, Security
<img width="30px" src="https://avatars.githubusercontent.com/u/19473391?s=400&v=4">  | [Greg Hanson](https://github.com/GregHanson)       | Solo.io     | Product Security
&nbsp;                                                                               | [Jianpeng He](https://github.com/zirain)           | Huawei     | Extensions and Telemetry
## Getting in touch

If you'd like to reach out to the working group leads, please drop a note to [wg-leads@discuss.istio.io](mailto:wg-leads@discuss.istio.io).
