package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Firmwares_Create_MinimumSet_Gen.json
func ExampleFirmwaresClient_Create_firmwaresCreateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirmwaresClient().Create(ctx, "rgworkspaces-firmwares", "A7", "umrkdttp", armiotfirmwaredefense.Firmware{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Firmware = armiotfirmwaredefense.Firmware{
	// }
}
