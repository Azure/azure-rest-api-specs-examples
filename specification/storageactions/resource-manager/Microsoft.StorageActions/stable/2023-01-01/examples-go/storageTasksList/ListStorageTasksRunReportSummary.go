package armstorageactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageactions/armstorageactions"
)

// Generated from example definition: 2023-01-01/storageTasksList/ListStorageTasksRunReportSummary.json
func ExampleStorageTasksReportClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageactions.NewClientFactory("1f31ba14-ce16-4281-b9b4-3e78da6e1616", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageTasksReportClient().NewListPager("rgroup1", "mytask1", nil)
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
		// page = armstorageactions.StorageTasksReportClientListResponse{
		// 	StorageTaskReportSummary: armstorageactions.StorageTaskReportSummary{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example-rg/providers/Microsoft.StorageActions/storageTasks/example-task/reports?api-version=2023-01-01&monitor=true&$skiptoken=fake-continuation-token-12345"),
		// 		Value: []*armstorageactions.StorageTaskReportInstance{
		// 			{
		// 				Name: to.Ptr("instance1"),
		// 				Type: to.Ptr("Microsoft.StorageActions/storageTasks/reports"),
		// 				ID: to.Ptr("/subscriptions/9d6baee9-681f-44d8-99f7-d06d180af6bb/resourcegroups/rgroup1/providers/Microsoft.StorageActions/storageTasks/myTaskName/reports/instance1"),
		// 				Properties: &armstorageactions.StorageTaskReportProperties{
		// 					FinishTime: to.Ptr("2023-06-23T00:40:10.2931264Z"),
		// 					ObjectFailedCount: to.Ptr("0"),
		// 					ObjectsOperatedOnCount: to.Ptr("150"),
		// 					ObjectsSucceededCount: to.Ptr("150"),
		// 					ObjectsTargetedCount: to.Ptr("150"),
		// 					RunResult: to.Ptr(armstorageactions.RunResultSucceeded),
		// 					RunStatusEnum: to.Ptr(armstorageactions.RunStatusEnumFinished),
		// 					RunStatusError: to.Ptr("0"),
		// 					StartTime: to.Ptr("2023-06-23T00:30:43.226744Z"),
		// 					StorageAccountID: to.Ptr("/subscriptions/9d6baee9-681f-44d8-99f7-d06d180af6bb/resourcegroups/rgroup1/providers/Microsoft.Storage/storageAccounts/acc123"),
		// 					SummaryReportPath: to.Ptr("https://acc123.blob.core.windows.net/result-container/{folderpath}/SummaryReport.json"),
		// 					TaskAssignmentID: to.Ptr("/subscriptions/9d6baee9-681f-44d8-99f7-d06d180af6bb/resourcegroups/rgroup1/providers/Microsoft.Storage/storageAccounts/acc123/storageTaskAssignments/assign1"),
		// 					TaskID: to.Ptr("/subscriptions/9d6baee9-681f-44d8-99f7-d06d180af6bb/resourcegroups/rgroup1/providers/Microsoft.StorageActions/storageTasks/mytask1"),
		// 					TaskVersion: to.Ptr("1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("instance2"),
		// 				Type: to.Ptr("Microsoft.StorageActions/storageTasks/reports"),
		// 				ID: to.Ptr("/subscriptions/9d6baee9-681f-44d8-99f7-d06d180af6bb/resourcegroups/rgroup1/providers/Microsoft.StorageActions/storageTasks/myTaskName/reports/instance2"),
		// 				Properties: &armstorageactions.StorageTaskReportProperties{
		// 					FinishTime: to.Ptr("2023-06-23T00:40:10.2931264Z"),
		// 					ObjectFailedCount: to.Ptr("0"),
		// 					ObjectsOperatedOnCount: to.Ptr("150"),
		// 					ObjectsSucceededCount: to.Ptr("150"),
		// 					ObjectsTargetedCount: to.Ptr("150"),
		// 					RunResult: to.Ptr(armstorageactions.RunResultSucceeded),
		// 					RunStatusEnum: to.Ptr(armstorageactions.RunStatusEnumFinished),
		// 					RunStatusError: to.Ptr("0"),
		// 					StartTime: to.Ptr("2023-06-23T00:30:43.226744Z"),
		// 					StorageAccountID: to.Ptr("/subscriptions/9d6baee9-681f-44d8-99f7-d06d180af6bb/resourcegroups/rgroup1/providers/Microsoft.Storage/storageAccounts/acc123"),
		// 					SummaryReportPath: to.Ptr("https://acc123.blob.core.windows.net/result-container/{folderpath}/SummaryReport.json"),
		// 					TaskAssignmentID: to.Ptr("/subscriptions/9d6baee9-681f-44d8-99f7-d06d180af6bb/resourcegroups/rgroup1/providers/Microsoft.Storage/storageAccounts/acc123/storageTaskAssignments/assign1"),
		// 					TaskID: to.Ptr("/subscriptions/9d6baee9-681f-44d8-99f7-d06d180af6bb/resourcegroups/rgroup1/providers/Microsoft.StorageActions/storageTasks/mytask1"),
		// 					TaskVersion: to.Ptr("1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
