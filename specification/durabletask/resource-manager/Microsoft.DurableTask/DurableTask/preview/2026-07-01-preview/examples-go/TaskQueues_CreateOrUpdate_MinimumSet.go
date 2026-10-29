package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2026-07-01-preview/TaskQueues_CreateOrUpdate_MinimumSet.json
func ExampleTaskQueuesClient_BeginCreateOrUpdate_taskQueuesCreateOrUpdateMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("EE9BD735-67CE-4A90-89C4-439D3F6A4C93", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTaskQueuesClient().BeginCreateOrUpdate(ctx, "rgopenapi", "testscheduler", "testtaskhub", "QueueA", armdurabletask.TaskQueue{
		Properties: &armdurabletask.TaskQueueProperties{},
	}, nil)
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
	// res = armdurabletask.TaskQueuesClientCreateOrUpdateResponse{
	// 	TaskQueue: armdurabletask.TaskQueue{
	// 		ID: to.Ptr("/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgopenapi/providers/Microsoft.DurableTask/schedulers/testscheduler/taskHubs/testtaskhub/taskQueues/QueueA"),
	// 		Name: to.Ptr("QueueA"),
	// 		Type: to.Ptr("Microsoft.DurableTask/schedulers/taskHubs/taskQueues"),
	// 		Properties: &armdurabletask.TaskQueueProperties{
	// 			ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
