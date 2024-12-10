# K8S Depends

<p>
  <a href="https://img.shields.io/github/go-mod/go-version/s0ders/k8s-depends"><img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/s0ders/k8s-depends"></a>
  <a href="https://img.shields.io/github/actions/workflow/status/s0ders/k8s-depends/main.yaml?label=CI"><img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/s0ders/k8s-depends/main.yaml?label=CI"></a>
  <a href="https://goreportcard.com/report/github.com/s0ders/k8s-depends"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/s0ders/k8s-depends"></a>
  <a href="https://github.com/s0ders/k8s-depends/blob/main/LICENSE.md"><img alt="GitHub License" src="https://img.shields.io/github/license/s0ders/k8s-depends?label=License"></a>
</p>

## Motivations 

This Go program aims to give Kubernetes a similar capability that Docker Compose offers : `depends_on`. In Docker 
Compose, this keyword allows to programmatically set the order in which the containers are to start.

Kubernetes does not propose this functionality as it is out of the tool's scope. Hence, this small Go program and its 
Docker image offers a simple solution to make sure that containers inside a Pod will not start until the specified 
Kubernetes service are alive. 

The program accepts the following flags:

| Flag        | Option                                                                        | Default Value |
|-------------|-------------------------------------------------------------------------------|---------------|
| `--timeout` | Set the timeout in seconds for each service after which it is considered down | 60            |
| `--sleep`   | Set the timeout in seconds between two consecutive polls of a given service   | 1             |

The rest of the arguments passed to the program are the services in form `SERVICE_NAME:SERVICE_PORT` you want to wait for.

## Install

This program is available as a Docker image and is hosted on Docker Hub:
```bash
$ docker pull s0ders/k8s-depends
```

## Examples

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
    image: s0ders/depends
    args: ["--timeout=120", "database-svc:5432", "cache-svc:6379"]
  containers:
  - name: web_server
    image: nginx
```

To check the logs of your "depends" `initContainers` you can use the following `kubectl` command:

```bash
kubectl logs <pod_name> -c <init_container_name>
```
