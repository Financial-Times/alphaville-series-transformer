# Alphaville Series Transformer

[![CircleCI](https://circleci.com/gh/Financial-Times/alphaville-series-transformer.svg?style=svg)](https://circleci.com/gh/Financial-Times/alphaville-series-transformer)

Retrieves Alphaville Series taxonomy from TME and transforms the series to the internal UP json model.
The service exposes endpoints for getting all the series and for getting series by uuid.

# Usage
`go get -u github.com/Financial-Times/alphaville-series-transformer`

`$GOPATH/bin/alphaville-series-transformer --port=8080 --base-url="http://localhost:8080/transformers/alphavillleseries/" --tme-base-url="https://tme.ft.com" --tme-username="user" --tme-password="pass" --token="token"`

```
export|set PORT=8080
export|set BASE_URL="http://localhost:8080/transformers/alphavillleseries/"
export|set TME_BASE_URL="https://tme.ft.com"
export|set TME_USERNAME="user"
export|set TME_PASSWORD="pass"
export|set TOKEN="token"
export|set CACHE_FILE_NAME="cache.db"
$GOPATH/bin/alphaville-series-transformer
```

### With Docker:

`docker build -t coco/alphaville-series-transformer .`

`docker run -ti --env BASE_URL=<base url> --env TME_BASE_URL=<structure service url> --env TME_USERNAME=<user> --env TME_PASSWORD=<pass> --env TOKEN=<token> --env CACHE_FILE_NAME=<file> coco/alphaville-series-transformer`

# Endpoints

* `/transformers/alphavillleseries` - Get all alphavillle series as APIURLs
* `/transformers/alphavillleseries/{uuid}` - Get alphavillle series data of the given uuid
