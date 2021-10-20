# Thanos

### Configuration

#### Install dependences
```
make install
```

#### Build
```
make build-with-docker
```

#### Configure Server
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