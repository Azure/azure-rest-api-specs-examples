package armplaywright_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwright/armplaywright"
)

// Generated from example definition: 2025-09-01/PlaywrightWorkspaces_CheckNameAvailability.json
func ExampleWorkspacesClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywright.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().CheckNameAvailability(ctx, armplaywright.CheckNameAvailabilityRequest{
		Name: to.Ptr("dummyName"),
		Type: to.Ptr("Microsoft.LoadTestService/PlaywrightWorkspaces"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armplaywright.WorkspacesClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResponse: &armplaywright.CheckNameAvailabilityResponse{
	// 		NameAvailable: to.Ptr(true),
	// 		Message: to.Ptr("Test message."),
	// 	},
	// }
}
