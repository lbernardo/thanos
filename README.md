# Thanos

### Configuration

#### Kubernetes Ingress
First execute after commands:

https://doc.traefik.io/traefik/v1.7/user-guide/kubernetes/

#### Configure server
```
thanos config:server
```

### Configure client
```
thanos config:client --secret %SECRET% --server %IP_SERVER%
```

### Start server in cluster

```
thanos server
```

### Apply config file with client

```
thanos apply example.json
```