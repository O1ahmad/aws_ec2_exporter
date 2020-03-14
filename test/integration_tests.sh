#!/bin/bash

set -euo pipefail

# Print all commands executed if DEBUG mode enabled
[ -n "${DEBUG:-""}" ] && set -x

# [Test-Setup]
docker build --file build/Containerfile --tag aws-ec2-exporter:testing .

# [Test-Run+Validate]
GOSS_FILES_PATH=test dgoss run aws-ec2-exporter:testing
