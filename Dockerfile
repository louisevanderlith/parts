FROM scratch

COPY cmd/cmd .

EXPOSE 8097

ENTRYPOINT [ "./cmd" ]