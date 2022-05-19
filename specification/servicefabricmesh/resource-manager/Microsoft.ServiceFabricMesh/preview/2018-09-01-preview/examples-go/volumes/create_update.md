Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabricmesh%2Farmservicefabricmesh%2Fv0.5.0/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/volumes/create_update.json
func ExampleVolumeClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicefabricmesh.NewVolumeClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"sbz_demo",
		"sampleVolume",
		armservicefabricmesh.VolumeResourceDescription{
			Location: to.Ptr("EastUS"),
			Tags:     map[string]*string{},
			Properties: &armservicefabricmesh.VolumeResourceProperties{
				Description: to.Ptr("Service Fabric Mesh sample volume."),
				AzureFileParameters: &armservicefabricmesh.VolumeProviderParametersAzureFile{
					AccountKey:  to.Ptr("provide-account-key-here"),
					AccountName: to.Ptr("sbzdemoaccount"),
					ShareName:   to.Ptr("sharel"),
				},
				Provider: to.Ptr(armservicefabricmesh.VolumeProviderSFAzureFile),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
