#!/bin/sh
set -eu

# Path to script
SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)

. "$SCRIPT_DIR/.env"

printf -- \
'--build-arg GO_VERSION=%s --build-arg BUILD_VERSION=%s' \
"$GO_VERSION" "$BUILD_VERSION"
