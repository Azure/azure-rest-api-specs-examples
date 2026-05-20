package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync/v2"
)

// Generated from example definition: 2022-09-01/StorageSyncServiceCheckNameAvailability_AlreadyExists.json
func ExampleServicesClient_CheckNameAvailability_storageSyncServiceCheckNameAvailabilityAlreadyExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("5c6bc8e1-1eaf-4192-94d8-58ce463ac86c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().CheckNameAvailability(ctx, "westus", armstoragesync.CheckNameAvailabilityParameters{
		Name: to.Ptr("newstoragesyncservicename"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragesync.ServicesClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResult: armstoragesync.CheckNameAvailabilityResult{
	// 		Message: to.Ptr("An account named 'newstoragesyncservicename' is already in use."),
	// 		NameAvailable: to.Ptr(false),
	// 		Reason: to.Ptr(armstoragesync.NameAvailabilityReasonAlreadyExists),
	// 	},
	// }
}
