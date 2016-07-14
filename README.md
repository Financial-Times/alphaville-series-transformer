# Alphaville Series Transformer

[![CircleCI](https://circleci.com/gh/Financial-Times/alphaville-series-transformer.svg?style=svg)](https://circleci.com/gh/Financial-Times/alphaville-series-transformer)

Retrieves Alphaville Series taxonomy from TME and transforms the series to the internal UP json model.
The service exposes endpoints for getting all the series and for getting series by uuid.

# Usage
`go get -u github.com/Financial-Times/alphaville-series-transformer`

## NB! change `--tme-taxonomy-name="topics"` to `--tme-taxonomy-name="alphavillesries"` when Alphaville Series taxonomy endpoint is exposed by livepub  
`$GOPATH/bin/alphaville-series-transformer --port=8080 --base-url="http://localhost:8080/transformers/alphaville-series/" --tme-base-url="https://tme.ft.com" --tme-username="user" --tme-password="pass" --token="token" --tme-taxonomy-name="topics"`

export|set PORT=8080  
export|set BASE_URL="http://localhost:8080/transformers/alphaville-series/"  
export|set TME_BASE_URL="https://tme.ft.com"  
export|set TME_USERNAME="user"  
export|set TME_PASSWORD="pass"  
export|set TOKEN="token"  
export|set CACHE_FILE_NAME="cache.db"  
$GOPATH/bin/alphaville-series-transformer  

### With Docker:

`docker build -t coco/alphaville-series-transformer .`
## NB! change `"TME_TAXONOMY_NAME=topics"` to `"TME_TAXONOMY_NAME=alphaville-series"` when Alphaville Series taxonomy endpoint is exposed by livepub  
`docker run -ti --env BASE_URL=<base url> --env TME_BASE_URL=<structure service url> --env TME_USERNAME=<user> --env TME_PASSWORD=<pass> --env TOKEN=<token> --env CACHE_FILE_NAME=<file> --env "TME_TAXONOMY_NAME=topics" coco/alphaville-series-transformer`

# Endpoints

* `/transformers/alphaville-series` - Get all alphavillle series as APIURLs
* `/transformers/alphavilles-eries/{uuid}` - Get alphavillle series data of the given uuid
* `/transformers/alphaville-series/__ids` - Get a stream of alphavillle series ids in this format {id : uuid}
* `/transformers/alphaville-series/__count` - Get count of alphavillle series
