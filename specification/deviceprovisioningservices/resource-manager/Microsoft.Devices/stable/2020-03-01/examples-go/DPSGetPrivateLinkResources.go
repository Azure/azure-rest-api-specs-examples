package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2020-03-01/examples/DPSGetPrivateLinkResources.json
func ExampleIotDpsResourceClient_GetPrivateLinkResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeviceprovisioningservices.NewIotDpsResourceClient("<subscription-id>", cred, nil)
	res, err := client.GetPrivateLinkResources(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<group-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IotDpsResourceClientGetPrivateLinkResourcesResult)
}
