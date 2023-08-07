package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	openapiclient "github.com/fbreckle/go-netbox/netbox/netbox"
)

func main() {
	configuration := openapiclient.NewConfiguration()
	// configuration.Scheme = "http"
	// configuration.Host = "localhost:8001"
	configuration.Servers[0].URL = "http://localhost:8001"
	spew.Dump(configuration)
	ctx := context.Background()

	apiKey := openapiclient.APIKey{
		Key: "0123456789abcdef0123456789abcdef01234567", Prefix: "Token",
	}

	ctx2 := context.WithValue(ctx, openapiclient.ContextAPIKeys, map[string]openapiclient.APIKey{"tokenAuth": apiKey})
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, _, err := apiClient.DcimApi.DcimSitesList(ctx2).Execute()
	if err != nil {
		fmt.Printf("Error when calling `DcimApi.DcimSitesList``: %v\n", err)
		// fmt.Printf("Full HTTP response: %v\n", r)
	}
	// response from `DcimSitesList`: PaginatedSiteList
	fmt.Printf("Response from `DcimApi.DcimSitesList`: %v\n", resp)
}
