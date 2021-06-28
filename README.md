# manifest-gen

`manifest-gen` is a tool to generate argo and helm manifests on the fly


## Usage

some useful commmands
```shell
$ make lint # run golangci-lint
$ make clean # clean binaries in ./bin
$ make build # compile binary to ./bin/manifest-gen
$ make install # install binary to ~/go/bin/manifest-gen
```

creating helm manifests
```shell
./bin/manifest-gen create helm --name sanjay --image "docker.io/nginx" --image-pull-policy "Always" --image-tag "latest" --service-type "ClusterIP" --service-port 8080
```
