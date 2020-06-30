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
curl --location --request POST 'http://localhost:9001/dns/1/drones/1' \
--header 'Content-Type: text/plain' \
--data-raw '{
"x": 123.12,
"y": 456.56,
"z": 789.89,
"vel": 20.0
}'

```
