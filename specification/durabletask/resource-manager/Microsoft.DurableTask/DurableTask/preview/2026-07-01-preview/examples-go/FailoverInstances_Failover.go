package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2026-07-01-preview/FailoverInstances_Failover.json
func ExampleFailoverInstancesClient_BeginFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("194D3C1E-462F-4738-9025-092A628C06EB", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFailoverInstancesClient().BeginFailover(ctx, "rgdurabletask", "testscheduler", nil)
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
	// res = armdurabletask.FailoverInstancesClientFailoverResponse{
	// 	FailoverInstance: armdurabletask.FailoverInstance{
	// 		Properties: &armdurabletask.FailoverInstanceProperties{
	// 			ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
	// 			ReplicationRole: to.Ptr(armdurabletask.ReplicationRolePrimary),
	// 			FailoverState: to.Ptr(armdurabletask.FailoverStateSucceeded),
	// 			PartnerSchedulerResourceID: to.Ptr("/subscriptions/194D3C1E-462F-4738-9025-092A628C06EB/resourceGroups/rgdurabletask-secondary/providers/Microsoft.DurableTask/schedulers/testscheduler-secondary"),
	// 			PartnerSchedulerLocation: to.Ptr("westus2"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/194D3C1E-462F-4738-9025-092A628C06EB/resourceGroups/rgdurabletask/providers/Microsoft.DurableTask/schedulers/testscheduler/failoverInstances/default"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.DurableTask/schedulers/failoverInstances"),
	// 		SystemData: &armdurabletask.SystemData{
	// 			CreatedBy: to.Ptr("zshkmc"),
	// 			CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2025, time.March, 31, 23, 34, 9, 612000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("ivqrae"),
	// 			LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.July, 13, 18, 30, 43, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
