package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/Workspaces_ListByResourceGroup_MaximumSet_Gen.json
func ExampleWorkspacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListByResourceGroupPager("rgdiscovery", nil)
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
		// page = armdiscovery.WorkspacesClientListByResourceGroupResponse{
		// 	WorkspaceListResult: armdiscovery.WorkspaceListResult{
		// 		Value: []*armdiscovery.Workspace{
		// 			{
		// 				ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/Workspaces/czxofsezpcqwhqie"),
		// 				Name: to.Ptr("czxofsezpcqwhqie"),
		// 				Tags: map[string]*string{
		// 					"key8931": to.Ptr("verirbmpdzupxkkeblzfq"),
		// 				},
		// 				Location: to.Ptr("uksouth"),
		// 				Type: to.Ptr("Microsoft.Discovery/Workspaces"),
		// 				SystemData: &armdiscovery.SystemData{
		// 					CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
		// 					CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
		// 					LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 				},
		// 				Properties: &armdiscovery.WorkspaceProperties{
		// 					ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
		// 					SupercomputerIDs: []*string{
		// 						to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/supercomputers/supercomputer12"),
		// 					},
		// 					WorkspaceAPIURI: to.Ptr("https://microsoft.com/a"),
		// 					WorkspaceUIURI: to.Ptr("https://microsoft.com/aygn"),
		// 					WorkspaceIdentity: &armdiscovery.Identity{
		// 						ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedid1"),
		// 					},
		// 					CustomerManagedKeys: to.Ptr(armdiscovery.CustomerManagedKeysEnabled),
		// 					KeyVaultProperties: &armdiscovery.KeyVaultProperties{
		// 						KeyVaultURI: to.Ptr("https://microsoft.com/a"),
		// 						KeyName: to.Ptr("cdrnokqonyvfzot"),
		// 						KeyVersion: to.Ptr("pxfpvedkfuagtnk"),
		// 					},
		// 					LogAnalyticsClusterID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.OperationalInsights/clusters/cluster1"),
		// 					PrivateEndpointConnections: []*armdiscovery.PrivateEndpointConnection{
		// 						{
		// 							Properties: &armdiscovery.PrivateEndpointConnectionProperties{
		// 								GroupIDs: []*string{
		// 									to.Ptr("bpyliugtuio"),
		// 								},
		// 								PrivateEndpoint: &armdiscovery.PrivateEndpoint{
		// 									ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Network/privateEndpoints/privateEndpoint1"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armdiscovery.PrivateLinkServiceConnectionState{
		// 									Status: to.Ptr(armdiscovery.PrivateEndpointServiceConnectionStatusPending),
		// 									Description: to.Ptr("km"),
		// 									ActionsRequired: to.Ptr("xbshniighjomlygqk"),
		// 								},
		// 								ProvisioningState: to.Ptr(armdiscovery.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 							},
		// 							ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/workspaces/workspaces1/privateEndpointConnections/privateEndpointConnection1"),
		// 							Name: to.Ptr("wdmo"),
		// 							Type: to.Ptr("qpcbcqayctbhbbtguk"),
		// 							SystemData: &armdiscovery.SystemData{
		// 								CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
		// 								CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
		// 								LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 							},
		// 						},
		// 					},
		// 					PublicNetworkAccess: to.Ptr(armdiscovery.PublicNetworkAccessEnabled),
		// 					AgentSubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/agentSubnet1"),
		// 					PrivateEndpointSubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/privateEndpointSubnet1"),
		// 					WorkspaceSubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/workspaceSubnet1"),
		// 					ManagedResourceGroup: to.Ptr("nfsdgndlqwmeuhbadtztztggiybk"),
		// 					ManagedOnBehalfOfConfiguration: &armdiscovery.WithMoboBrokerResources{
		// 						MoboBrokerResources: []*armdiscovery.MoboBrokerResource{
		// 							{
		// 								ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Storage/storageAccounts/storage1"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
