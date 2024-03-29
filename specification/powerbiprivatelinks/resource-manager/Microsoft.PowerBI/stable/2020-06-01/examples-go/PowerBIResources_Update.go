package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PowerBIResources_Update.json
func ExamplePowerBIResourcesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbiprivatelinks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPowerBIResourcesClient().Update(ctx, "resourceGroup", "azureResourceName", armpowerbiprivatelinks.TenantResource{
		Location: to.Ptr("global"),
		Properties: &armpowerbiprivatelinks.TenantProperties{
			TenantID: to.Ptr("ac2bc297-8a3e-46f3-972d-87c2b4ae6e2f"),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, &armpowerbiprivatelinks.PowerBIResourcesClientUpdateOptions{ClientTenantID: to.Ptr("ac2bc297-8a3e-46f3-972d-87c2b4ae6e2f")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TenantResource = armpowerbiprivatelinks.TenantResource{
	// 	Name: to.Ptr("myPrivateLinkServiceResource"),
	// 	Location: to.Ptr("global"),
	// 	Properties: &armpowerbiprivatelinks.TenantProperties{
	// 		TenantID: to.Ptr("ac2bc297-8a3e-46f3-972d-87c2b4ae6e2f"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// }
}
