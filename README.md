# Packform Dev Test Task

## Requirements
<br>

### Having the following:
- PostgreSQL database
- CSV files with test data for the DB
### The task is to develop:
- A script (Golang or Python preferred) to initialize the database and populate them
    with the data from the CSV
- A simple web application to display this data
<br><br>

## Quick Start
### Required setup before running this project
- Add your local variable to the .env at root folder, you can follow the <b>.envExample</b> file for reference.
- Golang is installed at machine(GO 1.17.6 was used)
- Python is installed at machine(Python 3.9.4 was used)
- Postgres db is installed and running on the target machine
- Vue cli is installed at machine(vue/cli 4.5.15 was used)
<br/><br/>

### For server(Go)
```bash
# Run database initialization script
$ cd scripts
$ python -m pip install python-dotenv
$ python .\database_script.py

# Install dependencies for server
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/gin-contrib/cors
$ go get -u github.com/joho/godotenv
$ go get -u gorm.io/driver/postgres
$ go get -u gorm.io/gorm

# Install dependencies for client
$ cd client
$ npm install

# Run the client & server with concurrently
$ npm run serve

# Run the Golang server only
$ go run main.go

# Run the React client only
$ npm run serve

# Server runs on http://localhost:5000 and client on http://localhost:8080
```
## Acceptance criteria
As a User
- When I open the /orders page
- Then I should see the list of all customer order
- And I should be able to search orders by part of the order or product name
- And I should be able to filter orders by date range (Melbourne/Australia TZ)
- And I should see maximum 5 orders per page
- And I should be able to switch between pages
- And for every order the following information should be displayed:
    - Order name
    - Customer Company name
    - Full-stack Software Engineer test 2
    - Customer name
    - Order date (Melbourne/Australia TZ)
    - Delivered amount (dash if nothing is delivered)
    - Total amount

## Tech requirements
- [x] UI should be implemented as a single page web application using Vue.js
- [x] Backend application should be implemented in Golang
- [x] There should be a Readme file with clear setup instructions
- [ ] (Bonus) Unit / functional / integration tests

## App Info

### Author

Ivan Chan
[Github](https://github.com/IvanCMC37/)

### Version

1.0.0

### License

This project is licensed under the MIT License
