package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/sourceControlSyncJob/createSourceControlSyncJob.json
func ExampleSourceControlSyncJobClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSourceControlSyncJobClient().Create(ctx, "rg", "myAutomationAccount33", "MySourceControl", "ce6fe3e3-9db3-4096-a6b4-82bfb4c10a9a", armautomation.SourceControlSyncJobCreateParameters{
		Properties: &armautomation.SourceControlSyncJobCreateProperties{
			CommitID: to.Ptr("9de0980bfb45026a3d97a1b0522d98a9f604226e"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
