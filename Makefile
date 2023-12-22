setup:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql
	go get github.com/spf13/viper
	go get gopkg.in/yaml.v3
	go get github.com/davecgh/go-spew/spew

compile:
	go build -o webapp

run:
	go build -o webapp
	./webapp $(arg)

mig:
	mysql -u moomin -pmoomin -e "drop database if exists test;"
	mysql -u moomin -pmoomin -e "create database test CHARACTER SET utf8mb4;"
	go build -o webapp
	./webapp m
	make drop_auto_gen_cols

inject:
	# works only if this workdir and review-service's workdir is on the same level
	mysql -u moomin -pmoomin -e "drop database if exists test;"
	mysql -u moomin -pmoomin -e "create database test;"
	echo '--------------------------'
	dumpfilefixer.py ../review-service/database/reviews_schema_only_2023-12-21.sql
	mysql -u moomin -pmoomin test < ../review-service/database/reviews_schema_only_2023-12-21.sql

drop_auto_gen_cols:
	mysql -u moomin -pmoomin -e "alter table test.votes drop column deleted_at;"

mi: # mig inject
	make mig; sqd -p; make inject
		
min: # mig inject
	make mig; sqd; make inject

miu: 
	make mig; ./scripts/update_comments.sh 

curl:
	curl localhost:8080

start-mysql:
	brew services start mysql

mysql:
	mysql -u moomin -pmoomin
