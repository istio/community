# Feature Lifecycle Checklist

This is the checklist describing the essential requirements for a feature to
be labeled as Alpha, Beta, or Stable. Once a feature meets the Alpha label
+criteria, the feature may be included in the
+[feature status page](https://istio.io/about/feature-stages/) on
+[Istio's Website](https://istio.io). This is a shortened version of the full
[Feature Lifecycle](FEATURE-LIFECYCLE.md) document.

- [Alpha](#alpha)
- [Beta](#beta)
- [Stable](#stable)

## Alpha

A development feature can be officially labeled as Alpha once it meets the following
requirements:

* Design
    * Requires Approved RFC.
* Config
    * Requires explicit user action to enable (e.g. a config field, config resource, or installation action).

* Docs
    * Reference docs are published on the Istio wiki.
    * Basic feature docs are published on istio.io describing what the feature does, how to use it, and any caveats.

* Tests
    * Automated integration tests cover core use cases with the feature enabled.
    * When disabled the feature does not affect system stability or performance.

* Performance
    * Performance impacts have been measured.

* API
    * Initial API review.

## Beta

An Alpha feature can be officially labeled as Beta once it meets the following additional requirements:

* Design
    * Requires Approved Design Doc in order to merge.
* Config
    * Can be enabled by default without requiring explicit user action.

* Docs
    * Documentation on istio.io includes performance expectations, may have caveats.
    * Documentation on istio.io includes samples/tutorials
    * Documentation on istio.io includes appropriate glossary entries

* Tests
    * Integration tests cover feature edge cases
    * End-to-end tests cover samples/tutorials
    * Fixed issues have tests to prevent regressions

* Performance
    * Feature has baseline performance tests.

* API
    * API has had a thorough API review and is thought to be complete.

* CLI
    * Any necessary CLI commands have been implemented and are complete.

* Bugs
    * Feature has no major known issues

## Stable

A Beta feature can be officially labeled as Stable once it meets the following additional requirements:

* Tests
    * Automated tests are in place to prevent regressions.

* Performance
    * Latency, throughput and scalability are quantified and documented on istio.io.

* Bugs
    * Feature has no major known issues.
