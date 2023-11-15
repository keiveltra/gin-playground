setup:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql

compile:
	go build -o webapp

run:
	go build -o webapp
	./webapp $(arg)

migrate:
	mysql -u moomin -pmoomin -e "drop database test; create database test"
	go build -o webapp
	./webapp m

curl:
	curl localhost:8080

start-mysql:
	brew services start mysql

mysql:
	mysql -u moomin -pmoomin
