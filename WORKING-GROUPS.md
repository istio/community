# Istio Working Groups

Most community activity is organized into *working groups*.

* [Working group meetings](#working-group-meetings)
* [Working group leads](#working-group-leads)
* [Getting in touch](#getting-in-touch)

Working groups follow the [contributing](CONTRIBUTING.md) guidelines although each of these groups may operate a little differently depending on
their needs and workflow.

When the need arises, a new working group can be created, please post to [technical-oversight-committee](https://discuss.istio.io/c/technical-oversight-committee)
working group if you think a new group is necessary.

The working groups generate design docs which are kept in a [Team Drive](https://drive.google.com/drive/u/0/folders/0AIS5p3eW9BCtUk9PVA).
Anybody can access the team drive for reading and commenting. To get access simply join the
[istio-team-drive-access@](https://groups.google.com/forum/#!forum/istio-team-drive-access) group.
Once you've done that, head to [Team Drive](https://drive.google.com/corp/drive/u/0/folders/0AIS5p3eW9BCtUk9PVA) and
behold all the docs.

The current working groups are:

| Group | Design Docs | Discussion Forum | Slack Channel | Meeting Notes | Meeting Link | Meeting Recordings | Description
|-------|-------------|-------|-------|---------------|--------------|--------------------|------------
| Config | [Folder](https://drive.google.com/corp/drive/folders/0B5CC9KT63DznUUQtSU9HTHBnb1E) | [Forum](https://discuss.istio.io/c/config) | [#config](https://istio.slack.com/messages/C7KSV4AHJ/) | [Notes](https://docs.google.com/document/d/1P3p7zOpX66hPoZBi_CiC36JW7JmoaLWqE2sgHvdq5tY/edit?ts=5a0b7200) | [Hangouts Meet](https://meet.google.com/uqh-cyie-hbp?hs=122) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffOC_hPQhaIDf8v1TxMboJj) |  Config API, config format, config management server/storage, config distribution/rollout, API documentation/style guide/governance
| Docs | [Folder](https://drive.google.com/corp/drive/u/0/folders/1C6X-UyN008fjBrGcWmRGxnliq0Jkpbih) | [Forum](https://discuss.istio.io/c/contributors/docs) | [#docs](https://istio.slack.com/messages/C50V5EATT/) | [Notes](https://docs.google.com/document/d/1RGb0NOp0J9QSIrMrZ6wokl16RbSfKHy-6NyTpOQdPa8/edit#heading=h.xjlp01fjb1kv) | [Hangouts Meet](https://meet.google.com/tfn-mmct-ped?hs=122) | n/a | User docs, information architecture, istio.io infrastructure
| Environments | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8JQWtfaS1MVk1pOHc) | [Forum](https://discuss.istio.io/c/environment) | [#environments](https://istio.slack.com/messages/C6KA8TTSS/) | [Notes](https://docs.google.com/document/d/1Ot9AeoiNYnI3fbQrq3w_-cyGxOqS8AD0RChkQfVxyhs/edit) | [Hangouts Meet](https://meet.google.com/pzi-ctet-ckx) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfelYS1XmTO1IaX4crk79tye) |  Raw VM support, Hybrid Mesh, Mac/Windows support, Cloud Foundry integration
| Networking | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8Jb1QwTXl1WTRqWHM) | [Forum](https://discuss.istio.io/c/networking) | [#networking](https://istio.slack.com/messages/C38CF1PEC/) | [Notes](https://docs.google.com/document/d/1xHy2jQ8oiwMponMVY2zJr2eUAmHW_Hi9JK42a7cg5Pc/edit#heading=h.2ju2wl4o5jbc) | [Hangouts Meet](https://meet.google.com/xjj-ujhi-qfk?hs=122) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffpmIKb5tVthiiQLSXenCK-) | Traffic Management, TCP Support, Additional L7 protocols, Proxy injection
| Performance and Scalability | [Folder](https://drive.google.com/corp/drive/u/0/folders/1Zpi5TcBPSqGno96WAq5QkcNKi16opaTL) | [Forum](https://discuss.istio.io/c/performance-and-scalability) | [#perf-and-scalability](https://istio.slack.com/messages/C7S85G0SY/) | [Notes](https://goo.gl/ENFQWb) | [Hangouts Meet](https://meet.google.com/njo-sthe-fqv?hs=122) | n/a | Performance and scalability characterization and improvements
| Policies and Telemetry | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8JbWtHOEo1STc1dGM) | [Forum](https://discuss.istio.io/c/policies-and-telemetry) | [#policies-telemetry](https://istio.slack.com/messages/C382V8Q92/) | [Notes](https://docs.google.com/document/d/1pn9QdRcoyT_nxOwzklsiYpt7OQraaSDfmtN14XTOrN0/edit) | [Hangouts Meet](https://meet.google.com/nmr-smbu-rtk?hs=122) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffF-9nyaw01Ni_0GOWBzaF4) | Mixer, Mixer Adapters, Rate Limiting, Tracing, Monitoring, Logging
| Security | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8Jb2hhRWQ2eTJYVzQ) | [Forum](https://discuss.istio.io/c/security) | [#security](https://istio.slack.com/messages/C3TEGNZ7W/) | [Notes](https://docs.google.com/document/d/1-Z7nSDV8f9psp5chO6H9PAUAinIVapOhQRaUIxbf0dk/edit#heading=h.ppfvkxi2bzdh) | [Hangouts Meet](https://meet.google.com/sxa-codd-cza?hs=122) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfd8ZbUNlZe-RYQKYTxTNJTR) | Service-to-service Auth, Identity/CA/SecretStore plugins, Identity Federation, End User Auth, Authority Delegation, Auditing
| Test and Release |[Folder](https://drive.google.com/corp/drive/u/0/folders/0B_ObhNOUiqZ0UWo1ODBVTkRkSzg) | [Forum](https://discuss.istio.io/c/test-and-release) |[#test-and-release](https://istio.slack.com/messages/C6FCV6WN4/) | [Notes](https://docs.google.com/document/d/18QgpvBH9N8Io5xU-0piysyOYif65U03m8GabpvHb4IQ/edit) | [Hangouts Meet](https://meet.google.com/rwi-kwui-axz) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfckNxJRS1ac9rqsx5b16NU_) | Build, test, release
| User Experience | [Folder](https://drive.google.com/drive/u/0/folders/1r3MDokrMU1R-jxrZh1rxBV3dYChAprJg) | [Forum](https://discuss.istio.io/c/UX) | [#user-experience](https://istio.slack.com/messages/CFTRP8NTW/) | [Notes](https://docs.google.com/document/d/1raZOoeYz3APZdQRlJlFA-zgfmEJu2YLaGa4_3KwbNRs/edit#heading=h.i71zsi7n759v) | [WebEx](https://ibm.webex.com/meet/snible) | n/a | User experience across Istio

To join Istio's Slack workspace, please fill out [this form](https://docs.google.com/forms/d/e/1FAIpQLSfdsupDfOWBtNVvVvXED6ULxtR4UIsYGCH_cQcRr0VcG1ZqQQ/viewform).

## Working group meetings

Working groups meet regularly. The meetings are open to anyone interested, so please join us. There's a number of ways to track the
meeting schedules:

* You can view the schedule of upcoming working group meetings in this [calendar](https://calendar.google.com/calendar/embed?src=4uhe8fi8sf1e3tvmvh6vrq2dog%40group.calendar.google.com&ctz=America%2FLos_Angeles).

* You can download [an ICS file](https://calendar.google.com/calendar/ical/4uhe8fi8sf1e3tvmvh6vrq2dog%40group.calendar.google.com/public/basic.ics)
to import the working group meetings into your favorite iCal client.

If you have a Google acount, you can also:

* Link the [working group event calendar](https://calendar.google.com/calendar?cid=NHVoZThmaThzZjFlM3R2bXZoNnZycTJkb2dAZ3JvdXAuY2FsZW5kYXIuZ29vZ2xlLmNvbQ)
directly to your account.

* Add yourself to the [istio-working-group-calendar-invites@](https://groups.google.com/forum/#!forum/istio-working-group-calendar-invites) group. This will
automatically invite you to all of the working group meetings on your primary Google calendar.

## Working group leads

Each working group has one or more leads which coordinate the activities of the group:

&nbsp; | Name | Company | Groups
-------|------|---------|----
&nbsp; | [Joshua Blatt](https://github.com/duderino) | Google | Test and Release
<img width="30px" src="https://avatars0.githubusercontent.com/u/2752495?s=400&v=4"> | [Frank Budinsky](https://github.com/frankbu) | IBM | Docs
&nbsp; | [Andra Cismaru](https://github.com/andraxylia) | Google | Networking
<img width="30px" src="https://avatars1.githubusercontent.com/u/5375600?s=400&v=4"> | [Spike Curtis](https://github.com/spikecurtis) | Tigera | Security
<img width="30px" src="https://avatars0.githubusercontent.com/u/755849?s=400&v=4"> | [Steven Dake](https://github.com/sdake) | NetApp | Environments
<img width="30px" src="https://avatars0.githubusercontent.com/u/13684010?s=400&v=4"> | [Surya V Duggirala](https://github.com/suryadu) | IBM | Performance and Scalability
<img width="30px" src="https://avatars2.githubusercontent.com/u/17071139?s=400&v=4"> | [Oz Evren](https://github.com/ozevren) | Google | Config
&nbsp; | [Mandar Jog](https://github.com/mandarjog) | Google | Policies and Telemetry, Performance and Scalability
<img width="30px" src="https://avatars0.githubusercontent.com/u/24381542?s=400&v=4"> | [Tao Li](https://github.com/wattli) | Google | Security
&nbsp; | [Costin Manolache](https://github.com/costinm) | Google | Environments
<img width="30px" src="https://avatars3.githubusercontent.com/u/8202871?s=400&v=4"> | [Shriram Rajagopalan](https://github.com/rshriram) | VMware | Networking
<img width="30px" src="https://avatars2.githubusercontent.com/u/21148125?s=400&v=4"> | [Douglas Reid](https://github.com/douglas-reid) | Google | Policies and Telemetry
<img width="30px" src="https://avatars3.githubusercontent.com/u/3237651?s=400&v=4"> | [Ed Snible](https://github.com/esnible) | IBM  | User Experience
<img width="30px" src="https://avatars1.githubusercontent.com/u/1588319?s=400&v=4">  | [Lin Sun](https://github.com/linsun) | IBM | Test and Release
<img width="30px" src="https://avatars3.githubusercontent.com/u/22780957?s=400&v=4"> | [Martin Taillefer](https://github.com/geeknoid) | Google  | Docs
&nbsp; | [Limin Wang](https://github.com/liminw) | Google | Security

## Getting in touch

If you'd like to reach out to the working group leads, please drop a note to [wg-leads@discuss.istio.io](mailto:wg-leads@discuss.istio.io).
