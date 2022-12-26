# -*- mode: Python -*-
load('ext://namespace', 'namespace_create', 'namespace_inject')
namespace_create('monitoring')
load('ext://helm_resource', 'helm_resource', 'helm_repo')
helm_repo('prometheus-community', 'https://prometheus-community.github.io/helm-charts')
helm_resource(name='prometheus-stack', chart='prometheus-community/kube-prometheus-stack', namespace='monitoring', flags=['-f', 'deploy/prometheus/values.yaml'])

docker_build('bwilczynski/go-svc', '.', dockerfile='Dockerfile')
k8s_yaml(kustomize('./deploy/go-svc'))
k8s_resource('go-svc', port_forwards='8000', resource_deps=['prometheus-stack'])
