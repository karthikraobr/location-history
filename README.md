# Location History

REST server which stores, retrieves and deletes orders' location history

## Requirements
- Docker
- go

## Assumptions and known issues
- PUT is not idempotent - repeated requests with same location and order appends the data. It also creates the history for a new order.
- HISTORY_SERVER_LISTEN_ADDR is assumed be a port number. Done to ease docker setup. 
- Logger isn't used. Can be used in the future.
- Retrieving order_id from the path can fail because it assumes we will always encounter urls of the form `/location/order_id`. This could be easily solved by using a third-party mux like gorilla mux.
- Tests for the server package are missing due to time constraints.
- CI/CD missing.
- Missing API documention (Swagger or OpenAPI).

Makefile contains all the commands to build, run, test this application. 
Contains no external dependencies apart from the go standard library.