# gone-nats
<<<<<<< HEAD
![License: MIT](https://img.shields.io/badge/Language-Golang-blue.svg)
![License: MIT](https://img.shields.io/badge/Database-NATS.io-magenta.svg)
[![Build GO workflow](https://github.com/edwinnduti/gone-nats/actions/workflows/deploy.yaml/badge.svg?branch=master)](https://github.com/edwinnduti/gone-nats/actions/workflows/deploy.yaml)

A REST API that uses the [NATS Message BUS/Broker](https://nats.io).

### Requirements
* Golang 
* Nats server installed( or running on docker)

### Run Locally
```bash
$ git clone github.com/edwinnduti/gone-nats.git
$ cd gone-nats
$ go mod download
$ go build -o natsapp
$ ./natsapp
```

 ### Run code using Docker
 ```
 $ git clone https://github.com/edwinnduti/gone-nats.git 
 $ cd gone-nats
 $ sudo docker build -t natsapp -f Dockerfile .
 $ sudo docker run -it -p 8010:8010 natsapp
 ```

#### Paths
Available :

| function              |   path                    |   method  |
|   ----                |   ----                    |   ----    |
| welcome new user           |   /			|	GET    |


Made by Edwin with ❤️ in Kenya.
=======
A code that uses the NATS.IO Message BUS/Broker from [nats.io](https://nats.io)
>>>>>>> ee35d641ef08561762e4b98837a040630a9dc01c
