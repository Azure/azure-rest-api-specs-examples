package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v3"
)

// Generated from example definition: 2025-07-01/WorkspacesPurge.json
func ExampleWorkspacePurgeClient_Purge_workspacePurge() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspacePurgeClient().Purge(ctx, "OIAutoRest5123", "aztest5048", armoperationalinsights.WorkspacePurgeBody{
		Filters: []*armoperationalinsights.WorkspacePurgeBodyFilters{
			{
				Column:   to.Ptr("TimeGenerated"),
				Operator: to.Ptr(">"),
				Value:    "2017-09-01T00:00:00",
			},
		},
		Table: to.Ptr("Heartbeat"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
