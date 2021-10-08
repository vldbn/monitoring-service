# Monitoring service

### Setting up

Monitoring service uses variables. If no variables are set, the default values listed below will be used:

```text
PORT=8000
SECRET_KEY=secretKeyString
DATABASE_URL=mongodb://localhost:27017
USERNAME=username
PASSWORD=password
```

For setting variable use command below:

```text
export PORT=8000
...
```

### Starting service

For running the service locally, perform the settings as above and start service by command:

```text
go run cmd/main/main.go
```

Also, you can run service with docker-compose by command:

```text
docker-compose up --build
```

If services is running in docker, it will use env variables from docker-compose file. By default, they are:

```text
PORT=8000
USERNAME=username
PASSWORD=password
```

### Open API

Open api available on address:

```text
http://localhost:8000/docs
```


##### Time spent
- reading fasthttp documentation: ~1h
- service implementation: ~8h