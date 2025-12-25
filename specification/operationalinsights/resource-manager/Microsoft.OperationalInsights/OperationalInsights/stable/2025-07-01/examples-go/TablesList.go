package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be46becafeb29aa993898709e35759d3643b2809/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/TablesList.json
func ExampleTablesClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTablesClient().NewListByWorkspacePager("oiautorest6685", "oiautorest6685", nil)
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
		// page.TablesListResult = armoperationalinsights.TablesListResult{
		// 	Value: []*armoperationalinsights.Table{
		// 		{
		// 			Name: to.Ptr("AzureNetworkFlow"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/Microsoft.OperationalInsights/workspaces/oiautorest6685/tables/AzureNetworkFlow"),
		// 			Properties: &armoperationalinsights.TableProperties{
		// 				Schema: &armoperationalinsights.Schema{
		// 					Name: to.Ptr("AzureNetworkFlow"),
		// 					Solutions: []*string{
		// 						to.Ptr("LogManagement")},
		// 						StandardColumns: []*armoperationalinsights.Column{
		// 							{
		// 								Name: to.Ptr("TenantId"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumGUID),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(true),
		// 							},
		// 							{
		// 								Name: to.Ptr("SourceSystem"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("TimeGenerated"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumDateTime),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("AgentID"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("SourceIP"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("Protocol"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("SourcePort"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("DestinationPort"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("TcpFlags"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("Packets"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("Bytes"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("BytesOut"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("DurationInMs"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("RstCount"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("MaxSampleRtt"),
		// 								Type: to.Ptr(armoperationalinsights.ColumnTypeEnumInt),
		// 								IsDefaultDisplay: to.Ptr(false),
		// 								IsHidden: to.Ptr(false),
		// 						}},
		// 						TableSubType: to.Ptr(armoperationalinsights.TableSubTypeEnumAny),
		// 						TableType: to.Ptr(armoperationalinsights.TableTypeEnumMicrosoft),
		// 					},
		// 					ArchiveRetentionInDays: to.Ptr[int32](25),
		// 					Plan: to.Ptr(armoperationalinsights.TablePlanEnumAnalytics),
		// 					ProvisioningState: to.Ptr(armoperationalinsights.ProvisioningStateEnumSucceeded),
		// 					RetentionInDays: to.Ptr[int32](45),
		// 					RetentionInDaysAsDefault: to.Ptr(false),
		// 					TotalRetentionInDays: to.Ptr[int32](70),
		// 					TotalRetentionInDaysAsDefault: to.Ptr(false),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("SurfaceHubDns"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/Microsoft.OperationalInsights/workspaces/oiautorest6685/tables/SurfaceHubDns"),
		// 				Properties: &armoperationalinsights.TableProperties{
		// 					Schema: &armoperationalinsights.Schema{
		// 						Name: to.Ptr("SurfaceHubDns"),
		// 						Solutions: []*string{
		// 							to.Ptr("LogManagement")},
		// 							StandardColumns: []*armoperationalinsights.Column{
		// 								{
		// 									Name: to.Ptr("TenantId"),
		// 									Type: to.Ptr(armoperationalinsights.ColumnTypeEnumGUID),
		// 									IsDefaultDisplay: to.Ptr(false),
		// 									IsHidden: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SourceSystem"),
		// 									Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
		// 									IsDefaultDisplay: to.Ptr(false),
		// 									IsHidden: to.Ptr(false),
		// 								},
		// 								{
		// 									Name: to.Ptr("TimeGenerated"),
		// 									Type: to.Ptr(armoperationalinsights.ColumnTypeEnumDateTime),
		// 									IsDefaultDisplay: to.Ptr(false),
		// 									IsHidden: to.Ptr(false),
		// 								},
		// 								{
		// 									Name: to.Ptr("QueryName"),
		// 									Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
		// 									IsDefaultDisplay: to.Ptr(false),
		// 									IsHidden: to.Ptr(false),
		// 								},
		// 								{
		// 									Name: to.Ptr("ComputerName"),
		// 									Type: to.Ptr(armoperationalinsights.ColumnTypeEnumString),
		// 									IsDefaultDisplay: to.Ptr(false),
		// 									IsHidden: to.Ptr(false),
		// 							}},
		// 							TableSubType: to.Ptr(armoperationalinsights.TableSubTypeEnumAny),
		// 							TableType: to.Ptr(armoperationalinsights.TableTypeEnumMicrosoft),
		// 						},
		// 						ArchiveRetentionInDays: to.Ptr[int32](0),
		// 						Plan: to.Ptr(armoperationalinsights.TablePlanEnumAnalytics),
		// 						ProvisioningState: to.Ptr(armoperationalinsights.ProvisioningStateEnumSucceeded),
		// 						RetentionInDays: to.Ptr[int32](30),
		// 						RetentionInDaysAsDefault: to.Ptr(false),
		// 						TotalRetentionInDays: to.Ptr[int32](30),
		// 						TotalRetentionInDaysAsDefault: to.Ptr(true),
		// 					},
		// 			}},
		// 		}
	}
}
