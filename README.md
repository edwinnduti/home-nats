# gone-nats

![License: MIT](https://img.shields.io/badge/Language-Golang-blue.svg)
![License: MIT](https://img.shields.io/badge/Database-NATS.io-magenta.svg)
[![Build GO workflow](https://github.com/edwinnduti/gone-nats/actions/workflows/deploy.yaml/badge.svg?branch=master)](https://github.com/edwinnduti/gone-nats/actions/workflows/deploy.yaml)

A REST API that uses the [NATS Message BUS/Broker](https://nats.io).

### Requirements
* Golang 
* Nats server installed( or running on docker)
* MySQL

# Note
- You will need to edit the <b>.env</b> file in the gone-nats directory with the credential details and paste it in the gone-nats directory.
- Also start the NATS server using:
    ```bash
    $ nats-server
    ```

### Run Locally
```bash
$ git clone github.com/edwinnduti/gone-nats.git
$ cd gone-nats
$ mysql -u <username> -p database_name_goes_here < db/Housedb.sql # database_name is houseinfodb
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

| function                   |   path               |   method  |
|   ----                     |   ----               |   ----    |
| welcome new user           |   /			        |	GET     |
| add new house to db        |   /add-house	        |	POST    |
| get house based on id      |   /house/{house_id}	|	GET     |



Happy Coding!

Made by Edwin with ❤️ in Kenya.
