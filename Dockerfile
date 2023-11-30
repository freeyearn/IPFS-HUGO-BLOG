# select image
FROM golang:alpine

# set environment
ENV GOPROXY https://goproxy.cn,direct

# create workdir
RUN mkdir /app 

# copy the source code to workdir
ADD . /app/

# change the workdir
WORKDIR /app

# build the binary
RUN go build -o IPFS-Blog-Hugo .

# expose the port
EXPOSE 8000

# run the binary
CMD ["./IPFS-Blog-Hugo"]