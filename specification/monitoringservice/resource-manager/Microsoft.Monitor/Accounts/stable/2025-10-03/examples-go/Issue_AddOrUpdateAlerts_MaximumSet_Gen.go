package armmonitorworkspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitorworkspaces"
)

// Generated from example definition: 2025-10-03/Issue_AddOrUpdateAlerts_MaximumSet_Gen.json
func ExampleIssueClient_AddOrUpdateAlerts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitorworkspaces.NewClientFactory("aceaa046-91f0-492a-96dc-45e10a9183dc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIssueClient().AddOrUpdateAlerts(ctx, "rg1", "myWorkspace", "3f29e1b2b05f8371595dc761fed8e8b3", armmonitorworkspaces.RelatedAlerts{
		Value: []*armmonitorworkspaces.RelatedAlert{
			{
				ID:        to.Ptr("/subscriptions/aceaa046-91f0-492a-96dc-45e10a9183dc/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1/Microsoft.AlertsManagement/2ca4a194-dd28-4d38-92ca-df48f6a3efc3"),
				Relevance: to.Ptr(armmonitorworkspaces.RelevanceRelevant),
				Origin: &armmonitorworkspaces.Origin{
					AddedBy:     to.Ptr("171a811c-2a3a-4e6c-b742-f78f5f6ca51c"),
					AddedByType: to.Ptr(armmonitorworkspaces.AddedByTypeManual),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitorworkspaces.IssueClientAddOrUpdateAlertsResponse{
	// 	RelatedAlerts: armmonitorworkspaces.RelatedAlerts{
	// 		Value: []*armmonitorworkspaces.RelatedAlert{
	// 			{
	// 				ID: to.Ptr("/subscriptions/aceaa046-91f0-492a-96dc-45e10a9183dc/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1/Microsoft.AlertsManagement/2ca4a194-dd28-4d38-92ca-df48f6a3efc3"),
	// 				Relevance: to.Ptr(armmonitorworkspaces.RelevanceRelevant),
	// 				Origin: &armmonitorworkspaces.Origin{
	// 					AddedBy: to.Ptr("171a811c-2a3a-4e6c-b742-f78f5f6ca51c"),
	// 					AddedByType: to.Ptr(armmonitorworkspaces.AddedByTypeManual),
	// 				},
	// 				AddedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-13T03:11:07"); return t}()),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-13T03:51:45"); return t}()),
	// 			},
	// 		},
	// 	},
	// }
}
