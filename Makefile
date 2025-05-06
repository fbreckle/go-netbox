swagger=docker run --rm -it --env GOPATH=$(HOME)/go:/go --user $$(id -u):$$(id -g) --volume $(HOME):$(HOME) --workdir $$(pwd) quay.io/goswagger/swagger:v0.31.0
generate:
	mkdir -p netbox
	$(swagger) generate client --target=./netbox --spec=./swagger.processed.json --copyright-file=./copyright_header.txt --skip-validation

clean:
	rm -rf netbox/

integration:
	go test ./... -tags=integration
