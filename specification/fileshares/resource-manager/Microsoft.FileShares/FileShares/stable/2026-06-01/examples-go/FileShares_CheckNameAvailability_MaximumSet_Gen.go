package armfileshares_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fileshares/armfileshares"
)

// Generated from example definition: 2026-06-01/FileShares_CheckNameAvailability_MaximumSet_Gen.json
func ExampleClient_CheckNameAvailability_fileSharesCheckNameAvailabilityMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfileshares.NewClientFactory("0681745E-3F9F-4966-80E6-69624A3B29F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CheckNameAvailability(ctx, "westus", armfileshares.CheckNameAvailabilityRequest{
		Name: to.Ptr("fvykqbgmd"),
		Type: to.Ptr("Microsoft.FileShares/fileShares"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armfileshares.ClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResponse: armfileshares.CheckNameAvailabilityResponse{
	// 		NameAvailable: to.Ptr(true),
	// 		Reason: to.Ptr(armfileshares.CheckNameAvailabilityReasonInvalid),
	// 		Message: to.Ptr("ysdbdamxulhxnqihecifxjepymgw"),
	// 	},
	// }
}
