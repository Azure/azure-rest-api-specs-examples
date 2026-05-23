package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v3"
)

// Generated from example definition: 2025-07-01/WorkspacesPurgeOperation.json
func ExampleWorkspacePurgeClient_GetPurgeStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacePurgeClient().GetPurgeStatus(ctx, "OIAutoRest5123", "aztest5048", "purge-970318e7-b859-4edb-8903-83b1b54d0b74", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armoperationalinsights.WorkspacePurgeClientGetPurgeStatusResponse{
	// 	WorkspacePurgeStatusResponse: armoperationalinsights.WorkspacePurgeStatusResponse{
	// 		Status: to.Ptr(armoperationalinsights.PurgeStateCompleted),
	// 	},
	// }
}
