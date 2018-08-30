# Seat Service

## How to run

- Start the server

```
$ go run cmd/server/main.go
// set -rows=<rows> to change the rows. Default to 60.
// set -cols=<cols> to change to seats per row. Default to 10.
```


## APIs

#### Set dimensions
```
curl -XPOST /dimensions?rows=1,2,3,4...&cols=A,B,C,D...
```
**NOTE**: the rows and cols must have the length equal to the rows and cols when starting the server.

#### Get seat name based on the dimensions above
```
curl /name?number=<seat_number>
```


## Testing

```
$ go test
```
