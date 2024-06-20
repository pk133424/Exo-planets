## Exo-planets

### To start application run below comands.
`sudo docker build -t exoplanet-service .`

`docker run -p 8080:8080 exoplanet-service`

### CURLS:
#### Create Exoplanets:

`curl --location 'http://localhost:8080/exoplanets' \
--header 'Content-Type: application/json' \
--data '{
  "name": "String",
  "description": "String",
  "distance": 10-1000,
  "radius": 0.1-10,
  "type": "GasGiant|Terrestrial"
}'`

#### Get All Exoplanets
`curl --location 'http://localhost:8080/exoplanets' \
--header 'Content-Type: application/json'`

#### GetByID Exoplanet
`curl --location 'http://localhost:8080/exoplanets/{ID}' \
--header 'Content-Type: application/json'`


#### Update Exoplanet
`curl --location --request PUT 'http://localhost:8080/exoplanets/4a73fda1-d04d-4419-afa9-fcc28e9269fb' \
--header 'Content-Type: application/json' \
--data '{
  "name": "String",
  "description": "String",
  "distance": 10-1000,
  "radius": 0.1-10,
  "type": "GasGiant|Terrestrial"
}'`

#### DeleteByID Exoplanet
`curl --location --request DELETE 'http://localhost:8080/exoplanets/7bc2740a-70d0-41e6-a6a7-a8eae18e2d34' \
--header 'Content-Type: application/json'`


#### Get Fuel Estimation
`curl --location 'http://localhost:8080/exoplanets/{ID}/fuel/{INT}' \
--header 'Content-Type: application/json'`

