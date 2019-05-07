#!/bin/bash

set -eu

# ROOT_PACKAGE :: the package (relative to $GOPATH/src) that is the target for code generation
ROOT_PACKAGE="github.com/deislabs/smi-sdk-go"
# CUSTOM_RESOURCE_NAME :: the name of the custom resource that we're generating client code for
CUSTOM_RESOURCE_NAME="trafficsplit"
# CUSTOM_RESOURCE_VERSION :: the version of the resource
CUSTOM_RESOURCE_VERSION="v1beta1"

bindir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
rootdir="$( cd $bindir/.. && pwd )"

# The generate-groups.sh script cannot handle group names with dashes, so we use
# smispec.io as the group name, then replace it with smi-spec.io after code
# generation.
find "$rootdir/gen/" -type f -exec sed -i 's/smi-spec.io/smispec.io/g' {} +

# run the code-generator entrypoint script
genver=$( awk '/k8s.io\/code-generator/ { print $2 }' "$rootdir/go.mod" )
gendir=${GOPATH}/pkg/mod/k8s.io/code-generator@${genver}
chmod +x ${gendir}/generate-groups.sh
${gendir}/generate-groups.sh all "$ROOT_PACKAGE/gen/client/$CUSTOM_RESOURCE_NAME" "$ROOT_PACKAGE/gen/apis/" "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

# The generate-groups.sh script cannot handle group names with dashes, so we use
# smispec.io as the group name, then replace it with smi-spec.io after code
# generation.
find "$rootdir/gen/" -type f -exec sed -i 's/smispec.io/smi-spec.io/g' {} +
