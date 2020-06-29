# Space_exploration
Space DNS for future 

### Create Sector
```
curl --location --request POST 'http://localhost:9001/sectors' \
--header 'Content-Type: text/plain' \
--data-raw '{
"sector_id" : 1,
"sector_name" : "sector1"
}'
```

### Create Drones
```
curl --location --request POST 'http://localhost:9001/drones' \
--header 'Content-Type: text/plain' \
--data-raw '{
"id" : 1,
"name" : "drone1",
"sector_id" : 1
}'
```

### Create DNS
```
curl --location --request POST 'http://localhost:9001/dns' \
--header 'Content-Type: text/plain' \
--data-raw '{
"id" : 1,
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

### Get All DNS inInfofo
```
curl --location --request GET 'http://localhost:9001/dns'
```
