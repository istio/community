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

* Development
    * Requires Approved RFC in order to merge.
* Config
    * Requires explicit user action to enable (e.g. a config field, config resource, or installation action).

* Docs
    * Reference docs are published on istio.io.
    * Basic feature docs are published on istio.io describing what the feature does, how to use it, and any caveats.
    * A reference to the design doc / issue is published on istio.io.

* Tests
    * Test coverage is >= 70%.
    * Integration tests cover core use cases with the feature enabled.
    * When disabled the feature does not affect system stability or performance.

* Performance
    * Performance requirements assessed as part of design.

* API
    * (Optional) Initial API review.

## Beta

An Alpha feature can be officially labeled as Beta once it meets the following additional requirements:

* Development
    * Requires Approved Design Doc in order to merge.
* Config
    * Can be enabled by default without requiring explicit user action.

* Docs
    * Performance expectations of the feature are documented, may have caveats.
    * Documentation on istio.io includes samples/tutorials
    * Documentation on istio.io includes appropriate glossary entries

* Tests
    * Test coverage is >= 80%
    * Integration tests cover feature edge cases
    * End-to-end tests cover samples/tutorials
    * Fixed issues have tests to prevent regressions

* Performance
    * Feature has baseline performance tests.

* API
    * API has had a thorough API review and is thought to be complete.

* CLI
    * Necessary CLI commands have been implemented and are complete.

* Bugs
    * Feature has no outstanding P0 bugs.

## Stable

A Beta feature can be officially labeled as Stable once it meets the following additional requirements:

* Tests
    * Test coverage is >= 90%
    * Automated tests are in place to prevent performance regressions.

* Performance
    * Latency, throughput and scale are quantified and documented.

* Bugs
    * Feature has no outstanding P0 or P1 bugs.
