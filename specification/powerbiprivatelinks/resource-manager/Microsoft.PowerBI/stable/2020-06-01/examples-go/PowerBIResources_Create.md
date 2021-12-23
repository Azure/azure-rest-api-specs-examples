Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbiprivatelinks%2Farmpowerbiprivatelinks%2Fv0.1.0/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/README.md) on how to add the SDK to your project and authenticate.

```go
package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks"
)

// x-ms-original-file: specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PowerBIResources_Create.json
func ExamplePowerBIResourcesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpowerbiprivatelinks.NewPowerBIResourcesClient("<subscription-id>",
		"<resource-group-name>",
		"<azure-resource-name>", cred, nil)
	res, err := client.Create(ctx,
		armpowerbiprivatelinks.TenantResource{
			Location: to.StringPtr("<location>"),
			Properties: &armpowerbiprivatelinks.TenantProperties{
				TenantID: to.StringPtr("<tenant-id>"),
			},
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
		},
		&armpowerbiprivatelinks.PowerBIResourcesCreateOptions{ClientTenantID: to.StringPtr("<client-tenant-id>")})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TenantResource.ID: %s\n", *res.ID)
}
```
