YQ="yq"

# Fix wrong types: (Issue: https://github.com/netbox-community/netbox/issues/11578)
${YQ} -i '.components.schemas.IPAddress.properties.family.properties.value.type = "integer" ' openapi.yaml
${YQ} -i '.paths["/api/ipam/prefixes/{id}/available-ips/"].post.requestBody.content["application/json"].schema.type = "array" ' openapi.yaml
${YQ} -i 'del(.paths["/api/ipam/prefixes/{id}/available-ips/"].post.requestBody.content["application/json"].schema["$ref"])' openapi.yaml
${YQ} -i '.paths["/api/ipam/prefixes/{id}/available-ips/"].post.requestBody.content["application/json"].schema.items["$ref"] = "#/components/schemas/WritableIPAddressRequest" ' openapi.yaml

# Set the default format for all integers to int64 if not specified otherwise.
${YQ} -i 'with(.components.schemas ; .[] | with(.properties[]; select( .type == "integer" and .format == null) | .format="int64"))' openapi.yaml
${YQ} -i 'with(.paths[]; with(.[]; with( .parameters[]; select(.schema.type=="integer" and .schema.format!="int64") | .schema.format="int64" )))' openapi.yaml

# Delete old generated file:
rm -rf client.gen.go

# Regenerate, but without formating, since the spec generates int constants with value <nil>. Which is not valid go code.
${topLevelDirectory}/bin/oapi-codegen --exclude-tags virtualization,wireless,circuits,schema,status,users -generate types,client,skip-fmt -package openapi openapi.yaml > client.gen.go

# Remove <nil> integers.
sed -i "/= <nil>/d" client.gen.go

# Format code
go fmt client.gen.go

# Remove unused imports
${topLevelDirectory}/bin/goimports -w client.gen.go
