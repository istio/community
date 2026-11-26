Istio Product Security Working Group (PSWG) Membership & Release Management
======

# Purpose
The Istio Product Security Working Group (PSWG) exists to:
- Receive and triage security vulnerability reports and CVEs impacting Istio.
- Coordinate private remediation and disclosure.
- Manage CVE assignment and security releases.
- Provide a consistent, timely, and accountable security response process.

# Membership

## Membership Definition
Membership criteria for the PSWG: https://github.com/istio/community/blob/master/EARLY-DISCLOSURE.md#membership-criteria

Each member organization **must name at least one primary representative**.

## Membership Requirements (Minimum)
PSWG member organizations must meet all requirements below:
### Meeting Attendance
  - A representative from the organization must be present at a majority of PSWG meetings within a rolling 90-day window.
  - Attendance is tracked via the meeting notes.
  - “Present” means the representative attended for a meaningful portion of the meeting and is responsive for action items.

### Active Participation in Security Work
Member organizations must actively participate in discussions and work related to:
  - Vulnerability reports and triage decisions
  - Severity and impact assessment
  - CVE planning and disclosure coordination
  - Patch/release timing and communication

### Release Manager Participation
  - See [Roles](#roles-security-release-managers)

### Operational Readiness
  Member organizations must ensure their representatives:
  - Can participate in private security channels as required.
  - Can review private patches and test candidate fixes within reasonable timeframes.
  - Can maintain confidentiality until coordinated disclosure.

## Roles: Security Release Managers
PSWG uses two operational roles for releases:
- Primary Security Release Manager (Primary SRM)
- Secondary Security Release Manager (Secondary SRM)

Roles are assigned based on a rotation schedule. See [Rotation](#rotation) for more information.

### Primary Security Release Manager (Primary SRM)
#### Responsibilities
The Primary SRM is accountable for driving the security response lifecycle. Responsibilities include:
- Intake & Response
  - Monitor and respond to incoming vulnerability report emails.
  - Acknowledge receipt and initiate triage within the defined [SLA](#sla-responding-to-vulnerability-reports)
- Tracking & Coordination
  - Create and maintain tracking artifacts (private GitHub issue(s)).
  - Ensure clear ownership, timelines, and status updates.
- Private Remediation
  - Coordinate development of private fixes (including PRs/patches) and ensure appropriate review.
  - Coordinate backports across supported release branches (as applicable) while coordinating with the respective Istio Release Managers.
  - Ensure tests and validations are executed (repro, integration converage, long running tests).
- CVE Handling
  - Coordinate Istio CVE creation (as applicable to Istio’s CVE process).
  - Ensure descriptions, affected versions, fixed versions, and references are accurate.
- Release Leadership
  - Determine release plan and timeline
  - Coordinate release notes and advisories
  - Coordinate embargo, disclosure, and publishing steps
  - Run the release checklist and ensure completion of required steps.

#### Authority
The Primary SRM may:
- Call emergency meetings.
- Escalate to PSWG leads / maintainers when timelines or risk requires.
- Propose fast-path decisions during active exploitation risk.

### Secondary Security Release Manager (Secondary SRM)
#### Responsibilities
The Secondary SRM is the Primary’s operational partner and continuity plan.
- Backup & Continuity
  - Be ready to assume Primary SRM duties if the Primary is unavailable.
- Active Support
  - Participate in triage decisions and investigation.
  - Review public and private fixes and advisories.
  - Help coordinate testing and validation.
  - Own assigned workstreams (e.g., branch backport coordination, advisory drafting, verification).
- Operational Load Sharing
  - Ensure the response does not depend on a single person.
  - Maintain awareness of current vulnerabilities tracked in GH issues.

## Rotation
### Goals
Rotation exists to:
- Build redundancy and shared context
- Ensure consistent operational coverage

### Model
PSWG maintains a rotating schedule for SRMs.
Each rotation assigns:
- 1 Primary SRM
- 1 Secondary SRM

Rotation length: 3 months (matching Istio minor release cadance).

### Scheduling Rules
Rotation schedule should be published at least one quarter in advance when possible.
Organizations may request swaps with reasonable notice; swaps should be recorded in meeting notes.
If a scheduled SRM is unavailable, the Secondary SRM becomes Acting Primary, and PSWG will designate a new Secondary as soon as possible.

### Coverage Expectations
Both Primary and Secondary SRMs should ensure coverage during:
- business hours for their region
- planned PTO (handoff required)
- high-severity incident response windows

## SLA: Responding to Vulnerability Reports
This SLA defines minimum expectations for responsiveness.

### Acknowledgement SLA
Within 1 business day of report receipt: 
- Acknowledge the report and confirm it is being reviewed
- Create a github issue in istio-private/pswg tracking the vulnerability and any context

### Triage SLA
Within 3 business days: initial triage decision, including:
- Whether the report is in-scope
- Initial severity estimate (preliminary)
- Request for additional information if needed
- Update to reporter if more investigation is required

### Escalated SLA (High Risk)
For reports indicating any of the following, escalate issue in PSWG slack and engage PSWG leads:
- Data plane crash
- Exposure of private information
- Bypass of security policy reaching application

### Status updates in the tracking issue
Provide:
- weekly updates for standard vulnerabilities
- updates every 48 hours for critical/high risk or actively exploited scenarios

## Membership Non-Compliance & Removal

### What Counts as Non-Compliance
A member organization may be considered non-compliant if it fails to meet one or more of:
- **Attendance**: Missing the majority of meetings over a rolling 90-day period. This requirement is flexible for timezone conflicts.
- **Participation**: Repeated lack of meaningful participation in vulnerability/CVE discussions or assigned actions.
- **Release Manager Commitment**:
  - Refusing to serve as Secondary SRM when scheduled without arranging a replacement.
  - Repeatedly failing to provide backup coverage when the Primary SRM is unavailable.
- **Operational Issues**:
  - Consistently missing agreed timelines for reviews/testing without proactive communication.
  - Violating confidentiality/embargo expectations (may trigger immediate action).

### Removal Process
PSWG uses a graduated process to be fair and transparent:
#### Step 1: Notice of Non-compliance

PSWG lead(s) notify the member org representatives that they are not meeting requirements.
The notice includes:
- Which requirement(s) are not met
- Evidence (meeting notes links, missed rotations, etc.)
- A proposed remediation plan and timeframe

#### Step 2: Final Notice of Non-compliance
If issues persist 30 days after initial notice and the remediation plan is not being followed a final notice will be issued and recorded in the PSWG notes.

The warning includes:
- Clear expectations
- Timeline for decision to remove from PSWG (30 days from final notice)
- Specific measurable criteria (e.g., “attend 3 of next 4 meetings,” “fulfill next Secondary SRM rotation,” etc.)

#### Step 3: Vote / Decision to Remove
If the member org does not remediate within the window:
- PSWG leads and Istio's TOC vote to remove member organization. Removal requires a majority vote. Eligible voters from the non-compliant member organization must abstain from voting.
- The outcome is recorded in meeting notes.

#### Step 4: Removal & Access Updates
Upon removal:
- The organization’s PSWG membership is revoked.
- Access to private channels, repos, and documents is removed per security policy.
- Removal is documented in PSWG notes.

### Immediate Removal (Exceptional)
PSWG leads may recommend immediate removal for severe breaches, such as:
- Confidentiality or embargo violations (including non-Istio embargos where Istio's PSWG is privileged to early disclosures)
- Malicious behavior
- Repeated actions that materially endanger users or the project
- Breach of CNCF code of conduct

### Reconsideration for Membership
Removed organizations may apply for reconsideration.

#### Eligibility Window
Reconsideration may be requested no earlier than 60 days after removal (unless PSWG explicitly allows sooner).

#### Reconsideration Requirements
The organization must provide a written statement describing:
- What caused the non-compliance
- What has changed
- How they will meet requirements going forward
- Identify at least 1 primary representative and 1 backup representative
- Commit to participating in SRM rotations.

#### Decision Process
PSWG reviews the request during a PSWG meeting.
Decision is recorded in meeting notes.

## Operational Notes

### Attendance Tracking
Attendance is tracked via meeting notes.
Each meeting notes doc should list attendees and organizations.

### Expected Conduct
Members must follow project and CNCF code of conduct and confidentiality expectations.
PSWG discussions may involve sensitive information; members must not disclose until authorized.
