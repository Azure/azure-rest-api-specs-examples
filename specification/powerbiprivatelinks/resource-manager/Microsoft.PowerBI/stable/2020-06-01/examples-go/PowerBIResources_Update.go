package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PowerBIResources_Update.json
func ExamplePowerBIResourcesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerbiprivatelinks.NewPowerBIResourcesClient("a0020869-4d28-422a-89f4-c2413130d73c",
		"resourceGroup",
		"azureResourceName", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		armpowerbiprivatelinks.TenantResource{
			Location: to.Ptr("global"),
			Properties: &armpowerbiprivatelinks.TenantProperties{
				TenantID: to.Ptr("ac2bc297-8a3e-46f3-972d-87c2b4ae6e2f"),
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		&armpowerbiprivatelinks.PowerBIResourcesClientUpdateOptions{ClientTenantID: to.Ptr("ac2bc297-8a3e-46f3-972d-87c2b4ae6e2f")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
