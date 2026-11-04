package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-08-15-preview/ExternalClusters_ListBySubscription.json
func ExampleExternalClustersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExternalClustersClient().NewListBySubscriptionPager(nil)
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
		// page = armnetapp.ExternalClustersClientListBySubscriptionResponse{
		// 	ExternalClusterList: armnetapp.ExternalClusterList{
		// 		Value: []*armnetapp.ExternalCluster{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contoso-external-cluster"),
		// 				Name: to.Ptr("contoso-external-cluster"),
		// 				Type: to.Ptr("Microsoft.NetApp/externalClusters"),
		// 				Location: to.Ptr("eastus"),
		// 				Identity: &armnetapp.ManagedServiceIdentity{
		// 					Type: to.Ptr(armnetapp.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armnetapp.UserAssignedIdentity{
		// 						"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contosoIdentity1": &armnetapp.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						},
		// 					},
		// 				},
		// 				Properties: &armnetapp.ExternalClusterProperties{
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					InterclusterLifsIPAddresses: []*string{
		// 						to.Ptr("10.0.1.100"),
		// 						to.Ptr("10.0.1.101"),
		// 					},
		// 					ClusterPeerPassphraseKeyVault: &armnetapp.KeyVaultClusterPeerPassphraseProperties{
		// 						KeyVaultURI: to.Ptr("https://contoso-vault.vault.azure.net/"),
		// 						SecretName: to.Ptr("cluster-passphrase"),
		// 						UserAssignedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contosoIdentity1"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg2/providers/Microsoft.NetApp/externalClusters/external-cluster-2"),
		// 				Name: to.Ptr("external-cluster-2"),
		// 				Type: to.Ptr("Microsoft.NetApp/externalClusters"),
		// 				Location: to.Ptr("westus"),
		// 				Identity: &armnetapp.ManagedServiceIdentity{
		// 					Type: to.Ptr(armnetapp.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armnetapp.UserAssignedIdentity{
		// 						"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg2/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": &armnetapp.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						},
		// 					},
		// 				},
		// 				Properties: &armnetapp.ExternalClusterProperties{
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					InterclusterLifsIPAddresses: []*string{
		// 						to.Ptr("10.0.2.100"),
		// 						to.Ptr("10.0.2.101"),
		// 					},
		// 					ClusterPeerPassphraseKeyVault: &armnetapp.KeyVaultClusterPeerPassphraseProperties{
		// 						KeyVaultURI: to.Ptr("https://contoso-vault2.vault.azure.net/"),
		// 						SecretName: to.Ptr("cluster-passphrase-2"),
		// 						UserAssignedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg2/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
