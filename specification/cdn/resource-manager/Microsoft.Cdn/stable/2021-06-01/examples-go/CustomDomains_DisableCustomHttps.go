package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/CustomDomains_DisableCustomHttps.json
func ExampleCustomDomainsClient_DisableCustomHTTPS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCustomDomainsClient().DisableCustomHTTPS(ctx, "RG", "profile1", "endpoint1", "www-someDomain-net", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
