FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-debug-zone"]
COPY baton-debug-zone /