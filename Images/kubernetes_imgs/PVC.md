# Persistant Volume Claim

## Pvc verifications

First ensure of the good running of the cluster:

![pvc](get_deployment.png)

Verify the creation of the pvc:

![pvc](get_pvc.png)

Get the redis pod name:

![pvc](get_pods.png)

Connect to the redis cli and set a key value:

![pvc](set_testkey.png)

Then delete the pod, kubernetes will recreate it automatically and verify that the key is still present:

![pvc](result.png)

