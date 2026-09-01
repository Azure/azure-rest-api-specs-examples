package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-07-15-preview/CreateAccountWithAgentHostingConfiguration.json
func ExampleAccountsClient_BeginCreate_createAFoundryAccountWithCustomerOwnedAksHosting() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "myResourceGroup", "foundryByocAccount", armcognitiveservices.Account{
		Identity: &armcognitiveservices.Identity{
			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcognitiveservices.UserAssignedIdentity{
				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/account-control-plane": {},
			},
		},
		Kind:     to.Ptr("AIServices"),
		Location: to.Ptr("West US"),
		Properties: &armcognitiveservices.AccountProperties{
			AgentHostingConfigurations: []armcognitiveservices.AgentHostingConfigurationClassification{
				&armcognitiveservices.ManagedClusterAgentHostingConfiguration{
					Name:                                to.Ptr("default"),
					HostingType:                         to.Ptr(armcognitiveservices.AgentHostingTypeManagedCluster),
					HostingManagementIdentityResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/account-control-plane"),
					WorkloadIdentityResourceID:          to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aks-workload"),
					ClusterResourceID:                   to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/cluster1"),
					StorageAccountResourceID:            to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/storage1"),
				},
			},
		},
		SKU: &armcognitiveservices.SKU{
			Name: to.Ptr("S0"),
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
	// res = armcognitiveservices.AccountsClientCreateResponse{
	// 	Account: armcognitiveservices.Account{
	// 		Name: to.Ptr("foundryByocAccount"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 		Etag: to.Ptr("W/\"datetime'2026-08-05T03%3A00%3A00.0000000Z'\""),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/foundryByocAccount"),
	// 		Identity: &armcognitiveservices.Identity{
	// 			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeUserAssigned),
	// 			TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			UserAssignedIdentities: map[string]*armcognitiveservices.UserAssignedIdentity{
	// 				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/account-control-plane": &armcognitiveservices.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("b5cf119e-a5c2-42c7-802f-592e0efb169f"),
	// 					ClientID: to.Ptr("2f7ee82a-3d7f-4a7f-b8de-3c43ad9478a2"),
	// 				},
	// 			},
	// 		},
	// 		Kind: to.Ptr("AIServices"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armcognitiveservices.AccountProperties{
	// 			Endpoint: to.Ptr("https://foundrybyocaccount.cognitiveservices.azure.com/"),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 			AgentHostingConfigurations: []armcognitiveservices.AgentHostingConfigurationClassification{
	// 				&armcognitiveservices.ManagedClusterAgentHostingConfiguration{
	// 					Name: to.Ptr("default"),
	// 					HostingType: to.Ptr(armcognitiveservices.AgentHostingTypeManagedCluster),
	// 					HostingManagementIdentityResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/account-control-plane"),
	// 					WorkloadIdentityResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aks-workload"),
	// 					ClusterResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/cluster1"),
	// 					StorageAccountResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/storage1"),
	// 				},
	// 			},
	// 		},
	// 		SKU: &armcognitiveservices.SKU{
	// 			Name: to.Ptr("S0"),
	// 		},
	// 	},
	// }
}
