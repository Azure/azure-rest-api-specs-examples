package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/Workspaces_CreateOrUpdate_CreateOnly.json
func ExampleWorkspacesClient_BeginCreateOrUpdate_createAWorkspaceOnlyWhenAbsentOrReturn412WhenItAlreadyExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginCreateOrUpdate(ctx, "exampleRG", "exampleWorkspace", armchaos.Workspace{
		Location: to.Ptr("westus2"),
		Properties: &armchaos.WorkspaceProperties{
			Scopes: []*string{
				to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG"),
			},
		},
	}, &armchaos.WorkspacesClientBeginCreateOrUpdateOptions{
		IfNoneMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
