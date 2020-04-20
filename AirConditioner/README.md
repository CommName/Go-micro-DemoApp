# AirConditioner Service

This is the AirConditioner service

Generated with

```
micro new AirConditioner --namespace=github.com/CommName --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: github.com/CommName.srv.AirConditioner
- Type: srv
- Alias: AirConditioner

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./AirConditioner-srv
```

Build a docker image
```
make docker
```