package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/listHybridRunbookWorker.json
func ExampleHybridRunbookWorkersClient_NewListByHybridRunbookWorkerGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHybridRunbookWorkersClient().NewListByHybridRunbookWorkerGroupPager("rg", "testaccount", "TestHybridGroup", nil)
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
		// page = armautomation.HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse{
		// 	HybridRunbookWorkersListResult: armautomation.HybridRunbookWorkersListResult{
		// 		Value: []*armautomation.HybridRunbookWorker{
		// 			{
		// 				Name: to.Ptr("c010ad12-ef14-4a2a-aa9e-ef22c4745ddd"),
		// 				Type: to.Ptr("Microsoft.Automation/AutomationAccounts/HybridRunbookWorkerGroups/HybridRunbookWorkers"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/testaccount/hybridRunbookWorkerGroups/TestHybridGroup/hybridRunbookWorkers/c010ad12-ef14-4a2a-aa9e-ef22c4745ddd"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.HybridRunbookWorkerProperties{
		// 					IP: to.Ptr("10.0.0.0"),
		// 					LastSeenDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					RegisteredDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					VMResourceID: to.Ptr("/subscriptions/vmsubid/resourceGroups/vmrg/providers/Microsoft.Compute/virtualMachines/vmname"),
		// 					WorkerName: to.Ptr("vmname"),
		// 					WorkerType: to.Ptr(armautomation.WorkerTypeHybridV2),
		// 				},
		// 				SystemData: &armautomation.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					CreatedBy: to.Ptr("foo@contoso.com"),
		// 					CreatedByType: to.Ptr(armautomation.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armautomation.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("e18fe971-75b1-4351-987f-6fe3604bc721"),
		// 				Type: to.Ptr("Microsoft.Automation/AutomationAccounts/HybridRunbookWorkerGroups/HybridRunbookWorkers"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/testaccount/hybridRunbookWorkerGroups/TestHybridGroup/hybridRunbookWorkers/e18fe971-75b1-4351-987f-6fe3604bc721"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.HybridRunbookWorkerProperties{
		// 					IP: to.Ptr("10.0.0.1"),
		// 					LastSeenDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-25T16:30:55+00:00"); return t}()),
		// 					RegisteredDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-25T16:30:55+00:00"); return t}()),
		// 					VMResourceID: to.Ptr(""),
		// 					WorkerName: to.Ptr("myworker"),
		// 					WorkerType: to.Ptr(armautomation.WorkerTypeHybridV1),
		// 				},
		// 				SystemData: &armautomation.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					CreatedBy: to.Ptr("foo@contoso.com"),
		// 					CreatedByType: to.Ptr(armautomation.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armautomation.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
