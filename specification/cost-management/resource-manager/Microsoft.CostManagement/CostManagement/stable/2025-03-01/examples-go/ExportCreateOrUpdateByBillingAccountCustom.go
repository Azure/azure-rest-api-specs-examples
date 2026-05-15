package armcostmanagement_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/ExportCreateOrUpdateByBillingAccountCustom.json
func ExampleExportsClient_CreateOrUpdate_exportCreateOrUpdateByBillingAccountCustom() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportsClient().CreateOrUpdate(ctx, "providers/Microsoft.Billing/billingAccounts/123456", "TestExport", armcostmanagement.Export{
		Identity: &armcostmanagement.SystemAssignedServiceIdentity{
			Type: to.Ptr(armcostmanagement.SystemAssignedServiceIdentityTypeSystemAssigned),
		},
		Location: to.Ptr("centralus"),
		Properties: &armcostmanagement.ExportProperties{
			Format:                to.Ptr(armcostmanagement.FormatTypeCSV),
			CompressionMode:       to.Ptr(armcostmanagement.CompressionModeTypeGzip),
			DataOverwriteBehavior: to.Ptr(armcostmanagement.DataOverwriteBehaviorTypeOverwritePreviousReport),
			Definition: &armcostmanagement.ExportDefinition{
				Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
				DataSet: &armcostmanagement.ExportDataset{
					Configuration: &armcostmanagement.ExportDatasetConfiguration{
						DataVersion: to.Ptr("2023-05-01"),
					},
					Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
				},
				TimePeriod: &armcostmanagement.ExportTimePeriod{
					From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-03T00:00:00.000Z"); return t }()),
					To:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-03T00:00:00.000Z"); return t }()),
				},
				Timeframe: to.Ptr(armcostmanagement.TimeframeTypeCustom),
			},
			DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
				Destination: &armcostmanagement.ExportDeliveryDestination{
					Type:           to.Ptr(armcostmanagement.DestinationTypeAzureBlob),
					Container:      to.Ptr("exports"),
					ResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
					RootFolderPath: to.Ptr("ad-hoc"),
				},
			},
			ExportDescription: to.Ptr("This is a test export."),
			PartitionData:     to.Ptr(true),
			Schedule: &armcostmanagement.ExportSchedule{
				Status: to.Ptr(armcostmanagement.StatusTypeInactive),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcostmanagement.ExportsClientCreateOrUpdateResponse{
	// 	Export: armcostmanagement.Export{
	// 		Name: to.Ptr("TestExport"),
	// 		Type: to.Ptr("Microsoft.CostManagement/exports"),
	// 		ETag: to.Ptr(azcore.ETag("\"00000000-0000-0000-0000-000000000000\"")),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/providers/Microsoft.CostManagement/exports/TestExport"),
	// 		Identity: &armcostmanagement.SystemAssignedServiceIdentity{
	// 			Type: to.Ptr(armcostmanagement.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Location: to.Ptr("centralus"),
	// 		Properties: &armcostmanagement.ExportProperties{
	// 			Format: to.Ptr(armcostmanagement.FormatTypeCSV),
	// 			CompressionMode: to.Ptr(armcostmanagement.CompressionModeTypeGzip),
	// 			DataOverwriteBehavior: to.Ptr(armcostmanagement.DataOverwriteBehaviorTypeOverwritePreviousReport),
	// 			Definition: &armcostmanagement.ExportDefinition{
	// 				Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
	// 				DataSet: &armcostmanagement.ExportDataset{
	// 					Configuration: &armcostmanagement.ExportDatasetConfiguration{
	// 						DataVersion: to.Ptr("2023-05-01"),
	// 					},
	// 					Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
	// 				},
	// 				TimePeriod: &armcostmanagement.ExportTimePeriod{
	// 					From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-03T00:00:00Z"); return t}()),
	// 					To: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-03T00:00:00Z"); return t}()),
	// 				},
	// 				Timeframe: to.Ptr(armcostmanagement.TimeframeTypeCustom),
	// 			},
	// 			DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
	// 				Destination: &armcostmanagement.ExportDeliveryDestination{
	// 					Container: to.Ptr("exports"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
	// 					RootFolderPath: to.Ptr("ad-hoc"),
	// 				},
	// 			},
	// 			ExportDescription: to.Ptr("This is a test export."),
	// 			PartitionData: to.Ptr(true),
	// 			Schedule: &armcostmanagement.ExportSchedule{
	// 				Recurrence: to.Ptr(armcostmanagement.RecurrenceType("None")),
	// 				Status: to.Ptr(armcostmanagement.StatusTypeInactive),
	// 			},
	// 		},
	// 	},
	// }
}
