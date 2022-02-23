Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fextendedlocation%2Farmextendedlocation%2Fv0.2.1/sdk/resourcemanager/extendedlocation/armextendedlocation/README.md) on how to add the SDK to your project and authenticate.

```go
package armextendedlocation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsCreate_Update.json
func ExampleCustomLocationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armextendedlocation.NewCustomLocationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armextendedlocation.CustomLocation{
			Location: to.StringPtr("<location>"),
			Identity: &armextendedlocation.Identity{
				Type: armextendedlocation.ResourceIdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armextendedlocation.CustomLocationProperties{
				Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
					Type:  to.StringPtr("<type>"),
					Value: to.StringPtr("<value>"),
				},
				ClusterExtensionIDs: []*string{
					to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedCluster/someCluster/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
				DisplayName:    to.StringPtr("<display-name>"),
				HostResourceID: to.StringPtr("<host-resource-id>"),
				Namespace:      to.StringPtr("<namespace>"),
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
	log.Printf("Response result: %#v\n", res.CustomLocationsClientCreateOrUpdateResult)
}
```
