Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv0.3.0/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.

```go
package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/PrivateEndpoint_Create.json
func ExamplePrivateEndpointsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewPrivateEndpointsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<private-endpoint-name>",
		armstreamanalytics.PrivateEndpoint{
			Properties: &armstreamanalytics.PrivateEndpointProperties{
				ManualPrivateLinkServiceConnections: []*armstreamanalytics.PrivateLinkServiceConnection{
					{
						Properties: &armstreamanalytics.PrivateLinkServiceConnectionProperties{
							GroupIDs: []*string{
								to.StringPtr("groupIdFromResource")},
							PrivateLinkServiceID: to.StringPtr("<private-link-service-id>"),
						},
					}},
			},
		},
		&armstreamanalytics.PrivateEndpointsClientCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointsClientCreateOrUpdateResult)
}
```
