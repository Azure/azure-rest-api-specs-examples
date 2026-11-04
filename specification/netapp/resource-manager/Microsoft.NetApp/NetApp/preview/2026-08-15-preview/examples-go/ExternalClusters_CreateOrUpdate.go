package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-08-15-preview/ExternalClusters_CreateOrUpdate.json
func ExampleExternalClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExternalClustersClient().BeginCreateOrUpdate(ctx, "contosoNetAppRg", "contoso-external-cluster", armnetapp.ExternalCluster{
		Location: to.Ptr("eastus"),
		Identity: &armnetapp.ManagedServiceIdentity{
			Type: to.Ptr(armnetapp.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armnetapp.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contosoIdentity1": {},
			},
		},
		Properties: &armnetapp.ExternalClusterProperties{
			InterclusterLifsIPAddresses: []*string{
				to.Ptr("10.0.1.100"),
				to.Ptr("10.0.1.101"),
			},
			ClusterPeerPassphraseKeyVault: &armnetapp.KeyVaultClusterPeerPassphraseProperties{
				KeyVaultURI:                    to.Ptr("https://contoso-vault.vault.azure.net/"),
				SecretName:                     to.Ptr("cluster-passphrase"),
				UserAssignedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contosoIdentity1"),
			},
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
	// res = armnetapp.ExternalClustersClientCreateOrUpdateResponse{
	// 	ExternalCluster: armnetapp.ExternalCluster{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contoso-external-cluster"),
	// 		Name: to.Ptr("contoso-external-cluster"),
	// 		Type: to.Ptr("Microsoft.NetApp/externalClusters"),
	// 		Location: to.Ptr("eastus"),
	// 		Identity: &armnetapp.ManagedServiceIdentity{
	// 			Type: to.Ptr(armnetapp.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armnetapp.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contosoIdentity1": &armnetapp.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armnetapp.ExternalClusterProperties{
	// 			ProvisioningState: to.Ptr(armnetapp.ProvisioningStateCreating),
	// 			InterclusterLifsIPAddresses: []*string{
	// 				to.Ptr("10.0.1.100"),
	// 				to.Ptr("10.0.1.101"),
	// 			},
	// 			ClusterPeerPassphraseKeyVault: &armnetapp.KeyVaultClusterPeerPassphraseProperties{
	// 				KeyVaultURI: to.Ptr("https://contoso-vault.vault.azure.net/"),
	// 				SecretName: to.Ptr("cluster-passphrase"),
	// 				UserAssignedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contosoIdentity1"),
	// 			},
	// 		},
	// 	},
	// }
}
