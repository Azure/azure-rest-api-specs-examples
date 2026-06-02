package armstorage_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-04-01/storageTaskAssignmentsCrud/PatchStorageTaskAssignment.json
func ExampleTaskAssignmentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("1f31ba14-ce16-4281-b9b4-3e78da6e1616", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTaskAssignmentsClient().BeginUpdate(ctx, "res4228", "sto4445", "myassignment1", armstorage.TaskAssignmentUpdateParameters{
		Properties: &armstorage.TaskAssignmentUpdateProperties{
			Description: to.Ptr("My Storage task assignment"),
			Enabled:     to.Ptr(true),
			ExecutionContext: &armstorage.TaskAssignmentUpdateExecutionContext{
				Target: &armstorage.ExecutionTarget{
					ExcludePrefix: []*string{},
					Prefix: []*string{
						to.Ptr("prefix1"),
						to.Ptr("prefix2"),
					},
				},
				Trigger: &armstorage.ExecutionTriggerUpdate{
					Type: to.Ptr(armstorage.TriggerTypeRunOnce),
					Parameters: &armstorage.TriggerParametersUpdate{
						StartOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-15T21:52:47.8145095Z"); return t }()),
					},
				},
			},
			Report: &armstorage.TaskAssignmentUpdateReport{
				Prefix: to.Ptr("container1"),
			},
		},
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
	// res = armstorage.TaskAssignmentsClientUpdateResponse{
	// 	TaskAssignment: armstorage.TaskAssignment{
	// 		Name: to.Ptr("myassignment1"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/storageTaskAssignments"),
	// 		ID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445/storageTaskAssignments/myassignment1"),
	// 		Properties: &armstorage.TaskAssignmentProperties{
	// 			Description: to.Ptr("My Storage task assignment"),
	// 			Enabled: to.Ptr(true),
	// 			ExecutionContext: &armstorage.TaskAssignmentExecutionContext{
	// 				Target: &armstorage.ExecutionTarget{
	// 					ExcludePrefix: []*string{
	// 					},
	// 					Prefix: []*string{
	// 						to.Ptr("prefix1"),
	// 						to.Ptr("prefix2"),
	// 					},
	// 				},
	// 				Trigger: &armstorage.ExecutionTrigger{
	// 					Type: to.Ptr(armstorage.TriggerTypeRunOnce),
	// 					Parameters: &armstorage.TriggerParameters{
	// 						StartOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-15T21:52:47.8145095Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armstorage.StorageTaskAssignmentProvisioningStateSucceeded),
	// 			Report: &armstorage.TaskAssignmentReport{
	// 				Prefix: to.Ptr("container1"),
	// 			},
	// 			TaskID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1"),
	// 		},
	// 	},
	// }
}
