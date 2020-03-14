# Contributing Guidelines

:heavy_check_mark::tada: Let's make code better together - Thanks for taking the time to contribute! :tada::heavy_check_mark:

The following is a set of guidelines for contributing to *0x0I Prometheus Exporters*, which are hosted under the [0x0I](https://github.com/0x0I?tab=repositories) developer account on GitHub. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

For general information and guidelines for contributing to these Prometheus exporters and the greater Prometheus community, see the [exporter maintainability page](https://prometheus.io/docs/instrumenting/writing_exporters/#maintainability-and-purity).

**Table of Contents**
  - [Pull Requests](#pull-requests)
      - [CI Pipeline](#ci-pipeline)
  - [Issues](#issues)
      - [Issue Types](#issue-types)
  - [Workflow and backlog](#workflow-and-backlog)
  - [Code of Conduct](#code-of-conduct)

## Pull Requests

All [PRs](https://github.com/0x0I/aws_ec2_exporter/pulls) are welcome! :+1: The following guidelines and CI pipeline are provided for validating container functionality and avoiding regressions for each submitted request:

**Note:** to build and test this container/image, a working copy of either the `docker` or `podman` container runtimes/engines and the `dgoss` testing framework are necessary.

#### CI Pipeline

| Step | Description | Dependencies | Command |
| --- | --- | --- | --- |
| :construction: `Image build` | Leverages [Ansible](https://docs.ansible.com/) for defining variable build arguments and building playbooks implementing the end-to-end Image release process. | [ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html) | `ansible-playbook build/build-playbook.yml` |
| :zap: `golint` | Validates Go code quality adheres to Golang standards and practices as [defined](https://github.com/golang/go/wiki/CodeReviewComments). | [Golint](https://github.com/golang/lint) (Go package) | `golint .` |
| :wrench: `integration testing` | Utilizing the Goss testing framework along with its associated `dgoss` extra package, integration testing of this container file is organized according to the various provisioning phases and should be executed prior to PR submission to validate new modifications and identify/prevent regressions. | [dgoss](https://github.com/aelsabbahy/goss/tree/master/extras/dgoss) | `test/integration_tests.sh` |
| :traffic_light: `Continuous Integration (CI)` | Automatic E2E testing of this role is accomplished leveraging the [Travis-CI](https://travis-ci.com/0x0I/container-file-kafka) test infrastructure platform and is executed on each pull request. Requests should not be merged unless all tests pass or the community approves otherwise. | *N/A* | *see* [.travis.yml](https://github.com/0x0I/container-file-kafka/blob/master/.travis.yml) for additional details |

## Issues

New GitHub issues can be [opened](https://github.com/0x0I/aws_ec2_exporter/issues/new) and [tracked](https://github.com/0x0I/aws_ec2_exporter/issues) in a similar fashion as with most Github repositories by making use of the standard Github issue management facilities.

Reference the following issue reporting guide for more details:

#### Issue Types

| Issue Type | Description |
| --- | --- |
| :arrow_up: `:enhancement:` | Feature requests. |
| :bug: `:bug:` | Confirmed bugs or reports that are very likely to be bugs. |
| :question: `:question:` | Questions more than bug reports or feature requests (e.g. how do I do X). |
| :eyeglasses: :heartpulse:`:feedback:` | General feedback more than bug reports or feature requests. |

## Workflow and backlog

Reference this repository's [wiki](https://github.com/0x0I/aws_ec2_exporter/wiki) to visualize the project roadmap, workflow and backlog to stay up to speed with development  plans and work in progress.

## Code of Conduct

See the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md).
