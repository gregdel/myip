FROM alpine:3.6 as downloader
ARG TAG=v0.2
ARG OS=linux
ARG ARCH=amd64
RUN apk --no-cache add wget ca-certificates && \
    wget https://github.com/gregdel/myip/releases/download/${TAG}/myip_${OS}_${ARCH} && \
    mv myip_${OS}_${ARCH} /app && \
    chmod +x /app

FROM scratch
COPY --from=downloader /app /app
EXPOSE 8080
CMD ["/app", "-ip", "0.0.0.0", "-port", "8080"]
