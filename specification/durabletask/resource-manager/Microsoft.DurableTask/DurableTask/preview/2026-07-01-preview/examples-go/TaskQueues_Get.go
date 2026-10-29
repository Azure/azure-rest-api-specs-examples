package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2026-07-01-preview/TaskQueues_Get.json
func ExampleTaskQueuesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("EE9BD735-67CE-4A90-89C4-439D3F6A4C93", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTaskQueuesClient().Get(ctx, "rgopenapi", "testscheduler", "testtaskhub", "QueueA", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdurabletask.TaskQueuesClientGetResponse{
	// 	TaskQueue: armdurabletask.TaskQueue{
	// 		ID: to.Ptr("/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgopenapi/providers/Microsoft.DurableTask/schedulers/testscheduler/taskHubs/testtaskhub/taskQueues/QueueA"),
	// 		Name: to.Ptr("QueueA"),
	// 		Type: to.Ptr("Microsoft.DurableTask/schedulers/taskHubs/taskQueues"),
	// 		Properties: &armdurabletask.TaskQueueProperties{
	// 			ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
	// 			Throughput: &armdurabletask.TaskQueueThroughput{
	// 				ActionsPerSecond: to.Ptr[int32](40),
	// 			},
	// 		},
	// 		SystemData: &armdurabletask.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.July, 1, 12, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.July, 1, 12, 0, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
