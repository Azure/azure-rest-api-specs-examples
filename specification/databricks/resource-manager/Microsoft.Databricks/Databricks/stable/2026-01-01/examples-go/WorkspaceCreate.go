package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks/v2"
)

// Generated from example definition: 2026-01-01/WorkspaceCreate.json
func ExampleWorkspacesClient_BeginCreateOrUpdate_createOrUpdateWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginCreateOrUpdate(ctx, "rg", "myWorkspace", armdatabricks.Workspace{
		Location: to.Ptr("westus"),
		Properties: &armdatabricks.WorkspaceProperties{
			AccessConnector: &armdatabricks.WorkspacePropertiesAccessConnector{
				ID:           to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/adbrg/providers/Microsoft.Databricks/accessConnectors/myAccessConnector"),
				IdentityType: to.Ptr(armdatabricks.IdentityTypeSystemAssigned),
			},
			ComputeMode: to.Ptr(armdatabricks.ComputeModeHybrid),
			DefaultCatalog: &armdatabricks.DefaultCatalogProperties{
				InitialName: to.Ptr(""),
				InitialType: to.Ptr(armdatabricks.InitialTypeUnityCatalog),
			},
			DefaultStorageFirewall: to.Ptr(armdatabricks.DefaultStorageFirewallEnabled),
			ManagedResourceGroupID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myManagedRG"),
		},
		SKU: &armdatabricks.SKU{
			Name: to.Ptr("premium"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabricks.WorkspacesClientCreateOrUpdateResponse{
	// 	Workspace: armdatabricks.Workspace{
	// 		Type: to.Ptr("Microsoft.Databricks/workspaces"),
	// 		ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg/providers/Microsoft.Databricks/workspaces/myWorkspace"),
	// 		Location: to.Ptr("East US 2"),
	// 		Properties: &armdatabricks.WorkspaceProperties{
	// 			AccessConnector: &armdatabricks.WorkspacePropertiesAccessConnector{
	// 				ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/adbrg/providers/Microsoft.Databricks/accessConnectors/myAccessConnector"),
	// 				IdentityType: to.Ptr(armdatabricks.IdentityTypeSystemAssigned),
	// 			},
	// 			Authorizations: []*armdatabricks.WorkspaceProviderAuthorization{
	// 				{
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					RoleDefinitionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				},
	// 			},
	// 			ComputeMode: to.Ptr(armdatabricks.ComputeModeHybrid),
	// 			CreatedBy: &armdatabricks.CreatedBy{
	// 				ApplicationID: to.Ptr("44444444-4444-4444-4444-444444444444"),
	// 				Oid: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 				Puid: to.Ptr("33333333"),
	// 			},
	// 			CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-20T00:10:29.2858439Z"); return t}()),
	// 			DefaultCatalog: &armdatabricks.DefaultCatalogProperties{
	// 				InitialName: to.Ptr(""),
	// 				InitialType: to.Ptr(armdatabricks.InitialTypeUnityCatalog),
	// 			},
	// 			DefaultStorageFirewall: to.Ptr(armdatabricks.DefaultStorageFirewallEnabled),
	// 			IsUcEnabled: to.Ptr(true),
	// 			ManagedResourceGroupID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myManagedRG"),
	// 			ProvisioningState: to.Ptr(armdatabricks.ProvisioningStateAccepted),
	// 			UIDefinitionURI: to.Ptr("https://path/to/workspaceCreateUiDefinition.json"),
	// 			UpdatedBy: &armdatabricks.CreatedBy{
	// 				ApplicationID: to.Ptr("44444444-4444-4444-4444-444444444444"),
	// 				Oid: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 				Puid: to.Ptr("33333333"),
	// 			},
	// 			WorkspaceID: to.Ptr("5555555555555555"),
	// 			WorkspaceURL: to.Ptr("adb-5555555555555555.19.azuredatabricks.net"),
	// 		},
	// 		SKU: &armdatabricks.SKU{
	// 			Name: to.Ptr("premium"),
	// 		},
	// 	},
	// }
}
