## Gin Playground [For Review Project]
The purpose of this repo is
- To verify designed tables [with regard to review] appears in real RDB schema as intended
- To test NoSQL version of it
- To PoC or test any golang/gin related functionality available.

## Setup

execute the following command:
```
$ make setup
```

After setup the repo, please setup the mysql:
```
brew install mysql
CREATE USER 'moomin'@'localhost' IDENTIFIED BY 'moomin';
GRANT ALL PRIVILEGES ON *.* TO 'moomin'@'localhost' WITH GRANT OPTION;
```

## Usage

First you need to make sure you have no mysql database named 'test'.
(Since it is going to be truncated)

```
$ make migrate
```

To run the webapp,

```
$ make run
```
