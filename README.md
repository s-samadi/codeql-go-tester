## CodeQL Go Example

### Create the CodeQL database
```
export PATH="$HOME/local/vmsync/codeql263:$PATH"
codeql database create --language="go" ./sample-db --overwrite
```

## Input Validation

A simple example of how to do input validation in Go.

## API

*GET: /user/{id}*  
params:  
path: id - account ID  
header: Authorization - authorization ID

*POST: /user*  
params:  
body: userId JSON, e.g
```
{
    "name": "peter",
    "mobile_number": "040312345"
}
```

## Example Usage

Build and run the server on localhost port 8000
```
$ go run .
```

### Create a User

Send a JSON payload using HTTPie:
```
http -v POST http://localhost:8000/user Authorization:123 Name="peter" MobileNumber="0403123567"
```

Or with cURL:
```
curl -v --header "Content-Type: application/json" --request POST --data '{"Name":"peter","MobileNumber":"0403123456"}'  http://localhost:8000/user
```

### Get a User

Where `{id}` is from the above POST API.

```
http  GET http://localhost:8000/user/{id} Authorization:123 
```

Or with cURL:
```
curl -v http://localhost:8000/user/{id}
```

## Validation

Send the /user API an invalid payload, e.g.

```
http -v POST http://localhost:8000/user Authorization:123 Name="peter" MobileNumber="040312345x"
```

