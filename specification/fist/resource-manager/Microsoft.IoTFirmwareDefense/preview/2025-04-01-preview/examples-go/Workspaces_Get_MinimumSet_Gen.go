package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-04-01-preview/Workspaces_Get_MinimumSet_Gen.json
func ExampleWorkspacesClient_Get_workspacesGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("9781B4B5-0922-472A-80F0-B743D0596694", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Get(ctx, "rgworkspaces", "E_US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotfirmwaredefense.WorkspacesClientGetResponse{
	// 	Workspace: &armiotfirmwaredefense.Workspace{
	// 		Name: to.Ptr("tbrqhnzpsatbrnhtj"),
	// 		Type: to.Ptr("angrpzpxmuppzqpzpljasjirao"),
	// 		ID: to.Ptr("/subscriptions/blah/resourceGroups/blah/providers/blah/workspaces/blah"),
	// 		Location: to.Ptr("zxwrauqottoy"),
	// 	},
	// }
}
