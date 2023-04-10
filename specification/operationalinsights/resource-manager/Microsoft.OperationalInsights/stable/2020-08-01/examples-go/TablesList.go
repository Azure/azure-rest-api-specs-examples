package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/TablesList.json
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
		// 			Name: to.Ptr("table1"),
		// 			ID: to.Ptr("/subscriptions/594038b5-1093-476e-a366-482775671c11/resourcegroups/calbot-rg/providers/microsoft.operationalinsights/workspaces/testresourcelock/tables/table1"),
		// 			Properties: &armoperationalinsights.TableProperties{
		// 				RetentionInDays: to.Ptr[int32](30),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("table2"),
		// 			ID: to.Ptr("/subscriptions/594038b5-1093-476e-a366-482775671c11/resourcegroups/calbot-rg/providers/microsoft.operationalinsights/workspaces/testresourcelock/tables/table2"),
		// 			Properties: &armoperationalinsights.TableProperties{
		// 				RetentionInDays: to.Ptr[int32](30),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("table3"),
		// 			ID: to.Ptr("/subscriptions/594038b5-1093-476e-a366-482775671c11/resourcegroups/calbot-rg/providers/microsoft.operationalinsights/workspaces/testresourcelock/tables/table3"),
		// 			Properties: &armoperationalinsights.TableProperties{
		// 				RetentionInDays: to.Ptr[int32](30),
		// 			},
		// 	}},
		// }
	}
}
