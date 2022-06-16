package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/WorkspacesUpdate.json
func ExampleWorkspacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armoperationalinsights.NewWorkspacesClient("00000000-0000-0000-0000-00000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"oiautorest6685",
		"oiautorest6685",
		armoperationalinsights.WorkspacePatch{
			Properties: &armoperationalinsights.WorkspaceProperties{
				RetentionInDays: to.Ptr[int32](30),
				SKU: &armoperationalinsights.WorkspaceSKU{
					Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
				},
				WorkspaceCapping: &armoperationalinsights.WorkspaceCapping{
					DailyQuotaGb: to.Ptr[float64](-1),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
