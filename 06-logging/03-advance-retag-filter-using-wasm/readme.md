# Preqrequisites
1. Install [TinyGo](https://tinygo.org/getting-started/)
2. Install [Golang <1.18>](https://golang.org/doc/install)

# Docker Enviornment

## Build WASM Filter (optional as build artifact named filter.wasm is already included)

```
tinygo build -wasm-abi=generic -target=wasi -o filter.wasm filter.go
```

## Run Fluent Bit In Docker Container

```
docker run \
  -v $(pwd)/filter.wasm:/fluent-bit/etc/filter.wasm \
  -v $(pwd)/fluent-bit.conf:/fluent-bit/etc/fluent-bit.conf \
  -ti cr.fluentbit.io/fluent/fluent-bit:2.0 \
  /fluent-bit/bin/fluent-bit \
  -c /fluent-bit/etc/fluent-bit.conf
```

## Example Output
[View the output](./example.output)

# Kubernetes Enviornment

## Build WASM Filter (optional as build artifact named filter.wasm is already included)

```
tinygo build -wasm-abi=generic -target=wasi -o filter.wasm filter.go
```

## Create ConfigMap

```
kubectl create configmap wasm-filter --from-file=filter.wasm  --namespace=default
```

## Run Fluent Bit In Kubernetes

```
helm upgrade -i fluent-bit fluent/fluent-bit --values values.yaml
```