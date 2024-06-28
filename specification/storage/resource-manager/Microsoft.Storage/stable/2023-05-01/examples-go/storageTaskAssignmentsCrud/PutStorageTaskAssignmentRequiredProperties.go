package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/storageTaskAssignmentsCrud/PutStorageTaskAssignmentRequiredProperties.json
func ExampleTaskAssignmentsClient_BeginCreate_putStorageTaskAssignmentRequiredProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTaskAssignmentsClient().BeginCreate(ctx, "res4228", "sto4445", "myassignment1", armstorage.TaskAssignment{
		Properties: &armstorage.TaskAssignmentProperties{
			Description: to.Ptr("My Storage task assignment"),
			Enabled:     to.Ptr(true),
			ExecutionContext: &armstorage.TaskAssignmentExecutionContext{
				Trigger: &armstorage.ExecutionTrigger{
					Type: to.Ptr(armstorage.TriggerTypeRunOnce),
					Parameters: &armstorage.TriggerParameters{
						StartOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-15T21:52:47.814Z"); return t }()),
					},
				},
			},
			Report: &armstorage.TaskAssignmentReport{
				Prefix: to.Ptr("container1"),
			},
			TaskID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TaskAssignment = armstorage.TaskAssignment{
	// 	Name: to.Ptr("myassignment1"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/storageTaskAssignments"),
	// 	ID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445/storageTaskAssignments/myassignment1"),
	// 	Properties: &armstorage.TaskAssignmentProperties{
	// 		Description: to.Ptr("My Storage task assignment"),
	// 		Enabled: to.Ptr(true),
	// 		ExecutionContext: &armstorage.TaskAssignmentExecutionContext{
	// 			Trigger: &armstorage.ExecutionTrigger{
	// 				Type: to.Ptr(armstorage.TriggerTypeRunOnce),
	// 				Parameters: &armstorage.TriggerParameters{
	// 					StartOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-15T21:52:47.814Z"); return t}()),
	// 				},
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 		Report: &armstorage.TaskAssignmentReport{
	// 			Prefix: to.Ptr("container1"),
	// 		},
	// 		TaskID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1"),
	// 	},
	// }
}
