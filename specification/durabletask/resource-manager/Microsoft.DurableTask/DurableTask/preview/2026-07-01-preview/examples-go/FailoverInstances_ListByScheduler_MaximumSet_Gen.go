package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2026-07-01-preview/FailoverInstances_ListByScheduler_MaximumSet_Gen.json
func ExampleFailoverInstancesClient_NewListBySchedulerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("194D3C1E-462F-4738-9025-092A628C06EB", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFailoverInstancesClient().NewListBySchedulerPager("rgdurabletask", "testscheduler", nil)
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
		// page = armdurabletask.FailoverInstancesClientListBySchedulerResponse{
		// 	FailoverInstanceListResult: armdurabletask.FailoverInstanceListResult{
		// 		Value: []*armdurabletask.FailoverInstance{
		// 			{
		// 				Properties: &armdurabletask.FailoverInstanceProperties{
		// 					ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
		// 					ReplicationRole: to.Ptr(armdurabletask.ReplicationRolePrimary),
		// 					FailoverState: to.Ptr(armdurabletask.FailoverStateSucceeded),
		// 					PartnerSchedulerResourceID: to.Ptr("/subscriptions/194D3C1E-462F-4738-9025-092A628C06EB/resourceGroups/rgdurabletask-secondary/providers/Microsoft.DurableTask/schedulers/testscheduler-secondary"),
		// 					PartnerSchedulerLocation: to.Ptr("westus2"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/194D3C1E-462F-4738-9025-092A628C06EB/resourceGroups/rgdurabletask/providers/Microsoft.DurableTask/schedulers/testscheduler/failoverInstances/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.DurableTask/schedulers/failoverInstances"),
		// 				SystemData: &armdurabletask.SystemData{
		// 					CreatedBy: to.Ptr("zshkmc"),
		// 					CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2025, time.March, 31, 23, 34, 9, 612000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("ivqrae"),
		// 					LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.March, 31, 23, 34, 9, 612000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/ai"),
		// 	},
		// }
	}
}
