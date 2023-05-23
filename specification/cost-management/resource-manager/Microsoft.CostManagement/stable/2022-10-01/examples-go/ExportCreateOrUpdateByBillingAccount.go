package armcostmanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportCreateOrUpdateByBillingAccount.json
func ExampleExportsClient_CreateOrUpdate_exportCreateOrUpdateByBillingAccount() {
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
		Properties: &armcostmanagement.ExportProperties{
			Format: to.Ptr(armcostmanagement.FormatTypeCSV),
			Definition: &armcostmanagement.ExportDefinition{
				Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
				DataSet: &armcostmanagement.ExportDataset{
					Configuration: &armcostmanagement.ExportDatasetConfiguration{
						Columns: []*string{
							to.Ptr("Date"),
							to.Ptr("MeterId"),
							to.Ptr("ResourceId"),
							to.Ptr("ResourceLocation"),
							to.Ptr("Quantity")},
					},
					Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
				},
				Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
			},
			DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
				Destination: &armcostmanagement.ExportDeliveryDestination{
					Container:      to.Ptr("exports"),
					ResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
					RootFolderPath: to.Ptr("ad-hoc"),
				},
			},
			Schedule: &armcostmanagement.ExportSchedule{
				Recurrence: to.Ptr(armcostmanagement.RecurrenceTypeWeekly),
				RecurrencePeriod: &armcostmanagement.ExportRecurrencePeriod{
					From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t }()),
					To:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-31T00:00:00Z"); return t }()),
				},
				Status: to.Ptr(armcostmanagement.StatusTypeActive),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Export = armcostmanagement.Export{
	// 	Name: to.Ptr("TestExport"),
	// 	Type: to.Ptr("Microsoft.CostManagement/exports"),
	// 	ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/123456/providers/Microsoft.CostManagement/exports/TestExport"),
	// 	Properties: &armcostmanagement.ExportProperties{
	// 		Format: to.Ptr(armcostmanagement.FormatTypeCSV),
	// 		Definition: &armcostmanagement.ExportDefinition{
	// 			Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
	// 			DataSet: &armcostmanagement.ExportDataset{
	// 				Configuration: &armcostmanagement.ExportDatasetConfiguration{
	// 					Columns: []*string{
	// 						to.Ptr("Date"),
	// 						to.Ptr("MeterId"),
	// 						to.Ptr("ResourceId"),
	// 						to.Ptr("ResourceLocation"),
	// 						to.Ptr("Quantity")},
	// 					},
	// 					Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
	// 				},
	// 				Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
	// 			},
	// 			DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
	// 				Destination: &armcostmanagement.ExportDeliveryDestination{
	// 					Container: to.Ptr("exports"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
	// 					RootFolderPath: to.Ptr("ad-hoc"),
	// 				},
	// 			},
	// 			Schedule: &armcostmanagement.ExportSchedule{
	// 				Recurrence: to.Ptr(armcostmanagement.RecurrenceTypeWeekly),
	// 				RecurrencePeriod: &armcostmanagement.ExportRecurrencePeriod{
	// 					From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t}()),
	// 					To: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-31T00:00:00Z"); return t}()),
	// 				},
	// 				Status: to.Ptr(armcostmanagement.StatusTypeActive),
	// 			},
	// 		},
	// 	}
}
