package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2020-03-01/examples/DPSCheckNameAvailability.json
func ExampleIotDpsResourceClient_CheckProvisioningServiceNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeviceprovisioningservices.NewIotDpsResourceClient("<subscription-id>", cred, nil)
	res, err := client.CheckProvisioningServiceNameAvailability(ctx,
		armdeviceprovisioningservices.OperationInputs{
			Name: to.StringPtr("<name>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IotDpsResourceClientCheckProvisioningServiceNameAvailabilityResult)
}
