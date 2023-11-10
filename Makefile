setup:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql

compile:
	go build -o webapp

run:
	go build -o webapp
	./webapp

curl:
	curl localhost:8080
