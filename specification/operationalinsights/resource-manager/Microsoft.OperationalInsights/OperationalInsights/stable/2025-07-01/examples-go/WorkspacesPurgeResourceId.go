package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be46becafeb29aa993898709e35759d3643b2809/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/WorkspacesPurgeResourceId.json
func ExampleWorkspacePurgeClient_Purge_workspacePurgeResourceId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspacePurgeClient().Purge(ctx, "OIAutoRest5123", "aztest5048", armoperationalinsights.WorkspacePurgeBody{
		Filters: []*armoperationalinsights.WorkspacePurgeBodyFilters{
			{
				Column:   to.Ptr("_ResourceId"),
				Operator: to.Ptr("=="),
				Value:    "/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/SomeResourceGroup/providers/microsoft.insights/components/AppInsightResource",
			}},
		Table: to.Ptr("Heartbeat"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
