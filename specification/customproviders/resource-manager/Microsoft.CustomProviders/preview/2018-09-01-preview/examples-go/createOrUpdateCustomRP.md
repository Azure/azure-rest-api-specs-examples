Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcustomproviders%2Farmcustomproviders%2Fv0.1.0/sdk/resourcemanager/customproviders/armcustomproviders/README.md) on how to add the SDK to your project and authenticate.

```go
package armcustomproviders_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders"
)

// x-ms-original-file: specification/customproviders/resource-manager/Microsoft.CustomProviders/preview/2018-09-01-preview/examples/createOrUpdateCustomRP.json
func ExampleCustomResourceProviderClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcustomproviders.NewCustomResourceProviderClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-provider-name>",
		armcustomproviders.CustomRPManifest{
			Resource: armcustomproviders.Resource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &armcustomproviders.CustomRPManifestProperties{
				Actions: []*armcustomproviders.CustomRPActionRouteDefinition{
					{
						CustomRPRouteDefinition: armcustomproviders.CustomRPRouteDefinition{
							Name:     to.StringPtr("<name>"),
							Endpoint: to.StringPtr("<endpoint>"),
						},
						RoutingType: armcustomproviders.ActionRoutingProxy.ToPtr(),
					}},
				ResourceTypes: []*armcustomproviders.CustomRPResourceTypeRouteDefinition{
					{
						CustomRPRouteDefinition: armcustomproviders.CustomRPRouteDefinition{
							Name:     to.StringPtr("<name>"),
							Endpoint: to.StringPtr("<endpoint>"),
						},
						RoutingType: armcustomproviders.ResourceTypeRoutingProxyCache.ToPtr(),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CustomRPManifest.ID: %s\n", *res.ID)
}
```
