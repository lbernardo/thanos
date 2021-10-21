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
thanos config:server --ip IP_SERVER --config $HOME/.thanos.json
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

#### Create Service

```json
{
    "name": "webapp",
    "environments": {
        "prod": {
            "image": "lbernardo/webapp",
            "replicas": 1,
            "service": {
                "host": "webapp.lubipay.com",
                "port": 80
            }
        }
    }
}
```

| Name | Description | Eg           |
|------|-------------|--------------|
| name | App name    | my-first-app |
| environments | `String: Object{}`  environments use | prod, staging, xxx, sb, etc |
| environments.X.image | Image for your deployment | nginx |
| environments.X.replicas | Number replicas | 2 |
| environments.X.service | Service configuration | |
| environments.X.service.host | Host configuration | web.localhost.com |
| environments.X.service.port | Port host service | 80 |

#### Apply with tag and environment
```
thanos apply example.json --tag v1.0 --environment prod
```