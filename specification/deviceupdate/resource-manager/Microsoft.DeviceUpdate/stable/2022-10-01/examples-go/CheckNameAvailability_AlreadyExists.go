package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/CheckNameAvailability_AlreadyExists.json
func ExampleClient_CheckNameAvailability_checkNameAvailabilityAlreadyExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CheckNameAvailability(ctx, armdeviceupdate.CheckNameAvailabilityRequest{
		Name: to.Ptr("contoso"),
		Type: to.Ptr("Microsoft.DeviceUpdate/accounts"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResponse = armdeviceupdate.CheckNameAvailabilityResponse{
	// 	Message: to.Ptr("Resource name already exists"),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armdeviceupdate.CheckNameAvailabilityReasonAlreadyExists),
	// }
}
