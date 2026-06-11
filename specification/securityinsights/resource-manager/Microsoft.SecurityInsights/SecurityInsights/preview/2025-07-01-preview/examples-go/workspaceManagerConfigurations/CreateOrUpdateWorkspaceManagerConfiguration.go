package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/workspaceManagerConfigurations/CreateOrUpdateWorkspaceManagerConfiguration.json
func ExampleWorkspaceManagerConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceManagerConfigurationsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "default", armsecurityinsights.WorkspaceManagerConfiguration{
		Properties: &armsecurityinsights.WorkspaceManagerConfigurationProperties{
			Mode: to.Ptr(armsecurityinsights.ModeEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.WorkspaceManagerConfigurationsClientCreateOrUpdateResponse{
	// 	WorkspaceManagerConfiguration: armsecurityinsights.WorkspaceManagerConfiguration{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/workspaceManagerConfigurations"),
	// 		Etag: to.Ptr("\"3f6451dd-1b58-4bef-bce7-72eba6b354d7\""),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/workspaceManagerConfigurations/default"),
	// 		Properties: &armsecurityinsights.WorkspaceManagerConfigurationProperties{
	// 			Mode: to.Ptr(armsecurityinsights.ModeEnabled),
	// 		},
	// 	},
	// }
}
