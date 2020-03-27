# ardan
web rapid development framework 

## Deps
- [Gin](https://gin-gonic.com/) HTTP web framework
- [XORM](https://xorm.io/) ORM Framework 
- [Container](https://github.com/teamlint/container) Dependency injection container

## CLI
ardan cli tools

### Install

#### Go get

```shell
go get github.com/teamlint/ardan
```

#### OR Git

```shell
git clone git@github.com:teamlint/ardan.git
cd ardan
go install
```

### Usages

#### Init Project
```go
mkdir myproject
cd myproject
go mod init <name>
ardan -s init  
```
#### Project layout

![](https://github.com/teamlint/ardan/blob/master/screenshots/layout.png?raw=true)

#### Run 

```shell
task run
// or
go run cmd/server/main.go
```

### Model Sync

#### Create model file.

Write model file in `app/model/tom.go`

```go
package model

import (
	"time"
)

//ardan:sync
//ardan:gen --service TomService --repository 
// Tom test model
type Tom struct{
	ID         string     `xorm:"not null pk unique CHAR(20) 'id'" json:"id"`
	Username   string     `xorm:"not null index VARCHAR(100)" json:"username"`
	CreatedAt  time.Time  `xorm:"not null created TIMESTAMPZ" json:"created_at"`
}
```
```shell
ardan -dc <databae-conn-str> sync 
```

#### view database

have fun!.

![](https://github.com/teamlint/ardan/blob/master/screenshots/sync_db.png?raw=true)

### Generate Tools

```shell
ardan gen 
```

## TODO

- new command (controller/middleware/module...)

