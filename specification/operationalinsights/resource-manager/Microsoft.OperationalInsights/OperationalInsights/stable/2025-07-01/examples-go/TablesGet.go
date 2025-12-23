package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f0a5127d4e8b1ea6007b0bf9570904d9df860b97/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/TablesGet.json
func ExampleTablesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTablesClient().Get(ctx, "oiautorest6685", "oiautorest6685", "table1_SRCH", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Table = armoperationalinsights.Table{
	// 	Name: to.Ptr("table1_SRCH"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/oiautorest6685/providers/Microsoft.OperationalInsights/workspaces/oiautorest6685/tables/table1_SRCH"),
	// 	Properties: &armoperationalinsights.TableProperties{
	// 		Schema: &armoperationalinsights.Schema{
	// 			Name: to.Ptr("table1_SRCH"),
	// 			Columns: []*armoperationalinsights.Column{
	// 			},
	// 			Solutions: []*string{
	// 				to.Ptr("LogManagement")},
	// 				StandardColumns: []*armoperationalinsights.Column{
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumGUID),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SourceSystem"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("MG"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumGUID),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("ManagementGroupName"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TimeGenerated"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumDateTime),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Computer"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("RawData"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 				}},
	// 				TableSubType: to.Ptr(armoperationalinsights.TableSubTypeEnumDataCollectionRuleBased),
	// 				TableType: to.Ptr(armoperationalinsights.TableTypeEnumSearchResults),
	// 			},
	// 			ArchiveRetentionInDays: to.Ptr[int32](0),
	// 			Plan: to.Ptr(armoperationalinsights.TablePlanEnumAnalytics),
	// 			ProvisioningState: to.Ptr(armoperationalinsights.ProvisioningStateEnumSucceeded),
	// 			ResultStatistics: &armoperationalinsights.ResultStatistics{
	// 				IngestedRecords: to.Ptr[int32](5),
	// 				Progress: to.Ptr[float32](75),
	// 				ScannedGb: to.Ptr[float32](23.5),
	// 			},
	// 			RetentionInDays: to.Ptr[int32](50),
	// 			RetentionInDaysAsDefault: to.Ptr(false),
	// 			SearchResults: &armoperationalinsights.SearchResults{
	// 				Description: to.Ptr("This table was created using a Search Job with the following query: 'Heartbeat | where SourceSystem != '' | project SourceSystem'."),
	// 				EndSearchTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T04:05:06.000Z"); return t}()),
	// 				Limit: to.Ptr[int32](1000),
	// 				Query: to.Ptr("Heartbeat | where SourceSystem != '' | project SourceSystem"),
	// 				SourceTable: to.Ptr("Heartbeat"),
	// 				StartSearchTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T04:05:06.000Z"); return t}()),
	// 			},
	// 			TotalRetentionInDays: to.Ptr[int32](50),
	// 			TotalRetentionInDaysAsDefault: to.Ptr(true),
	// 		},
	// 		SystemData: &armoperationalinsights.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T04:05:06.000Z"); return t}()),
	// 			CreatedBy: to.Ptr("00000000-0000-0000-0000-00000000000"),
	// 		},
	// 	}
}
