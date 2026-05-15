package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/ExportRunHistoryGetByEnrollmentAccount.json
func ExampleExportsClient_GetExecutionHistory_exportRunHistoryGetByEnrollmentAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportsClient().GetExecutionHistory(ctx, "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456", "TestExport", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcostmanagement.ExportsClientGetExecutionHistoryResponse{
	// 	ExportExecutionListResult: armcostmanagement.ExportExecutionListResult{
	// 		Value: []*armcostmanagement.ExportRun{
	// 			{
	// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456/providers/Microsoft.CostManagement/exports/JohnDoeSchedule/Run/1e25d58a-a3b0-4916-9542-6e04a89bc100"),
	// 				Properties: &armcostmanagement.ExportRunProperties{
	// 					EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-30T00:00:00Z"); return t}()),
	// 					ExecutionType: to.Ptr(armcostmanagement.ExecutionTypeOnDemand),
	// 					FileName: to.Ptr("ScheduledTestsForJohnDoe/JohnDoeSchedule/20180729-20180804/JohnDoeSchedule_1e25d58a-a3b0-4916-9542-6e04a89bc100.csv"),
	// 					ManifestFile: to.Ptr("ScheduledTestsForJohnDoe/JohnDoeSchedule/20180729-20180804/manifestfile"),
	// 					ProcessingEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T07:52:28.0373318Z"); return t}()),
	// 					ProcessingStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T07:52:16.9123797Z"); return t}()),
	// 					RunSettings: &armcostmanagement.CommonExportProperties{
	// 						Format: to.Ptr(armcostmanagement.FormatTypeCSV),
	// 						CompressionMode: to.Ptr(armcostmanagement.CompressionModeTypeGzip),
	// 						DataOverwriteBehavior: to.Ptr(armcostmanagement.DataOverwriteBehaviorTypeOverwritePreviousReport),
	// 						Definition: &armcostmanagement.ExportDefinition{
	// 							Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
	// 							DataSet: &armcostmanagement.ExportDataset{
	// 								Configuration: &armcostmanagement.ExportDatasetConfiguration{
	// 									DataVersion: to.Ptr("2023-05-01"),
	// 								},
	// 								Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
	// 							},
	// 							TimePeriod: &armcostmanagement.ExportTimePeriod{
	// 								From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t}()),
	// 								To: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-30T00:00:00Z"); return t}()),
	// 							},
	// 							Timeframe: to.Ptr(armcostmanagement.TimeframeTypeCustom),
	// 						},
	// 						DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
	// 							Destination: &armcostmanagement.ExportDeliveryDestination{
	// 								Container: to.Ptr("exports"),
	// 								ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
	// 								RootFolderPath: to.Ptr("ScheduledTestsForJohnDoe"),
	// 							},
	// 						},
	// 						ExportDescription: to.Ptr("This is a test export."),
	// 						PartitionData: to.Ptr(true),
	// 					},
	// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t}()),
	// 					Status: to.Ptr(armcostmanagement.ExecutionStatusCompleted),
	// 					SubmittedBy: to.Ptr("john.doe@gmail.com"),
	// 					SubmittedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T07:52:15.6016681Z"); return t}()),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456/providers/Microsoft.CostManagement/exports/JohnDoeSchedule/Run/11ac6811-dca3-46ad-b326-4704cf0c58ef"),
	// 				Properties: &armcostmanagement.ExportRunProperties{
	// 					EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-30T00:00:00Z"); return t}()),
	// 					ExecutionType: to.Ptr(armcostmanagement.ExecutionTypeScheduled),
	// 					FileName: to.Ptr("ScheduledTestsForJohnDoe/JohnDoeSchedule/20180729-20180804/JohnDoeSchedule_11ac6811-dca3-46ad-b326-4704cf0c58ef.csv"),
	// 					ManifestFile: to.Ptr("ScheduledTestsForJohnDoe/JohnDoeSchedule/20180729-20180804/manifestfile"),
	// 					ProcessingEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T09:04:19.7223808Z"); return t}()),
	// 					ProcessingStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T09:03:58.5710244Z"); return t}()),
	// 					RunSettings: &armcostmanagement.CommonExportProperties{
	// 						Format: to.Ptr(armcostmanagement.FormatTypeCSV),
	// 						CompressionMode: to.Ptr(armcostmanagement.CompressionModeTypeGzip),
	// 						DataOverwriteBehavior: to.Ptr(armcostmanagement.DataOverwriteBehaviorTypeOverwritePreviousReport),
	// 						Definition: &armcostmanagement.ExportDefinition{
	// 							Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
	// 							DataSet: &armcostmanagement.ExportDataset{
	// 								Configuration: &armcostmanagement.ExportDatasetConfiguration{
	// 									DataVersion: to.Ptr("2023-05-01"),
	// 								},
	// 								Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
	// 							},
	// 							TimePeriod: &armcostmanagement.ExportTimePeriod{
	// 								From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t}()),
	// 								To: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-30T00:00:00Z"); return t}()),
	// 							},
	// 							Timeframe: to.Ptr(armcostmanagement.TimeframeTypeCustom),
	// 						},
	// 						DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
	// 							Destination: &armcostmanagement.ExportDeliveryDestination{
	// 								Container: to.Ptr("exports"),
	// 								ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
	// 								RootFolderPath: to.Ptr("ScheduledTestsForJohnDoe"),
	// 							},
	// 						},
	// 						ExportDescription: to.Ptr("This is a test export."),
	// 						PartitionData: to.Ptr(true),
	// 					},
	// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t}()),
	// 					Status: to.Ptr(armcostmanagement.ExecutionStatusCompleted),
	// 					SubmittedBy: to.Ptr("System"),
	// 					SubmittedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T09:03:58.5710244Z"); return t}()),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
