# Testing kubernetes volumes

- A `hostPath` PersistentVolume is being created and then accessed through a PersistentVolumeClaim
- Because this is a kind cluster, the `hostPath` should be defined in the docker container holding the kind cluster
  because if it's only defined in the local fs, kubernetes will not pick it up as kubernetes is being executed on the
  kind cluster which is a docker container.
- Make sure to create the `pv` and `pvc` before creating the pod.
