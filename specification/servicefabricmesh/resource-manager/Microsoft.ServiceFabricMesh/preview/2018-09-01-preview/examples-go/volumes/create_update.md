Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabricmesh%2Farmservicefabricmesh%2Fv0.2.1/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/volumes/create_update.json
func ExampleVolumeClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabricmesh.NewVolumeClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<volume-resource-name>",
		armservicefabricmesh.VolumeResourceDescription{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armservicefabricmesh.VolumeResourceProperties{
				Description: to.StringPtr("<description>"),
				AzureFileParameters: &armservicefabricmesh.VolumeProviderParametersAzureFile{
					AccountKey:  to.StringPtr("<account-key>"),
					AccountName: to.StringPtr("<account-name>"),
					ShareName:   to.StringPtr("<share-name>"),
				},
				Provider: armservicefabricmesh.VolumeProvider("SFAzureFile").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VolumeClientCreateResult)
}
```
