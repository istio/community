# Early Disclosure of Security Vulnerabilities

When an Istio security vulnerability is discovered we follow the process outlined in 
[Security Vulnerabilities](https://istio.io/about/security-vulnerabilities/) in order to 
resolve the issue in a prompt and responsible manner.

The Istio project maintains a mailing list for private early disclosure of 
security vulnerabilities. 
The list is used to provide actionable information to close Istio partners.
The list is *NOT* intended for individuals to find out about
security issues.

Istio has upstream dependencies of its own. In certain situations, we may need to make emergency 
Istio releases to address vulnerabilities in these upstream dependencies. Due to embargo policies of
these dependencies, we may not be allowed to share any details about the vulnerabilities. In that
case, Istio's early disclosure will be limited, on a case-by-case basis.
Note that Envoy vulnerabilities are handled slightly differently, see [Membership Criteria](#membership-criteria)

## Embargo policy

The information members receive from the mailing list must not be
made public, shared, nor even hinted at anywhere beyond the need-to-know within
your specific team except with the list's explicit approval. This holds
true until the public disclosure date/time that was agreed upon by the list.
Members of the list and others may not use the information for anything other
than getting the issue fixed for your respective distribution's users.

Before any information from the list is shared with respective members of your
team required to fix said issue, they must agree to the same terms and only
find out information on a need-to-know basis.

In the unfortunate event you share the information beyond what is allowed by
this policy, you _must_ urgently inform the mailing list by replying to the 
initial disclosure email. The email should specify exactly what information
leaked and to whom. A retrospective will take place after the leak so
we can assess how to not make the same mistake in the future.

If you continue to leak information and break the policy outlined here, you
will be removed from the early disclosure list.

## Membership

| Email		| Organization	|
| ------------- |:-------------:|
| tanzu.psirt@broadcom.com | Broadcom |
| istio-security-reports@f5.com | F5 Networks |
| istio-security-vulnerability-notifications@google.com | Google |
| argoprod@us.ibm.com | IBM |
| envoy-security@microsoft.com | Microsoft |
| istio-security@redhat.com | RedHat |
| cve@solo.io | Solo.io |
| istio-security@tetrate.io | Tetrate |

### Membership criteria

To be eligible for the early disclosure mailing list, your
organization must:

1. Have an actively monitored security email alias for our project. This cannot be a personal
email address, it must be a corporate address owned by the organization.
2. Be involved in redistributing Istio or have a product with close coupling to Istio.
3. Have a user base not limited to your own organization.
4. Be a participant and active contributor in the Istio community.
5. Accept the [Embargo Policy](#embargo-policy) that is outlined above.
6. Have a member of your organization be a participant in the Istio Product Security Working Group. This is a private meeting held alternating Wednesdays at 11am Pacific or 3pm Pacific to accommodate a world-wide community. Participants will be expected to actively contribute and fill the roll as security release lead on a quarterly rotational basis.

#### Removal

If your organization stops meeting one or more of these criteria
after joining the list then you will be unsubscribed.

#### Envoy

Because Istio works closely with Envoy, and is often involved in security vulnerabilities, members of the early disclosure list will inherently gain access to details about embargoed Envoy vulnerability details.
While we do not currently require members to be a part of the [Envoy early disclosure list](https://github.com/envoyproxy/envoy/blob/main/SECURITY.md#members), we do expect that members will follow their embargo guidelines.
Failure to do so, or at the request of the Envoy security team, may result in your removal from the early disclosure list.

### Request to join

New membership requests should be emailed to [istio-security-vulnerability-reports@googlegroups.com](mailto:istio-security-vulnerability-reports@googlegroups.com)

In the body of your email please specify how you qualify and fulfill each
criterion listed in [Membership criteria](#membership-criteria).

Here is an example:

```
To: istio-security-vulnerability-reports@googlegroups.com
Subject: Seven-Corp Membership to Early Disclosure Mailing List

Below are each criterion and why I think we, Seven-Corp, qualify.

> 1. Have an actively monitored security alias email for our project.

Yes, please subscribe istio-security-team@example.com to the early disclosure
list.

> 2. Be involved in redistributing Istio or have a product with close coupling to Istio.

We distribute the "Seven" distribution of Istio [link]. We have been doing
this since 2017.

> 3. Have a user base not limited to your own organization.

Our user base spans of the extensive "Seven" community. We have a slack and
GitHub repos and mailing lists where the community hangs out. [links]

> 4. Be a participant and active contributor in the Istio community.

Our members, Acidburn, Cereal, and ZeroCool are outstanding members and are well
known throughout the Istio community. Especially for their contributions
in hacking the Gibson.

> 5. Accept the Embargo Policy that is outlined above.

We accept.

> 6. Have a member of your organization be a participant in the Istio Product Security Working Group.

John/Jane Doe (Email: person@myorganization.com, Istio Slack username: person) will represent our company in the Istio Product Security Working Group.
```

Final approval is required from the Istio TOC pending a formal request from the PSWG.
