# Onboarding Technologies to Istio

On occasion, we encounter situations where software developed outside of the scope of the Istio project has
sufficient alignment with the main Istio mission that it should become an integral part of Istio proper.

This document describes a set of requirements that must be satisfied before externally developed software can
be folded into the Istio project.

## Forms of Integration

Externally developed software can be integrated into the Istio project in a number of ways, falling into three broad categories:

* **Augment**. Represents changes and additions to the main code base to introduce small new features and capabilities. These changes would typically
be introduced as normal PRs into the existing Istio repository, possibly within a feature branch.

* **Extend**. Represents bigger changes to the platform, including the addition of new components. An example of this could be introducing a standard
Open Service Broker implementation for Istio. This type of change would typically be added in a subdirectory of the main Istio repo, and introduced via
a feature branch. Technologies of this type are subject to Istio's normal release cadence.

* **Support**. Represents functionality that is peripheral to the main Istio product. These tools and other binaries can be added as
standalone repositories within the Istio GitHub organization. Technologies of this type may be have a different release cadence than
the Istio's normal release cadence.

## Prerequisites

* Based on technical merits and alignment, the TOC must agree to the integration, must establish the proper form of integration, and must determine the overall protocol to follow (feature branch vs. master branch, timeline, testing requirements, doc requirements, etc).

* Contributors of the new technology must be owners of the copyright of all subject material.

* The transitive dependencies for **Augment** and **Extend** integrations must all use one of the supported Istio dependency license types
(Apache, MIT, and BSD). **Support** integrations can have dependencies with additional license types (TBD: enumerate which ones)

* Copyright for all contributed material must be explicitly assigned in the source code to "Istio Authors" and must be licensed under the Apache license.

* Contributors must agree to whatever CLA the Istio project is currently subject to.

* Contributors must agree to supply the initial documentation for the technology to <https://istio.io> or to Istio’s developer wiki on GitHub, as appropriate.

* Contributors must agree to supply test suites that ensure the new technology functions properly and meets Istio’s testing and quality standards.

* Contributors must agree to update the technology's API surface to comply with Istio’s overall design guidelines.
