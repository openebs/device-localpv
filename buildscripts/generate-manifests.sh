#!/bin/bash

# Copyright 2019 The Kubernetes Authors.
# Copyright 2020 The OpenEBS Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

## find or download controller-gen
CONTROLLER_GEN=$(which controller-gen)

if [ "$CONTROLLER_GEN" = "" ]
then
  echo "ERROR: failed to get controller-gen, Please run make bootstrap to install it";
  exit 1;
fi

$CONTROLLER_GEN crd:trivialVersions=false,preserveUnknownFields=false paths=./pkg/apis/... output:crd:artifacts:config=deploy/yamls

## create the the crd yamls

echo '

##############################################
###########                       ############
###########   DeviceVolume CRD    ############
###########                       ############
##############################################

# DeviceVolume CRD is autogenerated via `make manifests` command.
# Do the modification in the code and run the `make manifests` command
# to generate the CRD definition' > deploy/yamls/devicevolume-crd.yaml

cat deploy/yamls/local.openebs.io_devicevolumes.yaml >> deploy/yamls/devicevolume-crd.yaml
rm deploy/yamls/local.openebs.io_devicevolumes.yaml

echo '

##############################################
###########                       ############
###########   DeviceNode CRD      ############
###########                       ############
##############################################

# DeviceVolume CRD is autogenerated via `make manifests` command.
# Do the modification in the code and run the `make manifests` command
# to generate the CRD definition' > deploy/yamls/devicenode-crd.yaml


cat deploy/yamls/local.openebs.io_devicenodes.yaml >> deploy/yamls/devicenode-crd.yaml
rm deploy/yamls/local.openebs.io_devicenodes.yaml

## create the operator file using all the yamls

echo '# This manifest is autogenerated via `make manifests` command
# Do the modification to the device-driver.yaml in directory deploy/yamls/
# and then run `make manifests` command

# This manifest deploys the OpenEBS Device control plane components,
# with associated CRs & RBAC rules.
' > deploy/device-operator.yaml

# Add namespace creation to the Operator yaml
cat deploy/yamls/namespace.yaml >> deploy/device-operator.yaml

# Add DeviceVolume v1alpha1 CRDs to the Operator yaml
cat deploy/yamls/devicevolume-crd.yaml >> deploy/device-operator.yaml

# Add DeviceNode v1alpha1 CRDs to the Operator yaml
cat deploy/yamls/devicenode-crd.yaml >> deploy/device-operator.yaml

# Add the driver deployment to the Operator yaml
cat deploy/yamls/device-driver.yaml >> deploy/device-operator.yaml

# To use your own boilerplate text use:
#   --go-header-file ${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt
