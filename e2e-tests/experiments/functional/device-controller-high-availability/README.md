## About this experiment

This functional experiment scale up the device-controller replicas to use it in high availability mode and then verify the device-localpv behaviour when one of the replicas go down. This experiment checks the initial number of replicas of device-controller and scale it by one if a free node is present which should be able to schedule the pods. Default value for device-controller statefulset replica is one.

## Supported platforms:

K8s : 1.18+

OS : Ubuntu

## Entry-Criteria

- k8s cluster should be in healthy state including all the nodes in ready state.
- device-controller and csi node-agent daemonset pods should be in running state.

## Exit-Criteria

- device-controller statefulset should be scaled up by one replica.
- All the replias should be in running state.
- device-localpv volumes should be healthy and data after scaling up controller should not be impacted.
- This experiment makes one of the device-controller statefulset replica to go down, as a result active/master replica of device-controller prior to the experiment will be changed to some other remaining replica after the experiment completes. This happens because of the lease mechanism, which is being used to decide which replica will be serving as master. At a time only one replica will be master and other replica will follow the anti-affinity rules so that these replica pods will be present on different nodes only.
- Volumes provisioning / deprovisioning should not be impacted if any one replica goes down.

## How to run

- This experiment accepts the parameters in form of kubernetes job environmental variables.
- For running this experiment of deploying device-localpv provisioner, clone openens/device-localpv[https://github.com/openebs/device-localpv] repo and then first apply rbac and crds for e2e-framework.
```
kubectl apply -f device-localpv/e2e-tests/hack/rbac.yaml
kubectl apply -f device-localpv/e2e-tests/hack/crds.yaml
```
then update the needed test specific values in run_e2e_test.yml file and create the kubernetes job.
```
kubectl create -f run_e2e_test.yml
```
All the env variables description is provided with the comments in the same file.