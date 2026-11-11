package armsovereign_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
)

// Generated from example definition: 2026-10-31-preview/Operations_CheckNameAvailability.json
func ExampleOperationsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsovereign.NewClientFactory("ecb23b6b-b38a-42c7-b8ea-01a4aad4483a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().CheckNameAvailability(ctx, armsovereign.CheckNameAvailabilityRequest{
		Name: to.Ptr("name1"),
		Type: to.Ptr("Microsoft.Sovereign/sovereignViews"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsovereign.OperationsClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResponse: armsovereign.CheckNameAvailabilityResponse{
	// 		NameAvailable: to.Ptr(true),
	// 	},
	// }
}
