# Temporal Sample

This repo includes:

1. Nomad Jobspecs for setting up Temporal in Nomad.
2. Go source files from the Temporal Hello World project. For more info, check out Temporal's [Hello World Tutorial](https://docs.temporal.io/docs/go/hello-world-tutorial).

## Instructions

1. Run the worker and starter included in the project.

```bash
go run worker/main.go
go run start/main.go
```

If you have [`nodemon`](https://nodemon.io/) installed, you can automatically reload when you change any files: `nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run worker/main.go`


## Additional Notes

### Testing the gRPC Connection 

To test the gRPC connection, you can install the Temporal CLI, `tctl`:

```bash
brew install tctl
```

For non-Macs, check out [these instructions](https://docs.temporal.io/docs/tctl/how-to-install-tctl).

Then run the command below. This command just lists the namespaces in Temporal.

```bash
tctl --address <grpc_endpoint> --context_timeout <timeout_in_sec> namespace list
```

Where:

* `<grpc_endpoint>` is your gRPC endpoint. For example: `temporal-app.localhost:7233`
* `<timeout_in_sec>` is the context timeout, in seconds. This is handy if your DNS resolution is slow. Seems not be pretty fast in HCE Pre.

## Web UI

You can access the Temporal Web UI based on the value defined for the `traefik.http.routers.temporal-web.rule` in the `service` stanza of `temporal.nomad`. 
### MySQL Database

The DB backend that we're using for Temporal is a MySQL database, which is currently running in Nomad. It is not mounted to a fixed volume, so when the container dies, we lose our storage. This is NOT recommended for prod.