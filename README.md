## Gin Playground [For Review Project]
The purpose of this repo is
- To verify designed tables [with regard to review] appears in real RDB schema as intended
- To test NoSQL version of it
- To PoC or test any golang/gin related functionality available.

## Setup

make sure $GOPATH is valid.

```
make setup
mkdir -p $GOPATH/src/github.com/keiveltra/gin-playground && cd "$_"
curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
```

After setup the repo, please setup the mysql:
```
brew install mysql
CREATE USER 'moomin'@'localhost' IDENTIFIED BY 'moomin';
GRANT ALL PRIVILEGES ON *.* TO 'moomin'@'localhost' WITH GRANT OPTION;
```
