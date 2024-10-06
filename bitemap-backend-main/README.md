# bitemap-backend
Backend code for the bitemap application.

## Steps to setup the project locally
Excecute the commands sequentially in order to successfully run the application locally on a Mac OS

### Install Golang
```
brew install golang

go version
```
### Install other required packages
```
brew install golang-migrate

go install -v github.com/golang/mock/mockgen@v1.6.0
```
### Add PATH to shell
- open your .zshrc or .bashrc file.
- add - export PATH=$PATH:~/go/bin
- save and close the file
- open terminal and excute 
```
source ~/.zshrc
```

### Setup Database
Excecute following commands on the terminal after installing postgres locally or using docker - 
```
psql -U postgres

create database bitemap;
```

#### Run Database migrations
```
make migrateup
``` 

### Run the server
```
make server
```

### Build Docker image
```
docker build -t bitemap/bitemap-backend:latest .
```

### Run as Docker container
```
docker run \
      --name bitemap \
      --rm -it \
      -v /:/host:ro \
      -v /var/run/docker.sock:/var/run/docker.sock:ro \
      --privileged \
      --pid=host \
      --network=host \
      docker pull bitemap/bitemap-backend:latest
```





