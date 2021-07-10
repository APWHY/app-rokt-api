# CDS Backend Engineer Challenge

To run locally:
`go mod vendor`

## Example Queries  

### Geolocation Search

`long` must be in the range [-180,180]. `lat` must be in the range [-90,90]. `radius` is optional and in metres (defaults to 0).
<http://localhost:8080/nearby?long=-122.39006184632663&lat=37.722629217598346>  
<http://localhost:8080/nearby?long=-122.388369&lat=37.725498&radius=90>