# Gin Playground [For Review Project]

The purpose of this repo is

- To verify designed tables [with regard to review] appears in real RDB schema as intended
- To test NoSQL version of it
- To **PoC** or test any golang/gin related functionality available.

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
$ make mig
```

To run the webapp,

```
$ make run
```

## data fixture

`db_init.go` is responsible for the data fostering.

it reads test/fixtures/*.yaml

**This repo is POC of testing data design requirement and whether it works in gin+golang, no other further stuff is in scope.**

You can choose whatever fixture(aka test data) fostering protocol or any testing framework. This repo does not cover those.

## TODO

```
- review.+category_id
- view? 追加
```
