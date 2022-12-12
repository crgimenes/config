# Simple to use

This example shows the simplest way to use config.

Try the following commands:

- See the help using the -h parameter

```
go run main.go -h
```

- Change parameters via command line

```
go run main.go -db=true -mongodb_port=9090
```

- Try to create an environment variable MONGODB_HOST and see how the content of the struct changes

```
export MONGODB_HOST="myhost.com"
go run main.go
```
