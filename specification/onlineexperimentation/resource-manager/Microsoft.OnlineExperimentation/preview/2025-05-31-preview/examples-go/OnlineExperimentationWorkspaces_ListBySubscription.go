package armonlineexperimentation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/onlineexperimentation/armonlineexperimentation"
)

// Generated from example definition: 2025-05-31-preview/OnlineExperimentationWorkspaces_ListBySubscription.json
func ExampleWorkspacesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armonlineexperimentation.NewClientFactory("fa5fc227-a624-475e-b696-cdd604c735bc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armonlineexperimentation.WorkspacesClientListBySubscriptionResponse{
		// 	WorkspaceListResult: armonlineexperimentation.WorkspaceListResult{
		// 		Value: []*armonlineexperimentation.Workspace{
		// 			{
		// 				ID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OnlineExperimentation/workspaces/expworkspace3"),
		// 				Name: to.Ptr("expworkspace3"),
		// 				Type: to.Ptr("Microsoft.OnlineExperimentation/workspaces"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armonlineexperimentation.WorkspaceProperties{
		// 					WorkspaceID: to.Ptr("02270d43-f68a-401b-b526-86942525c352"),
		// 					ProvisioningState: to.Ptr(armonlineexperimentation.ResourceProvisioningStateSucceeded),
		// 					LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OperationalInsights/workspaces/log9871"),
		// 					LogsExporterStorageAccountResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto9871"),
		// 					AppConfigurationResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.AppConfiguration/configurationStores/appconfig9871"),
		// 					Endpoint: to.Ptr("https://expworkspace3.westus2.exp.azure.net"),
		// 				},
		// 				Identity: &armonlineexperimentation.ManagedServiceIdentity{
		// 					Type: to.Ptr(armonlineexperimentation.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armonlineexperimentation.UserAssignedIdentity{
		// 						"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armonlineexperimentation.UserAssignedIdentity{
		// 							ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
		// 							PrincipalID: to.Ptr("075a0ca6-43f6-4434-9abf-c9b1b79f9219"),
		// 						},
		// 						"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": &armonlineexperimentation.UserAssignedIdentity{
		// 							ClientID: to.Ptr("47429305-c0d3-40bc-8595-61946c4df3dc"),
		// 							PrincipalID: to.Ptr("bf9ebbc8-b92d-4752-8e66-c999000326e0"),
		// 						},
		// 					},
		// 				},
		// 				SKU: &armonlineexperimentation.WorkspaceSKU{
		// 					Name: to.Ptr(armonlineexperimentation.WorkspaceSKUNameF0),
		// 					Tier: to.Ptr(armonlineexperimentation.WorkspaceSKUTierFree),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OnlineExperimentation/workspaces/expworkspace2"),
		// 				Name: to.Ptr("expworkspace2"),
		// 				Type: to.Ptr("Microsoft.OnlineExperimentation/workspaces"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armonlineexperimentation.WorkspaceProperties{
		// 					WorkspaceID: to.Ptr("02270d43-f68a-401b-b526-86942525c353"),
		// 					ProvisioningState: to.Ptr(armonlineexperimentation.ResourceProvisioningStateSucceeded),
		// 					LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OperationalInsights/workspaces/log9871"),
		// 					LogsExporterStorageAccountResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto9871"),
		// 					AppConfigurationResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.AppConfiguration/configurationStores/appconfig9871"),
		// 					Endpoint: to.Ptr("https://expworkspace2.westus2.exp.azure.net"),
		// 				},
		// 				Identity: &armonlineexperimentation.ManagedServiceIdentity{
		// 					Type: to.Ptr(armonlineexperimentation.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armonlineexperimentation.UserAssignedIdentity{
		// 						"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armonlineexperimentation.UserAssignedIdentity{
		// 							ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
		// 							PrincipalID: to.Ptr("075a0ca6-43f6-4434-9abf-c9b1b79f9219"),
		// 						},
		// 						"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": &armonlineexperimentation.UserAssignedIdentity{
		// 							ClientID: to.Ptr("47429305-c0d3-40bc-8595-61946c4df3dc"),
		// 							PrincipalID: to.Ptr("bf9ebbc8-b92d-4752-8e66-c999000326e0"),
		// 						},
		// 					},
		// 				},
		// 				SKU: &armonlineexperimentation.WorkspaceSKU{
		// 					Name: to.Ptr(armonlineexperimentation.WorkspaceSKUNameF0),
		// 					Tier: to.Ptr(armonlineexperimentation.WorkspaceSKUTierFree),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OnlineExperimentation/workspaces/expworkspace4"),
		// 				Name: to.Ptr("expworkspace4"),
		// 				Type: to.Ptr("Microsoft.OnlineExperimentation/workspaces"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armonlineexperimentation.WorkspaceProperties{
		// 					WorkspaceID: to.Ptr("02270d43-f68a-401b-b526-86942525c354"),
		// 					ProvisioningState: to.Ptr(armonlineexperimentation.ResourceProvisioningStateSucceeded),
		// 					LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OperationalInsights/workspaces/log9871"),
		// 					LogsExporterStorageAccountResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto9871"),
		// 					AppConfigurationResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.AppConfiguration/configurationStores/appconfig9871"),
		// 					Endpoint: to.Ptr("https://expworkspace4.westus2.exp.azure.net"),
		// 				},
		// 				Identity: &armonlineexperimentation.ManagedServiceIdentity{
		// 					Type: to.Ptr(armonlineexperimentation.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armonlineexperimentation.UserAssignedIdentity{
		// 						"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armonlineexperimentation.UserAssignedIdentity{
		// 							ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
		// 							PrincipalID: to.Ptr("075a0ca6-43f6-4434-9abf-c9b1b79f9219"),
		// 						},
		// 						"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": &armonlineexperimentation.UserAssignedIdentity{
		// 							ClientID: to.Ptr("47429305-c0d3-40bc-8595-61946c4df3dc"),
		// 							PrincipalID: to.Ptr("bf9ebbc8-b92d-4752-8e66-c999000326e0"),
		// 						},
		// 					},
		// 				},
		// 				SKU: &armonlineexperimentation.WorkspaceSKU{
		// 					Name: to.Ptr(armonlineexperimentation.WorkspaceSKUNameF0),
		// 					Tier: to.Ptr(armonlineexperimentation.WorkspaceSKUTierFree),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OnlineExperimentation/workspaces/expworkspace5"),
		// 				Name: to.Ptr("expworkspace5"),
		// 				Type: to.Ptr("Microsoft.OnlineExperimentation/workspaces"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armonlineexperimentation.WorkspaceProperties{
		// 					WorkspaceID: to.Ptr("02270d43-f68a-401b-b526-86942525c354"),
		// 					ProvisioningState: to.Ptr(armonlineexperimentation.ResourceProvisioningStateSucceeded),
		// 					LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OperationalInsights/workspaces/log9871"),
		// 					LogsExporterStorageAccountResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto9871"),
		// 					AppConfigurationResourceID: to.Ptr("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.AppConfiguration/configurationStores/appconfig9871"),
		// 					Endpoint: to.Ptr("https://expworkspace5.westus2.exp.azure.net"),
		// 				},
		// 				Identity: &armonlineexperimentation.ManagedServiceIdentity{
		// 					Type: to.Ptr(armonlineexperimentation.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armonlineexperimentation.UserAssignedIdentity{
		// 						"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armonlineexperimentation.UserAssignedIdentity{
		// 							ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
		// 							PrincipalID: to.Ptr("075a0ca6-43f6-4434-9abf-c9b1b79f9219"),
		// 						},
		// 						"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": &armonlineexperimentation.UserAssignedIdentity{
		// 							ClientID: to.Ptr("47429305-c0d3-40bc-8595-61946c4df3dc"),
		// 							PrincipalID: to.Ptr("bf9ebbc8-b92d-4752-8e66-c999000326e0"),
		// 						},
		// 					},
		// 				},
		// 				SKU: &armonlineexperimentation.WorkspaceSKU{
		// 					Name: to.Ptr(armonlineexperimentation.WorkspaceSKUNameF0),
		// 					Tier: to.Ptr(armonlineexperimentation.WorkspaceSKUTierFree),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
