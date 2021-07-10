# CDS Backend Engineer Challenge

To run locally:  
`go run main.go`  
To run tests:  
`go test ./...`  
To build and run docker image:

```bash
docker build -t apptest .
docker run -it --rm -p 8080:8080 apptest
```

## Example Queries  

### Geolocation Search

`long` must be in the range [-180,180]. `lat` must be in the range [-90,90]. `radius` is optional and in metres (defaults to 0).  
<http://localhost:8080/nearby?long=-122.39006184632663&lat=37.722629217598346>  
<http://localhost:8080/nearby?long=-122.388369&lat=37.725498&radius=90>  

### Autocomplete Applicant

`term` is not case sensitive. Returns all applicants that start with `term`.  
<http://localhost:8080/autocomplete/applicant?term=Ev>

### Autocomplete Address

`term` is not case sensitive. Returns all addresses that contain `term`.  
<http://localhost:8080/autocomplete/address?term=MINNE>