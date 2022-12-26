# go-svc

Template for building services in Go.

## Features

- [x] Minimum dependencies service template
- [x] Structure logging with [zerolog](https://github.com/rs/zerolog)
- [x] [Prometheus](https://prometheus.io/) [metrics](./pkg/httpe/metrics/)
- [x] [Admin endpoints](./pkg/httpe/admin/routes.go) on separate port
- [x] HTTP [response helpers](./pkg/httpe/response.go)
- [x] HTTP chainable [middlewares](./pkg/httpe/middleware.go)
- [x] Dockerfile and Kubernetes deployments
- [x] Local development with [Tilt](https://tilt.dev/) and [Kind](https://kind.sigs.k8s.io/)
- [ ] Github Actions
- [ ] Grafana dashboard

## Sample requests

Requests for:

1. [Sample service](service.http)
1. [Admin service](admin.http)

Can be run directly from VS Code using [Rest Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).
