package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSCheckNameAvailability.json
func ExampleIotDpsResourceClient_CheckProvisioningServiceNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIotDpsResourceClient().CheckProvisioningServiceNameAvailability(ctx, armdeviceprovisioningservices.OperationInputs{
		Name: to.Ptr("test213123"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NameAvailabilityInfo = armdeviceprovisioningservices.NameAvailabilityInfo{
	// 	Message: to.Ptr("name is valid"),
	// 	NameAvailable: to.Ptr(true),
	// 	Reason: to.Ptr(armdeviceprovisioningservices.NameUnavailabilityReasonInvalid),
	// }
}
