### Development

# install golang
# https://golang.org/doc/install

setup.sh

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t example-scratch -f Dockerfile.scratch .
docker run -it example-scratch