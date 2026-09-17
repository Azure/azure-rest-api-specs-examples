package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/SandboxGroups_CreateOrUpdate.json
func ExampleSandboxGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSandboxGroupsClient().BeginCreateOrUpdate(ctx, "examplerg", "testgroup", armappcontainers.SandboxGroup{
		Location: to.Ptr("East US"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.SandboxGroupsClientCreateOrUpdateResponse{
	// 	SandboxGroup: armappcontainers.SandboxGroup{
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/sandboxGroups/testgroup"),
	// 		Name: to.Ptr("testgroup"),
	// 		Type: to.Ptr("Microsoft.App/sandboxGroups"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armappcontainers.SandboxGroupProperties{
	// 			ProvisioningState: to.Ptr(armappcontainers.SandboxGroupProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
