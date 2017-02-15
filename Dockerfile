FROM scratch

ADD kubernetes-pod-discovery /kubernetes-pod-discovery

CMD ["/kubernetes-pod-discovery"]
