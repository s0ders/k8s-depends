## Depends - K8S utility


### Motivations 

This Go program aims to give Kubernetes a similar capability that Docker Compose offers : `depends_on`. In Docker Compose, this keyword allows to programmaticaly set the order 
in which the containers are to start.

Kubernetes does not propose this functionality as it is out of the tool's scope. Hence, this small Go program and its Docker image offers a simple solution to make sure 
that containers inside of a Pod will not start until the specified Kubernetes service are alive. 

The image accepts the following flags:

Flag | Option | Default Value
--- | --- | ---
| `--timeout` | Set the timeout in seconds for each service after which it is considered down | 60 |

The rest of the arguments passed to the program are the services in form `SERVICE_NAME:SERVICE_PORT` you want to wait for.

### Examples

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: hello_world
spec:
  # Image is used as an initContainers to prevent containers from starting
  # before the given services are up and running.
  initContainers:
  - name: depends
    image: soders/depends
    args: ["--timeout=120", "database-svc:5432", "cache-svc:6379"]
  containers:
  - name: web_server
    image: nginx
```

To check the logs of your "depends" `initContainers` you can use the following `kubectl` command:

```bash
kubectl logs <pod_name> -c <init_container_name>
```
