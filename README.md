# boltdbweb
A simple web based boltdb GUI Admin panel.


##### Installation
```
go get github.com/evnix/boltdbweb
```

##### Usage
```
boltdbweb --db-name=<DBfilename>[required] --port=<port>[optional] --static-path=<static-path>[optional]
```
- `--db-name:` The file name of the DB.
    - NOTE: If 'file.db' does not exist. it will be created as a BoltDB file.
- `--port:` Port for listening on... (Default: 8080)
- `--static-path:` If you moved the binary to different folder you can determin the path of the `web` folder. (Default: Same folder where the binary is located.)


##### Example
```
boltdbweb --db-name=test.db --port=8089 --static-path=/home/user/github/boltdbweb
```
Goto: http://localhost:8089

##### Screenshots:

![](https://github.com/evnix/boltdbweb/blob/master/screenshots/1.png?raw=true)

![](https://github.com/evnix/boltdbweb/blob/master/screenshots/2.png?raw=true)

![](https://github.com/evnix/boltdbweb/blob/master/screenshots/3.png?raw=true)
