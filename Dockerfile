FROM scratch

ADD dist/linux_amd64_kubernetes-pod-discovery /kubernetes-pod-discovery

ENTRYPOINT ["/kubernetes-pod-discovery"]
