package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControlSyncJob/createSourceControlSyncJob.json
func ExampleSourceControlSyncJobClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
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
