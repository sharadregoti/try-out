# Building Docker Images for all services
`make build`

# Running all Services
`make apply`

# Running Jaeger in Kubernetes
`make apply-jaeger`

# Running OpenTelemetry Collector in Kubernetes
`make apply-otel-collector`

# Accessing Jaeger UI
`kubectl port-forward svc/jaeger-query 16686:16686`

# Making a request to the API
- Exec into any of the pod container
- Install curl
- `curl account-service:5000/getBalance/1`
You should get this response `{"message":"Balance request received for account ID: 1!"}`