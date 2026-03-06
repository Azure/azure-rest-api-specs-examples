package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/LocationCheckNameAvailability_AlreadyExists.json
func ExampleLocationClient_CheckNameAvailability_locationCheckNameAvailabilityAlreadyExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationClient().CheckNameAvailability(ctx, "japaneast", armbatch.CheckNameAvailabilityParameters{
		Name: to.Ptr("existingaccountname"),
		Type: to.Ptr("Microsoft.Batch/batchAccounts"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbatch.LocationClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResult: &armbatch.CheckNameAvailabilityResult{
	// 		Message: to.Ptr("An account named 'existingaccountname' is already in use."),
	// 		NameAvailable: to.Ptr(false),
	// 		Reason: to.Ptr(armbatch.NameAvailabilityReasonAlreadyExists),
	// 	},
	// }
}
