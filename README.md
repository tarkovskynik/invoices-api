# REST API For Invoices App

Use command: `git clone https://github.com/tarkovskynik/invoices-api.git`

## Dependencies:
- Docker & docker-compose
- golang-migrate (brew install golang-migrate)

## Build & Run:
Use this command in the directory
- `make build`
- `make run` (if you have an "Error connecting to Database" during the first command, please enter it again)
- If you are running this app for the first time, use `make migrate` in a new terminal window

## Authorization:
In Postman you choose "Authorization" menu, type "basic auth"

- Login: admin   Password: qwerty

## CRUD operations:
POST - "/invoice" - create invoice

GET - "/invoice/:id" - get invoice by id

GET "/invoices" - get all invoices

PUT "/invoice/:id" - update invoice by id

DELETE "/invoice/:id" - delete invoice by id

## Example of creating an invoice:
In Postman you choose "Body" menu, POST `localhost:8080/invoice` and type for example:

{  
"title": "inv-0101",  
"description": "for rent office",  
"company_name": "Advantio",  
"date": "06.05.2020",  
"total_cost": 10000  
}
