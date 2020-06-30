# Space_exploration
Space DNS for future 

### Create Sector
```
curl --location --request POST 'http://localhost:9001/sectors' \
--header 'Content-Type: text/plain' \
--data-raw '{
"sector_name" : "sector1"
}'
```

### Create Drones
```
curl --location --request POST 'http://localhost:9001/drones' \
--header 'Content-Type: text/plain' \
--data-raw '{
"name" : "drone1",
"sector_id" : 1
}'
```

### Create DNS
```
curl --location --request POST 'http://localhost:9001/dns' \
--header 'Content-Type: text/plain' \
--data-raw '{
"name" : "dns1",
"sectors" : [1]
}'
```

### Get All Sector Info
```
curl --location --request GET 'http://localhost:9001/sectors'
```
### Get All Drone Info
```
curl --location --request GET 'http://localhost:9001/drones'
```

### Get All DNS Info
```
curl --location --request GET 'http://localhost:9001/dns'
```

### Get location of given drone from given DNS
```
curl --location --request POST 'http://localhost:9001/dns/1/drones/1/location' \
--header 'Content-Type: text/plain' \
--data-raw '{
"x": 123.12,
"y": 456.56,
"z": 789.89,
"vel": 20.0
}'

```
### Get location of given drone from given DNS : custome for friend company
```
curl --location --request POST 'http://localhost:9001/dns/1/drones/1/location?is_custom=true' \
--header 'Content-Type: text/plain' \
--data-raw '{
"x": 123.12,
"y": 456.56,
"z": 789.89,
"vel": 20.0
}'

```

### Run using Compiled Binary
```
cd space_exploration/cmd/space_explore
./space_explore
```

## Running tests

A proper Go environment is required in order to run this project.
Once setup, tests can be run with the following command:
```
cd space_exploration/cmd/space_explore
go test -v ./...
```

### Running with Docker

```
To build the image from the Dockerfile, run:

`docker build -t space_exploration .`

To start an interactive shell, run:

`docker run -p 9001:9001 -it --rm --name space space_exploration`

```