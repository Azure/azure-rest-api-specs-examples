package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/workspaceManagerAssignments/CreateJob.json
func ExampleWorkspaceManagerAssignmentJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceManagerAssignmentJobsClient().Create(ctx, "myRg", "myWorkspace", "47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.WorkspaceManagerAssignmentJobsClientCreateResponse{
	// 	Job: armsecurityinsights.Job{
	// 		Name: to.Ptr("cfbe1338-8276-4d5d-8b96-931117f9fa0e"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/workspaceManagerAssignments/jobs"),
	// 		Etag: to.Ptr("\"f20a2523-7817-47b5-a3b2-21539c00c788\""),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58/jobs/cfbe1338-8276-4d5d-8b96-931117f9fa0e"),
	// 		Properties: &armsecurityinsights.JobProperties{
	// 			ProvisioningState: to.Ptr(armsecurityinsights.ProvisioningStateInProgress),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-14T04:47:52.9614956Z"); return t}()),
	// 		},
	// 	},
	// }
}
