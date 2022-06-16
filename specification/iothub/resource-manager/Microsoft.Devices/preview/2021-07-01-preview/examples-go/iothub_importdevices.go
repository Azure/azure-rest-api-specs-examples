package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2021-07-01-preview/examples/iothub_importdevices.json
func ExampleResourceClient_ImportDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiothub.NewResourceClient("<subscription-id>", cred, nil)
	res, err := client.ImportDevices(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armiothub.ImportDevicesRequest{
			InputBlobContainerURI:  to.StringPtr("<input-blob-container-uri>"),
			OutputBlobContainerURI: to.StringPtr("<output-blob-container-uri>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ResourceClientImportDevicesResult)
}
