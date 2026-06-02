package armmonitorworkspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitorworkspaces"
)

// Generated from example definition: 2025-10-03/Issue_List_MaximumSet_Gen.json
func ExampleIssueClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitorworkspaces.NewClientFactory("aceaa046-91f0-492a-96dc-45e10a9183dc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIssueClient().NewListPager("rg1", "myWorkspace", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armmonitorworkspaces.IssueClientListResponse{
		// 	IssueResourceListResult: armmonitorworkspaces.IssueResourceListResult{
		// 		Value: []*armmonitorworkspaces.IssueResource{
		// 			{
		// 				Properties: &armmonitorworkspaces.IssueProperties{
		// 					Title: to.Ptr("Alert fired on VM CPU"),
		// 					Status: to.Ptr(armmonitorworkspaces.StatusNew),
		// 					Severity: to.Ptr("Sev2"),
		// 					Investigations: []*armmonitorworkspaces.InvestigationMetadata{
		// 					},
		// 					ImpactTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-13T02:45:33"); return t}()),
		// 					InvestigationsCount: to.Ptr[int32](0),
		// 					Notifications: &armmonitorworkspaces.Notifications{
		// 						UpdateTypes: []armmonitorworkspaces.IssueNotificationTypeClassification{
		// 							&armmonitorworkspaces.IssueCreationNotificationType{
		// 								UpdateType: to.Ptr(armmonitorworkspaces.UpdateTypeIssueCreation),
		// 							},
		// 							&armmonitorworkspaces.OnChangeNotificationType{
		// 								UpdateType: to.Ptr(armmonitorworkspaces.UpdateTypeOnChange),
		// 							},
		// 							&armmonitorworkspaces.TimeBasedUpdatesNotificationType{
		// 								UpdateType: to.Ptr(armmonitorworkspaces.UpdateTypeTimeBased),
		// 								UpdateInterval: to.Ptr("PT1H"),
		// 							},
		// 						},
		// 						ActionGroupIDs: []*string{
		// 							to.Ptr("/subscriptions/aceaa046-91f0-492a-96dc-45e10a9183dc/resourceGroups/rg1/providers/Microsoft.Insights/actionGroups/myActionGroup"),
		// 						},
		// 						ExcludeDefaultActionGroups: to.Ptr(false),
		// 					},
		// 					ProvisioningState: to.Ptr(armmonitorworkspaces.ResourceProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/aceaa046-91f0-492a-96dc-45e10a9183dc/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.AlertsManagement/issues/3f29e1b2b05f8371595dc761fed8e8b3"),
		// 				Name: to.Ptr("3f29e1b2b05f8371595dc761fed8e8b3"),
		// 				SystemData: &armmonitorworkspaces.SystemData{
		// 					CreatedBy: to.Ptr("171a811c-2a3a-4e6c-b742-f78f5f6ca51c"),
		// 					CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByType("Manual")),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-13T03:11:07"); return t}()),
		// 					LastModifiedBy: to.Ptr("171a811c-2a3a-4e6c-b742-f78f5f6ca51c"),
		// 					LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByType("Manual")),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-13T03:51:45"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
