
# select image
FROM golang:alpine

# set environment
ENV GOPROXY https://goproxy.cn,direct

# set environment
WORKDIR /docker/deploy/IPFS-Blog-Hugo

# copy the source code to workdir
COPY . .

# build the binary
RUN go build -o IPFS-Blog-Hugo .

# expose the port
EXPOSE 8000

# run the binary
CMD ["./IPFS-Blog-Hugo"]
