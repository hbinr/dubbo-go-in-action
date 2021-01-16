# Run
## Start nacos server
If you are using [nacos](https://nacos.io/en-us/docs/quick-start.html) as a service registry, then it must be started first.

Or if you use another registry, like zookeeper, please also start the zookeeper before executing main.go
## Set env
### Step 1
```sh
cd cmd/
```
### Step 2
#### For Linux
```sh
export CONF_PROVIDER_FILE_PATH="../conf/server.yml"
export APP_LOG_CONF_FILE="../conf/log.yml"
```

#### For Windows
```sh
set CONF_PROVIDER_FILE_PATH="../conf/server.yml"
set APP_LOG_CONF_FILE="../conf/log.yml"
```

## excute main.go

```go
go run main.go
```