package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2026-07-01-preview/TaskQueues_ListByTaskHub.json
func ExampleTaskQueuesClient_NewListByTaskHubPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("EE9BD735-67CE-4A90-89C4-439D3F6A4C93", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTaskQueuesClient().NewListByTaskHubPager("rgopenapi", "testscheduler", "testtaskhub", nil)
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
		// page = armdurabletask.TaskQueuesClientListByTaskHubResponse{
		// 	TaskQueueListResult: armdurabletask.TaskQueueListResult{
		// 		Value: []*armdurabletask.TaskQueue{
		// 			{
		// 				ID: to.Ptr("/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgopenapi/providers/Microsoft.DurableTask/schedulers/testscheduler/taskHubs/testtaskhub/taskQueues/QueueA"),
		// 				Name: to.Ptr("QueueA"),
		// 				Type: to.Ptr("Microsoft.DurableTask/schedulers/taskHubs/taskQueues"),
		// 				Properties: &armdurabletask.TaskQueueProperties{
		// 					ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
		// 					Throughput: &armdurabletask.TaskQueueThroughput{
		// 						ActionsPerSecond: to.Ptr[int32](40),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgopenapi/providers/Microsoft.DurableTask/schedulers/testscheduler/taskHubs/testtaskhub/taskQueues/QueueB"),
		// 				Name: to.Ptr("QueueB"),
		// 				Type: to.Ptr("Microsoft.DurableTask/schedulers/taskHubs/taskQueues"),
		// 				Properties: &armdurabletask.TaskQueueProperties{
		// 					ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
		// 					Throughput: &armdurabletask.TaskQueueThroughput{
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgopenapi/providers/Microsoft.DurableTask/schedulers/testscheduler/taskHubs/testtaskhub/taskQueues?api-version=2026-07-01-preview&$skiptoken=nextPage"),
		// 	},
		// }
	}
}
