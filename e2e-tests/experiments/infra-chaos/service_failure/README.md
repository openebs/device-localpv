## About this experiment

This experiment validates the stability and fault-tolerance of application pod consuming device-localpv storage. It is a infra-chaos kind of e2e-test where docker/kubelet services are stopped on the node where application pod is scheduled and wait for application pod eviction time. After that wait-time, restart the docker/kubelet services on the node and tests the recovery workflow of the application pod.

## Supported platforms:

K8S : 1.18+

OS : Ubuntu

## Entry-Criteria

- One application should be deployed consuming device-localpv storage.
- Application services are accessible & pods are healthy.
- Application writes are successful.
- device-controller and csi node-agent daemonset pods should be in running state.

## Exit-Criteria

- Application services are accessible & pods are healthy.
- Data written prior to infra-chaos is successfully retrieved/read.
- Data consistency is maintained as per integrity check utils.

## Steps performed

- Get the application pod name and its pvc name.
- Dump some dummy data into the application mount point to check data consistency after chaos injection.
- Get the node details where application pod is running.
- Call the service_failure chaos utils where we first get the node ip and stop the docker/kubelet service (based on SVC_TYPE env).
- After stopping the service, check the node_status till its status becomes NotReady.
- Wait for the pod eviction time (by default it is 300 seconds). Though it is a localpv storage that's why pod will stuck to its node only where its storage is present. There will be two application pod one will be in Terminating state and other one will be in Pending state.
- Now restart the services on node and check for the application pod status. Older pod should be terminate successfully and newer pod should be in Running state.
- Validate the data consistency by checking the md5sum of test data.

## How to run

- This experiment accepts the parameters in form of kubernetes job environmental variables.
- For running this infra-chaos experiment of node_failure, clone openens/device-localpv[https://github.com/openebs/device-localpv] repo and then first apply rbac and crds for e2e-framework.
```
kubectl apply -f device-localpv/e2e-tests/hack/rbac.yaml
kubectl apply -f device-localpv/e2e-tests/hack/crds.yaml
```
then update the needed test specific values in run_e2e_test.yml file and create the kubernetes job.
```
kubectl create -f run_e2e_test.yml
```
All the env variables description is provided with the comments in the same file.

After creating kubernetes job, when the job’s pod is instantiated, we can see the logs of that pod which is executing the test-case.

```
kubectl get pods -n e2e
kubectl logs -f <service-failure-chaos-xxxxx-xxxxx> -n e2e
```
To get the test-case result, get the corresponding e2e custom-resource `e2eresult` (short name: e2er ) and check its phase (Running or Completed) and result (Pass or Fail).

```
kubectl get e2er
kubectl get e2er {{svc_type}}-service-failure -n e2e --no-headers -o custom-columns=:.spec.testStatus.phase
kubectl get e2er {{svc_type}}-service-failure -n e2e --no-headers -o custom-columns=:.spec.testStatus.result
```