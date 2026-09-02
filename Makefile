swagger?=docker run --rm -it --env GOPATH=$(HOME)/go:/go --user $$(id -u):$$(id -g) --volume $(HOME):$(HOME) --workdir $$(pwd) quay.io/goswagger/swagger:v0.36.5
# The NetBox OpenAPI 3 document the client is built from, downloaded from /api/schema/.
v3spec?=$(shell ls NetBox*.yaml 2>/dev/null | head -1)

# swagger.processed.json is the OpenAPI 2 conversion of the v3 document, rebuilt on every run and
# not tracked. --name pins the generated client type to NetBoxAPI; without it go-swagger derives
# the name from info.title, which is NetBox's ("NetBox REST API") and would rename the type on
# every wording change.
generate: spec
	mkdir -p netbox
	$(swagger) generate client --name="NetBox API" --target=./netbox --spec=./swagger.processed.json --copyright-file=./copyright_header.txt --skip-validation

spec:
	python3 tools/v3migrate/migrate.py --v3 "$(v3spec)"

# The drift test: the tracked client must be what the tracked document generates.
check: clean generate
	git diff --exit-code --stat -- netbox/

clean:
	rm -rf netbox/

integration:
	go test ./... -tags=integration

.PHONY: generate spec check clean integration
