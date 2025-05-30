package armterraform_test

import (
	"context"
	// "github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/terraform/armterraform"
)

// Generated from example definition: 2023-07-01-preview/ExportTerraform.json
func ExampleTerraformClient_BeginExportTerraform() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armterraform.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	// armterraform.ExportResourceGroup{
	// 	Type:              to.Ptr(armterraform.TypeExportResourceGroup),
	// 	ResourceGroupName: to.Ptr("rg1"),
	// }
	var b armterraform.BaseExportModelClassification
	poller, err := clientFactory.NewTerraformClient().BeginExportTerraform(ctx, b, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
