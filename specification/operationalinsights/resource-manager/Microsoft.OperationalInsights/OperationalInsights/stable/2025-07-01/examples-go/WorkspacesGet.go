package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be46becafeb29aa993898709e35759d3643b2809/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/WorkspacesGet.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Get(ctx, "oiautorest6685", "oiautorest6685", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workspace = armoperationalinsights.Workspace{
	// 	Name: to.Ptr("oiautorest6685"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/workspaces/oiautorest6685"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armoperationalinsights.WorkspaceProperties{
	// 		CustomerID: to.Ptr("5b02755b-5bf4-430c-9487-45502a2a7e62"),
	// 		Failover: &armoperationalinsights.WorkspaceFailoverProperties{
	// 			LastModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-05T07:35:51.000Z"); return t}()),
	// 			State: to.Ptr(armoperationalinsights.WorkspaceFailoverStateActive),
	// 		},
	// 		ProvisioningState: to.Ptr(armoperationalinsights.WorkspaceEntityStatusSucceeded),
	// 		PublicNetworkAccessForIngestion: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		PublicNetworkAccessForQuery: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		Replication: &armoperationalinsights.WorkspaceReplicationProperties{
	// 			CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-04T07:35:51.000Z"); return t}()),
	// 			Enabled: to.Ptr(true),
	// 			LastModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-04T07:35:51.000Z"); return t}()),
	// 			Location: to.Ptr("westus3"),
	// 			ProvisioningState: to.Ptr(armoperationalinsights.WorkspaceReplicationStateSucceeded),
	// 		},
	// 		RetentionInDays: to.Ptr[int32](30),
	// 		SKU: &armoperationalinsights.WorkspaceSKU{
	// 			Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
	// 		},
	// 	},
	// }
}
