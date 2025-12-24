package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be46becafeb29aa993898709e35759d3643b2809/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/WorkspacesCreate.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginCreateOrUpdate(ctx, "oiautorest6685", "oiautorest6685", armoperationalinsights.Workspace{
		Location: to.Ptr("australiasoutheast"),
		Tags: map[string]*string{
			"tag1": to.Ptr("val1"),
		},
		Properties: &armoperationalinsights.WorkspaceProperties{
			RetentionInDays: to.Ptr[int32](30),
			SKU: &armoperationalinsights.WorkspaceSKU{
				Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
			},
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
	// res.Workspace = armoperationalinsights.Workspace{
	// 	Name: to.Ptr("AzTest2170"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/workspaces/aztest2170"),
	// 	Location: to.Ptr("australiasoutheast"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("val1"),
	// 	},
	// 	Properties: &armoperationalinsights.WorkspaceProperties{
	// 		CustomerID: to.Ptr("bc089d7b-485c-4aff-a71e-c00f362d8d2f"),
	// 		ProvisioningState: to.Ptr(armoperationalinsights.WorkspaceEntityStatusSucceeded),
	// 		PublicNetworkAccessForIngestion: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		PublicNetworkAccessForQuery: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		RetentionInDays: to.Ptr[int32](30),
	// 		SKU: &armoperationalinsights.WorkspaceSKU{
	// 			Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
	// 		},
	// 	},
	// }
}
