Installing Redis
=======================================================

After updating your system dependencies install redis by
typing the following command on your terminal.
```
$ sudo apt update
$ sudo apt install redis-server
```

now check the status of your redis server.
```
$ sudo systemctl status redis-server
$ redis-cli
```
if redis is setup properly on your machine, the CLI must
give a response as follows.
```
$ 127.0.0.1> ping
\ PONG
```



Clone this repository
```
$ git clone https://github.com/NeerajKadayan/redis.git
```
Install dependencies on your machine
=======================================================
```
$ rm go.mod go.sum
$ go get -u github.com/gin-gonic/gin
$ go get github.com/go-redis/redis/v8
$ go get github.com/gorilla/websocket
$ go get gopkg.in/olahol/melody.v1
$ go mod init
$ go mod tidy
$ go mod vendor (optional)
```

Running the Application
=======================================================
Go to app/ and hit :
```
$ go run main.go
```

open the browser and hit `http://localhost:8080`
You must be greeted with "REDIS MONITOR"

now lets get back to the repository, and hit :
(choose any text editor of your choice)
```
vim ~/server/route.go
```

Now hit the APIs by pasting the string at the end of the URL,
of `localhost:8080` and look at the documentation of the API
to know where you are landed.





***
HAPPY HACKING !!!
***




