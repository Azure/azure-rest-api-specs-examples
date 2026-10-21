package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/Clusters_CreateOrReplaceWithDataEncryption.json
func ExampleClustersClient_BeginCreateOrReplace_createAHorizonDbClusterWithCustomerManagedKeyDataEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrReplace(ctx, "example-resource-group", "example-cluster", armhorizondb.Cluster{
		Location: to.Ptr("eastus2"),
		Identity: &armhorizondb.ClusterManagedServiceIdentity{
			Type: to.Ptr(armhorizondb.ClusterIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armhorizondb.UserAssignedIdentity{
				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/primary-encryption-identity": {},
			},
		},
		Properties: &armhorizondb.ClusterProperties{
			PostgreSQLVersion: to.Ptr(armhorizondb.PostgreSQLVersionV17),
			Backup: &armhorizondb.BackupProperties{
				GeoRedundancy: to.Ptr(armhorizondb.GeoRedundancyStateDisabled),
				RetentionDays: to.Ptr[int32](14),
			},
			DataEncryption: &armhorizondb.ClusterDataEncryption{
				Type:                          to.Ptr(armhorizondb.ClusterDataEncryptionTypeAzureKeyVault),
				PrimaryKeyURI:                 to.Ptr("https://primary-key-vault.vault.azure.net/keys/horizondb-primary-key"),
				PrimaryUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/primary-encryption-identity"),
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
	// res = armhorizondb.ClustersClientCreateOrReplaceResponse{
	// 	Cluster: armhorizondb.Cluster{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster"),
	// 		Name: to.Ptr("example-cluster"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters"),
	// 		Location: to.Ptr("eastus2"),
	// 		Identity: &armhorizondb.ClusterManagedServiceIdentity{
	// 			Type: to.Ptr(armhorizondb.ClusterIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armhorizondb.UserAssignedIdentity{
	// 				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/primary-encryption-identity": &armhorizondb.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					ClientID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armhorizondb.ClusterProperties{
	// 			PostgreSQLVersion: to.Ptr(armhorizondb.PostgreSQLVersionV17),
	// 			CreationMode: to.Ptr(armhorizondb.ClusterCreateModeCreate),
	// 			Network: &armhorizondb.Network{
	// 				PublicAccess: to.Ptr(armhorizondb.PublicNetworkAccessStateDisabled),
	// 			},
	// 			Backup: &armhorizondb.BackupProperties{
	// 				GeoRedundancy: to.Ptr(armhorizondb.GeoRedundancyStateDisabled),
	// 				RetentionDays: to.Ptr[int32](14),
	// 			},
	// 			DataEncryption: &armhorizondb.ClusterDataEncryption{
	// 				Type: to.Ptr(armhorizondb.ClusterDataEncryptionTypeAzureKeyVault),
	// 				PrimaryKeyURI: to.Ptr("https://primary-key-vault.vault.azure.net/keys/horizondb-primary-key"),
	// 				PrimaryUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/primary-encryption-identity"),
	// 				PrimaryEncryptionKeyStatus: to.Ptr(armhorizondb.ClusterEncryptionKeyStatusValid),
	// 				PrimaryEncryptionKeyCurrentVersion: to.Ptr("33333333333333333333333333333333"),
	// 			},
	// 			State: to.Ptr(armhorizondb.ClusterStateUpdating),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateInProgress),
	// 			ReadWritePool: &armhorizondb.ReadWritePoolProperties{
	// 			},
	// 		},
	// 	},
	// }
}
