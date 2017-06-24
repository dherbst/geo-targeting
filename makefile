

GOLANG := golang:1.8

test:
	docker run --rm -e GOPATH=/ -v ${PWD}/src:/src -w /src ${GOLANG} go test geo
