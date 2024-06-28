package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/storageTaskAssignmentsList/ListStorageTaskAssignmentsForAccount.json
func ExampleTaskAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTaskAssignmentsClient().NewListPager("res4228", "sto4445", &armstorage.TaskAssignmentsClientListOptions{Maxpagesize: nil})
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
		// page.TaskAssignmentsList = armstorage.TaskAssignmentsList{
		// 	Value: []*armstorage.TaskAssignment{
		// 		{
		// 			Name: to.Ptr("myassignment1"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/storageTaskAssignments"),
		// 			ID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445/storageTaskAssignments/myassignment1"),
		// 			Properties: &armstorage.TaskAssignmentProperties{
		// 				Description: to.Ptr("My Storage task assignment #1"),
		// 				Enabled: to.Ptr(true),
		// 				ExecutionContext: &armstorage.TaskAssignmentExecutionContext{
		// 					Target: &armstorage.ExecutionTarget{
		// 						ExcludePrefix: []*string{
		// 						},
		// 						Prefix: []*string{
		// 							to.Ptr("prefix1"),
		// 							to.Ptr("prefix2")},
		// 						},
		// 						Trigger: &armstorage.ExecutionTrigger{
		// 							Type: to.Ptr(armstorage.TriggerTypeRunOnce),
		// 							Parameters: &armstorage.TriggerParameters{
		// 								StartOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-15T21:52:47.814Z"); return t}()),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 					Report: &armstorage.TaskAssignmentReport{
		// 						Prefix: to.Ptr("container1"),
		// 					},
		// 					TaskID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myassignment2"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/storageTaskAssignments"),
		// 				ID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445/storageTaskAssignments/myassignment2"),
		// 				Properties: &armstorage.TaskAssignmentProperties{
		// 					Description: to.Ptr("My Storage task assignment #2"),
		// 					Enabled: to.Ptr(true),
		// 					ExecutionContext: &armstorage.TaskAssignmentExecutionContext{
		// 						Target: &armstorage.ExecutionTarget{
		// 							ExcludePrefix: []*string{
		// 							},
		// 							Prefix: []*string{
		// 								to.Ptr("prefix3"),
		// 								to.Ptr("prefix4")},
		// 							},
		// 							Trigger: &armstorage.ExecutionTrigger{
		// 								Type: to.Ptr(armstorage.TriggerTypeRunOnce),
		// 								Parameters: &armstorage.TriggerParameters{
		// 									StartOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-15T21:52:47.814Z"); return t}()),
		// 								},
		// 							},
		// 						},
		// 						ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 						Report: &armstorage.TaskAssignmentReport{
		// 							Prefix: to.Ptr("container2"),
		// 						},
		// 						TaskID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask2"),
		// 					},
		// 			}},
		// 		}
	}
}
