Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fextendedlocation%2Farmextendedlocation%2Fv0.2.1/sdk/resourcemanager/extendedlocation/armextendedlocation/README.md) on how to add the SDK to your project and authenticate.

```go
package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsPatch.json
func ExampleCustomLocationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armextendedlocation.NewCustomLocationsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armextendedlocation.PatchableCustomLocations{
			Identity: &armextendedlocation.Identity{
				Type: armextendedlocation.ResourceIdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armextendedlocation.CustomLocationProperties{
				ClusterExtensionIDs: []*string{
					to.StringPtr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension"),
					to.StringPtr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/barExtension")},
			},
			Tags: map[string]*string{
				"archv3": to.StringPtr(""),
				"tier":   to.StringPtr("testing"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CustomLocationsClientUpdateResult)
}
```
