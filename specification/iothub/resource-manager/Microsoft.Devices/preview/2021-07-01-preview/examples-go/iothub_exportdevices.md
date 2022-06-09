```go
package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2021-07-01-preview/examples/iothub_exportdevices.json
func ExampleResourceClient_ExportDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiothub.NewResourceClient("<subscription-id>", cred, nil)
	res, err := client.ExportDevices(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armiothub.ExportDevicesRequest{
			AuthenticationType:     armiothub.AuthenticationType("identityBased").ToPtr(),
			ExcludeKeys:            to.BoolPtr(true),
			ExportBlobContainerURI: to.StringPtr("<export-blob-container-uri>"),
			Identity: &armiothub.ManagedIdentity{
				UserAssignedIdentity: to.StringPtr("<user-assigned-identity>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ResourceClientExportDevicesResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiothub%2Farmiothub%2Fv0.3.1/sdk/resourcemanager/iothub/armiothub/README.md) on how to add the SDK to your project and authenticate.
