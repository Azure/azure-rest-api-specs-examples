package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be46becafeb29aa993898709e35759d3643b2809/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/TablesUpsert.json
func ExampleTablesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTablesClient().BeginUpdate(ctx, "oiautorest6685", "oiautorest6685", "AzureNetworkFlow", armoperationalinsights.Table{
		Properties: &armoperationalinsights.TableProperties{
			Schema: &armoperationalinsights.Schema{
				Name: to.Ptr("AzureNetworkFlow"),
				Columns: []*armoperationalinsights.Column{
					{
						Name: to.Ptr("MyNewColumn"),
						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumGUID),
					}},
			},
			RetentionInDays:      to.Ptr[int32](45),
			TotalRetentionInDays: to.Ptr[int32](70),
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
	// res.Table = armoperationalinsights.Table{
	// 	Name: to.Ptr("AzureNetworkFlow"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/Microsoft.OperationalInsights/workspaces/oiautorest6685/tables/AzureNetworkFlow"),
	// 	Properties: &armoperationalinsights.TableProperties{
	// 		Schema: &armoperationalinsights.Schema{
	// 			Name: to.Ptr("AzureNetworkFlow"),
	// 			Columns: []*armoperationalinsights.Column{
	// 				{
	// 					Name: to.Ptr("MyNewColumn"),
	// 					Type: to.Ptr(armoperationalinsights.ColumnTypeEnumGUID),
	// 					IsDefaultDisplay: to.Ptr(false),
	// 					IsHidden: to.Ptr(false),
	// 			}},
	// 			Solutions: []*string{
	// 				to.Ptr("LogManagement")},
	// 				StandardColumns: []*armoperationalinsights.Column{
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumGUID),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(true),
	// 					},
	// 					{
	// 						Name: to.Ptr("SourceSystem"),
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
	// 						Name: to.Ptr("AgentID"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SourceIP"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Protocol"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SourcePort"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("DestinationPort"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TcpFlags"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Packets"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Bytes"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("BytesOut"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("DurationInMs"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("RstCount"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("MaxSampleRtt"),
	// 						Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
	// 						IsDefaultDisplay: to.Ptr(false),
	// 						IsHidden: to.Ptr(false),
	// 				}},
	// 				TableSubType: to.Ptr(armoperationalinsights.TableSubTypeEnumDataCollectionRuleBased),
	// 				TableType: to.Ptr(armoperationalinsights.TableTypeEnumMicrosoft),
	// 			},
	// 			ArchiveRetentionInDays: to.Ptr[int32](25),
	// 			Plan: to.Ptr(armoperationalinsights.TablePlanEnumAnalytics),
	// 			ProvisioningState: to.Ptr(armoperationalinsights.ProvisioningStateEnumSucceeded),
	// 			RetentionInDays: to.Ptr[int32](45),
	// 			RetentionInDaysAsDefault: to.Ptr(false),
	// 			TotalRetentionInDays: to.Ptr[int32](70),
	// 			TotalRetentionInDaysAsDefault: to.Ptr(false),
	// 		},
	// 	}
}
