package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2020-03-01/examples/DPSGetOperationResult.json
func ExampleIotDpsResourceClient_GetOperationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeviceprovisioningservices.NewIotDpsResourceClient("<subscription-id>", cred, nil)
	res, err := client.GetOperationResult(ctx,
		"<operation-id>",
		"<resource-group-name>",
		"<provisioning-service-name>",
		"<asyncinfo>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IotDpsResourceClientGetOperationResultResult)
}
