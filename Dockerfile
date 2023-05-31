FROM centos:7

WORKDIR /var/www/IPFS-Blog-Hugo

COPY . .

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64



EXPOSE 8000

CMD ["./IPFS-Blog-Hugo"]