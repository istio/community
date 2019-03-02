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

| Group | Forum | Meeting Notes | Doc Folder | Recorded Meetings | Description
|-------|-------|---------------|------------|-------------------|-------------
| API Management | [Forum](https://discuss.istio.io/c/api-management) | [Notes](https://docs.google.com/document/d/1ub9jIKGGQitivYOeT3AgIGDhEZF7qq9jznC6DVhsozs/edit) | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8JRFJXWTZjMXY5MDQ) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfe2819aE1NawcPeWZanM-26) | API Keys, Content Mediation, Content Translation, OpenAPI Ingestion
| Config | [Forum](https://discuss.istio.io/c/config) | [Notes](https://docs.google.com/document/d/1P3p7zOpX66hPoZBi_CiC36JW7JmoaLWqE2sgHvdq5tY/edit?ts=5a0b7200) | [Folder](https://drive.google.com/corp/drive/folders/0B5CC9KT63DznUUQtSU9HTHBnb1E) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffOC_hPQhaIDf8v1TxMboJj) | Config API, config format, config management server/storage, config distribution/rollout, API documentation/style guide/governance
| Docs | [Forum](https://discuss.istio.io/c/contributors/docs) | [Notes](https://docs.google.com/document/d/1RGb0NOp0J9QSIrMrZ6wokl16RbSfKHy-6NyTpOQdPa8/edit#heading=h.xjlp01fjb1kv) | [Folder](https://drive.google.com/corp/drive/u/0/folders/1C6X-UyN008fjBrGcWmRGxnliq0Jkpbih) | n/a | User docs, information architecture, istio.io infrastructure
| Environments | [Forum](https://discuss.istio.io/c/environments) | [Notes](https://docs.google.com/document/d/1Ot9AeoiNYnI3fbQrq3w_-cyGxOqS8AD0RChkQfVxyhs/edit) | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8JQWtfaS1MVk1pOHc) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfelYS1XmTO1IaX4crk79tye) | Raw VM support, Hybrid Mesh, Mac/Windows support, Cloud Foundry integration
| Networking | [Forum](https://discuss.istio.io/c/networking) | [Notes](https://docs.google.com/document/d/1xHy2jQ8oiwMponMVY2zJr2eUAmHW_Hi9JK42a7cg5Pc/edit#heading=h.2ju2wl4o5jbc) | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8Jb1QwTXl1WTRqWHM) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffpmIKb5tVthiiQLSXenCK-) | Traffic Management, TCP Support, Additional L7 protocols, Proxy injection
| Performance and Scalability | [Forum](https://discuss.istio.io/c/performance-and-scalability) | [Notes](https://goo.gl/ENFQWb) | [Folder](https://drive.google.com/corp/drive/u/0/folders/1Zpi5TcBPSqGno96WAq5QkcNKi16opaTL) | n/a | Performance and scalability characterization and improvements
| Policies and Telemetry | [Forum](https://discuss.istio.io/c/policies-and-telemetry) | [Notes](https://docs.google.com/document/d/1pn9QdRcoyT_nxOwzklsiYpt7OQraaSDfmtN14XTOrN0/edit) | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8JbWtHOEo1STc1dGM) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdffF-9nyaw01Ni_0GOWBzaF4) | Mixer, Mixer Adapters, Rate Limiting, Tracing, Monitoring, Logging
| Security | [Forum](https://discuss.istio.io/c/security) | [Notes](https://docs.google.com/document/d/1-Z7nSDV8f9psp5chO6H9PAUAinIVapOhQRaUIxbf0dk/edit#heading=h.ppfvkxi2bzdh) | [Folder](https://drive.google.com/corp/drive/u/0/folders/0BzW5bSyKst8Jb2hhRWQ2eTJYVzQ) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfd8ZbUNlZe-RYQKYTxTNJTR) | Service-to-service Auth, Identity/CA/SecretStore plugins, Identity Federation, End User Auth, Authority Delegation, Auditing
| Test and Release | [Forum](https://discuss.istio.io/c/test-and-release) | [Notes](https://docs.google.com/document/d/18QgpvBH9N8Io5xU-0piysyOYif65U03m8GabpvHb4IQ/edit) | [Folder](https://drive.google.com/corp/drive/u/0/folders/0B_ObhNOUiqZ0UWo1ODBVTkRkSzg) | [YouTube](https://www.youtube.com/playlist?list=PL7wB27eZmdfckNxJRS1ac9rqsx5b16NU_) | Build, test, release
| User Experience | [Forum](https://discuss.istio.io/c/UX) | [Notes](https://docs.google.com/document/d/1raZOoeYz3APZdQRlJlFA-zgfmEJu2YLaGa4_3KwbNRs/edit#heading=h.i71zsi7n759v) | [Folder](https://drive.google.com/drive/u/0/folders/1r3MDokrMU1R-jxrZh1rxBV3dYChAprJg) | n/a | User experience across Istio

## Working group meetings

Working groups meet regularly. You can view the schedule of upcoming working group meetings in this [calendar](https://calendar.google.com/calendar/embed?src=4uhe8fi8sf1e3tvmvh6vrq2dog%40group.calendar.google.com&ctz=America%2FLos_Angeles).
You can download [an ICS file](https://calendar.google.com/calendar/ical/4uhe8fi8sf1e3tvmvh6vrq2dog%40group.calendar.google.com/public/basic.ics)
to import the working group meetings into your favorite iCal client.
And finally, if you have a Google account, you can link the [working group event calendar](https://calendar.google.com/calendar?cid=NHVoZThmaThzZjFlM3R2bXZoNnZycTJkb2dAZ3JvdXAuY2FsZW5kYXIuZ29vZ2xlLmNvbQ)
directly to your account.

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
<img width="30px" src="https://avatars3.githubusercontent.com/u/22780957?s=400&v=4"> | [Martin Taillefer](https://github.com/geeknoid) | Google  | API Management, Docs
&nbsp; | [Limin Wang](https://github.com/liminw) | Google | Security

## Getting in touch

If you'd like to reach out to the working group leads, please drop a note to [wg-leads@discuss.istio.io](mailto:wg-leads@discuss.istio.io).
