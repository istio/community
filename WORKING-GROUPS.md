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
| Docs | [Folder](https://drive.google.com/corp/drive/u/0/folders/1C6X-UyN008fjBrGcWmRGxnliq0Jkpbih) | [Forum](https://discuss.istio.io/c/contributors/docs) | [#docs](https://istio.slack.com/messages/C50V5EATT/) | [Notes](https://docs.google.com/document/d/1RGb0NOp0J9QSIrMrZ6wokl16RbSfKHy-6NyTpOQdPa8/edit#heading=h.xjlp01fjb1kv) | [Hangouts Meet](https://meet.google.com/tfn-mmct-ped?hs=122) | n/a | User docs, information architecture, istio.io infrastructure
| Environments | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8JQWtfaS1MVk1pOHc) | [Forum](https://discuss.istio.io/c/environment) | [#environments](https://istio.slack.com/messages/C6KA8TTSS/) | [Notes](https://docs.google.com/document/d/1Ot9AeoiNYnI3fbQrq3w_-cyGxOqS8AD0RChkQfVxyhs/edit) | [Hangouts Meet](https://meet.google.com/pzi-ctet-ckx) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfelYS1XmTO1IaX4crk79tye) |  Raw VM support, Hybrid Mesh, Mac/Windows support, Cloud Foundry integration
| Networking | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8Jb1QwTXl1WTRqWHM) | [Forum](https://discuss.istio.io/c/networking) | [#networking](https://istio.slack.com/messages/C38CF1PEC/) | [Notes](https://docs.google.com/document/d/1xHy2jQ8oiwMponMVY2zJr2eUAmHW_Hi9JK42a7cg5Pc/edit#heading=h.2ju2wl4o5jbc) | [Hangouts Meet](https://meet.google.com/xjj-ujhi-qfk?hs=122) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffpmIKb5tVthiiQLSXenCK-) | Traffic Management, TCP Support, Additional L7 protocols, Proxy injection
| Performance and Scalability | [Folder](https://drive.google.com/corp/drive/u/0/folders/1Zpi5TcBPSqGno96WAq5QkcNKi16opaTL) | [Forum](https://discuss.istio.io/c/performance-and-scalability) | [#perf-and-scalability](https://istio.slack.com/messages/C7S85G0SY/) | [Notes](https://goo.gl/ENFQWb) | [Hangouts Meet](https://meet.google.com/njo-sthe-fqv) | n/a | Performance and scalability characterization and improvements
| Policies and Telemetry | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8JbWtHOEo1STc1dGM) | [Forum](https://discuss.istio.io/c/policies-and-telemetry) | [#policies-telemetry](https://istio.slack.com/messages/C382V8Q92/) | [Notes](https://docs.google.com/document/d/1pn9QdRcoyT_nxOwzklsiYpt7OQraaSDfmtN14XTOrN0/edit) | [Hangouts Meet](https://meet.google.com/cmo-jgug-uhj) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffF-9nyaw01Ni_0GOWBzaF4) | Istio extensions for features such as Rate Limiting, Tracing, Monitoring, Logging
| Security | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8Jb2hhRWQ2eTJYVzQ) | [Forum](https://discuss.istio.io/c/security) | [#security](https://istio.slack.com/messages/C3TEGNZ7W/) | [Notes](https://docs.google.com/document/d/12Xz8fCuql2tPL-cpdLXGGE44kYT4o9NB6mljSiS_01U/edit#heading=h.o8pz6aqnzzgk) | [Hangouts Meet](https://meet.google.com/aop-xfen-ynm) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfd8ZbUNlZe-RYQKYTxTNJTR) | Service-to-service Auth, Identity/CA/SecretStore plugins, Identity Federation, End User Auth, Authority Delegation, Auditing
| Test and Release |[Folder](https://drive.google.com/corp/drive/u/0/folders/0B_ObhNOUiqZ0UWo1ODBVTkRkSzg) | [Forum](https://discuss.istio.io/c/test-and-release) |[#test-and-release](https://istio.slack.com/messages/C6FCV6WN4/) | [Notes](https://docs.google.com/document/d/18QgpvBH9N8Io5xU-0piysyOYif65U03m8GabpvHb4IQ/edit) | [Hangouts Meet](https://meet.google.com/bis-vvyb-vof) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfckNxJRS1ac9rqsx5b16NU_) | Build, test, release
| User Experience | [UX](https://drive.google.com/drive/u/0/folders/1r3MDokrMU1R-jxrZh1rxBV3dYChAprJg) [_Config(old)_](https://drive.google.com/corp/drive/folders/0B5CC9KT63DznUUQtSU9HTHBnb1E) | [Forum](https://discuss.istio.io/c/UX) | [#user-experience](https://istio.slack.com/messages/CFTRP8NTW/) [#config](https://istio.slack.com/messages/C7KSV4AHJ/)| [Notes](https://docs.google.com/document/d/1raZOoeYz3APZdQRlJlFA-zgfmEJu2YLaGa4_3KwbNRs/edit#heading=h.i71zsi7n759v) | [WebEx](https://ibm.webex.com/meet/snible) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfcaVTlYkvV3e09Gak6zhaLj) | User experience across Istio, API and CLI guidelines and support  

_NOTE: Config working group responsibilities have been merged into User Experience._ 

To join Istio's Slack workspace, please fill out [this form](https://docs.google.com/forms/d/e/1FAIpQLSfdsupDfOWBtNVvVvXED6ULxtR4UIsYGCH_cQcRr0VcG1ZqQQ/viewform).

## Working group meetings

Working groups meet regularly. The meetings are open to anyone interested, so please join us. There's a number of ways to track the
meeting schedules:

* You can view the schedule of upcoming working group meetings in this [calendar](https://calendar.google.com/calendar/embed?src=4uhe8fi8sf1e3tvmvh6vrq2dog%40group.calendar.google.com&ctz=America%2FLos_Angeles).

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
<img width="30px" src="https://avatars3.githubusercontent.com/u/8484260?s=400&v=4">  | [Brian Avery](https://github.com/brian-avery)      | Red Hat    | Test and Release
&nbsp;                                                                               | [Costin Manolache](https://github.com/costinm)     | Google     | Networking - Control Plane
<img width="30px" src="https://avatars0.githubusercontent.com/u/2752495?s=400&v=4">  | [Frank Budinsky](https://github.com/frankbu)       | IBM        | Docs
<img width="30px" src="https://avatars0.githubusercontent.com/u/755849?s=400&v=4">   | [Steven Dake](https://github.com/sdake)            | IBM        | Environments
<img width="30px" src="https://avatars0.githubusercontent.com/u/13684010?s=400&v=4"> | [Surya V Duggirala](https://github.com/suryadu)    | IBM        | Performance and Scalability
<img width="30px" src="https://avatars1.githubusercontent.com/u/623453?s=400&v=4">   | [John Howard](https://github.com/howardjohn)       | Google     | Environments
&nbsp;                                                                               | [Mandar Jog](https://github.com/mandarjog)         | Google     | Policies and Telemetry, Performance and Scalability
<img width="30px" src="https://avatars0.githubusercontent.com/u/102881?s=400&v=4">   | [Oliver Liu](https://github.com/myidpt)            | Google     | Security
&nbsp;                                                                               | [Martin Ostrowski](https://github.com/ostromart)   | Google     | Environments
<img width="30px" src="https://avatars2.githubusercontent.com/u/4336228?s=400&v=4">  | [Francois Pesce](https://github.com/fpesce)        | Google     | Test and Release
<img width="30px" src="https://avatars2.githubusercontent.com/u/21148125?s=400&v=4"> | [Douglas Reid](https://github.com/douglas-reid)    | Google     | Policies and Telemetry
&nbsp;                                                                               | [Piotr Sikora](https://github.com/piotrsikora)     | Google     | Networking - Data Plane
<img width="30px" src="https://avatars3.githubusercontent.com/u/3237651?s=400&v=4">  | [Ed Snible](https://github.com/esnible)            | IBM        | User Experience
<img width="30px" src="https://avatars1.githubusercontent.com/u/10537847?s=400&v=4"> | [Eric Van Norman](https://github.com/ericvn)       | IBM        | Test and Release
&nbsp;                                                                               | [Limin Wang](https://github.com/liminw)            | Google     | Security
<img width="30px" src="https://avatars3.githubusercontent.com/u/116378?s=400&v=4">   | [Jason Young](https://github.com/ayj)              | Google     | User Experience
<img width="30px" src="https://avatars3.githubusercontent.com/u/12534779?s=400&v=4"> | [Neeraj Poddar](https://github.com/nrjpoddar)      | Aspen Mesh | Networking - Control Plane
&nbsp;                                                                               | [Lizan Zhou](https://github.com/lizan)             | Tetrate    | Networking - Data Plane

## Getting in touch

If you'd like to reach out to the working group leads, please drop a note to [wg-leads@discuss.istio.io](mailto:wg-leads@discuss.istio.io).
