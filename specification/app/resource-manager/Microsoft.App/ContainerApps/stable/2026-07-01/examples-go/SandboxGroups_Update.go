package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/SandboxGroups_Update.json
func ExampleSandboxGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSandboxGroupsClient().BeginUpdate(ctx, "examplerg", "testgroup", armappcontainers.SandboxGroupPatch{
		Tags: map[string]*string{
			"environment": to.Ptr("test"),
		},
		Properties: &armappcontainers.SandboxGroupPatchProperties{
			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/exampleenv"),
		},
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
	// res = armappcontainers.SandboxGroupsClientUpdateResponse{
	// 	SandboxGroup: armappcontainers.SandboxGroup{
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/sandboxGroups/testgroup"),
	// 		Name: to.Ptr("testgroup"),
	// 		Type: to.Ptr("Microsoft.App/sandboxGroups"),
	// 		Location: to.Ptr("East US"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("test"),
	// 		},
	// 		Properties: &armappcontainers.SandboxGroupProperties{
	// 			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/exampleenv"),
	// 			DefaultDomain: to.Ptr("icycliff-79e92900.eastus.azurecontainerapps.io"),
	// 			ProvisioningState: to.Ptr(armappcontainers.SandboxGroupProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
