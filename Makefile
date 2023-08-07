OPENAPI-GENERATOR=docker run --rm --user $$(id -u):$$(id -g) --volume "${PWD}:/local" openapitools/openapi-generator-cli:v6.6.0
GOIMPORTS=goimports

generate:
	mkdir -p netbox
	$(OPENAPI-GENERATOR) generate -g go -i /local/openapi.yaml -o /local/netbox --git-host github.com --git-repo-id go-netbox --git-user-id fbreckle --package-name netbox # --enable-post-process-file
	$(GOIMPORTS) -w netbox/*.go netbox/test/*.go

preprocess:
	python3 preprocess.py

clean:
	rm -rf netbox/
