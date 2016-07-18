# boltdbweb
a simple web based boltdb GUI Admin panel


##### Installation
```
$ go get github.com/gin-gonic/gin
$ go get github.com/boltdb/bolt/...
$ go get github.com/evnix/boltdbweb

cd boltdbweb
go build boltdbweb.go

To run:
./boltdbweb file.db 8080
Goto: http://localhost:8080
NOTE: If 'file.db' does not exist. it will be created as a BoltDB file.
```

##### Screenshots:

![](https://github.com/evnix/boltdbweb/blob/master/screenshots/1.png?raw=true)

![](https://github.com/evnix/boltdbweb/blob/master/screenshots/2.png?raw=true)

![](https://github.com/evnix/boltdbweb/blob/master/screenshots/3.png?raw=true)
