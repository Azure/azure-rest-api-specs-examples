package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/ExportsGetBySubscription.json
func ExampleExportsClient_List_exportsGetBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportsClient().List(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcostmanagement.ExportsClientListResponse{
	// 	ExportListResult: armcostmanagement.ExportListResult{
	// 		Value: []*armcostmanagement.Export{
	// 			{
	// 				Name: to.Ptr("TestExport1"),
	// 				Type: to.Ptr("Microsoft.CostManagement/exports"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/exports/TestExport1"),
	// 				Identity: &armcostmanagement.SystemAssignedServiceIdentity{
	// 					Type: to.Ptr(armcostmanagement.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 				Location: to.Ptr("centralus"),
	// 				Properties: &armcostmanagement.ExportProperties{
	// 					Format: to.Ptr(armcostmanagement.FormatTypeCSV),
	// 					CompressionMode: to.Ptr(armcostmanagement.CompressionModeTypeGzip),
	// 					DataOverwriteBehavior: to.Ptr(armcostmanagement.DataOverwriteBehaviorTypeOverwritePreviousReport),
	// 					Definition: &armcostmanagement.ExportDefinition{
	// 						Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
	// 						DataSet: &armcostmanagement.ExportDataset{
	// 							Configuration: &armcostmanagement.ExportDatasetConfiguration{
	// 								DataVersion: to.Ptr("2023-05-01"),
	// 							},
	// 							Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
	// 						},
	// 						TimePeriod: &armcostmanagement.ExportTimePeriod{
	// 							From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t}()),
	// 							To: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-30T00:00:00Z"); return t}()),
	// 						},
	// 						Timeframe: to.Ptr(armcostmanagement.TimeframeTypeCustom),
	// 					},
	// 					DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
	// 						Destination: &armcostmanagement.ExportDeliveryDestination{
	// 							Container: to.Ptr("exports"),
	// 							ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
	// 							RootFolderPath: to.Ptr("ad-hoc"),
	// 						},
	// 					},
	// 					ExportDescription: to.Ptr("This is a test export."),
	// 					PartitionData: to.Ptr(true),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("TestExport2"),
	// 				Type: to.Ptr("Microsoft.CostManagement/exports"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/exports/TestExport2"),
	// 				Identity: &armcostmanagement.SystemAssignedServiceIdentity{
	// 					Type: to.Ptr(armcostmanagement.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 				Location: to.Ptr("centralus"),
	// 				Properties: &armcostmanagement.ExportProperties{
	// 					Format: to.Ptr(armcostmanagement.FormatTypeCSV),
	// 					CompressionMode: to.Ptr(armcostmanagement.CompressionModeTypeGzip),
	// 					DataOverwriteBehavior: to.Ptr(armcostmanagement.DataOverwriteBehaviorTypeOverwritePreviousReport),
	// 					Definition: &armcostmanagement.ExportDefinition{
	// 						Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
	// 						DataSet: &armcostmanagement.ExportDataset{
	// 							Configuration: &armcostmanagement.ExportDatasetConfiguration{
	// 								DataVersion: to.Ptr("2023-05-01"),
	// 							},
	// 							Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
	// 						},
	// 						TimePeriod: &armcostmanagement.ExportTimePeriod{
	// 							From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t}()),
	// 							To: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-30T00:00:00Z"); return t}()),
	// 						},
	// 						Timeframe: to.Ptr(armcostmanagement.TimeframeTypeCustom),
	// 					},
	// 					DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
	// 						Destination: &armcostmanagement.ExportDeliveryDestination{
	// 							Container: to.Ptr("exports"),
	// 							ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
	// 							RootFolderPath: to.Ptr("ad-hoc"),
	// 						},
	// 					},
	// 					ExportDescription: to.Ptr("This is a test export."),
	// 					PartitionData: to.Ptr(true),
	// 					Schedule: &armcostmanagement.ExportSchedule{
	// 						Recurrence: to.Ptr(armcostmanagement.RecurrenceTypeWeekly),
	// 						RecurrencePeriod: &armcostmanagement.ExportRecurrencePeriod{
	// 							From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t}()),
	// 							To: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-30T00:00:00Z"); return t}()),
	// 						},
	// 						Status: to.Ptr(armcostmanagement.StatusTypeActive),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
