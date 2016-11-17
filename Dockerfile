FROM scratch
ADD bin/ozuio_linux_amd64 /ozuio
EXPOSE 8080
VOLUME /data
WORKDIR /
CMD ["/ozuio", "--port", "8080", "--dir", "data"]
