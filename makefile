.PHONY=test

GOLANG := golang:1.12

test:
	docker run --rm -v ${PWD}:/go/src/github.com/dherbst/geo-targeting -w /go/src/github.com/dherbst/geo-targeting ${GOLANG} go test -coverprofile=coverage.out github.com/dherbst/geo-targeting/src/geo
	docker run --rm -v ${PWD}:/go/src/github.com/dherbst/geo-targeting -w /go/src/github.com/dherbst/geo-targeting ${GOLANG} go tool cover -html=coverage.out -o coverage.html
