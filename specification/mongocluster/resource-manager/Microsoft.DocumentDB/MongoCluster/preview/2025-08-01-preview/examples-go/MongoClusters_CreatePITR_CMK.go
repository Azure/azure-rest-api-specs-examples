package armmongocluster_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
)

// Generated from example definition: 2025-08-01-preview/MongoClusters_CreatePITR_CMK.json
func ExampleMongoClustersClient_BeginCreateOrUpdate_createsAMongoClusterResourceWithCustomerManagedKeyEncryptionFromAPointInTimeRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongocluster.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMongoClustersClient().BeginCreateOrUpdate(ctx, "TestResourceGroup", "myMongoCluster", armmongocluster.MongoCluster{
		Identity: &armmongocluster.ManagedServiceIdentity{
			Type: to.Ptr(armmongocluster.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armmongocluster.UserAssignedIdentity{
				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": {},
			},
		},
		Location: to.Ptr("westus2"),
		Properties: &armmongocluster.Properties{
			CreateMode: to.Ptr(armmongocluster.CreateModePointInTimeRestore),
			RestoreParameters: &armmongocluster.RestoreParameters{
				PointInTimeUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-13T20:07:35Z"); return t }()),
				SourceResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myOtherMongoCluster"),
			},
			Encryption: &armmongocluster.EncryptionProperties{
				CustomerManagedKeyEncryption: &armmongocluster.CustomerManagedKeyEncryptionProperties{
					KeyEncryptionKeyIdentity: &armmongocluster.KeyEncryptionKeyIdentity{
						IdentityType:                   to.Ptr(armmongocluster.KeyEncryptionKeyIdentityTypeUserAssignedIdentity),
						UserAssignedIdentityResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
					},
					KeyEncryptionKeyURL: to.Ptr("https://myVault.vault.azure.net/keys/myKey"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmongocluster.MongoClustersClientCreateOrUpdateResponse{
	// 	MongoCluster: &armmongocluster.MongoCluster{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myMongoCluster"),
	// 		Name: to.Ptr("myMongoCluster"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/mongoClusters"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("westus2"),
	// 		SystemData: &armmongocluster.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armmongocluster.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armmongocluster.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-02T17:18:19.1234567Z"); return t}()),
	// 		},
	// 		Properties: &armmongocluster.Properties{
	// 			ProvisioningState: to.Ptr(armmongocluster.ProvisioningStateSucceeded),
	// 			ClusterStatus: to.Ptr(armmongocluster.StatusReady),
	// 			Administrator: &armmongocluster.AdministratorProperties{
	// 				UserName: to.Ptr("mongoAdmin"),
	// 			},
	// 			AuthConfig: &armmongocluster.AuthConfigProperties{
	// 				AllowedModes: []*armmongocluster.AuthenticationMode{
	// 					to.Ptr(armmongocluster.AuthenticationModeNativeAuth),
	// 				},
	// 			},
	// 			ServerVersion: to.Ptr("7.0"),
	// 			Storage: &armmongocluster.StorageProperties{
	// 				SizeGb: to.Ptr[int64](32),
	// 				Type: to.Ptr(armmongocluster.StorageTypePremiumSSD),
	// 			},
	// 			Compute: &armmongocluster.ComputeProperties{
	// 				Tier: to.Ptr("M30"),
	// 			},
	// 			Sharding: &armmongocluster.ShardingProperties{
	// 				ShardCount: to.Ptr[int32](1),
	// 			},
	// 			HighAvailability: &armmongocluster.HighAvailabilityProperties{
	// 				TargetMode: to.Ptr(armmongocluster.HighAvailabilityModeSameZone),
	// 			},
	// 			Backup: &armmongocluster.BackupProperties{
	// 				EarliestRestoreTime: to.Ptr("2025-06-01T20:07:35Z"),
	// 			},
	// 			PublicNetworkAccess: to.Ptr(armmongocluster.PublicNetworkAccessEnabled),
	// 			ConnectionString: to.Ptr("mongodb+srv://<user>:<password>@myMongoCluster.mongocluster.cosmos.azure.com"),
	// 			Replica: &armmongocluster.ReplicationProperties{
	// 				Role: to.Ptr(armmongocluster.ReplicationRolePrimary),
	// 				ReplicationState: to.Ptr(armmongocluster.ReplicationStateActive),
	// 			},
	// 			InfrastructureVersion: to.Ptr("2.0"),
	// 			DataAPI: &armmongocluster.DataAPIProperties{
	// 				Mode: to.Ptr(armmongocluster.DataAPIModeDisabled),
	// 			},
	// 			CreateMode: to.Ptr(armmongocluster.CreateModeDefault),
	// 			Encryption: &armmongocluster.EncryptionProperties{
	// 				CustomerManagedKeyEncryption: &armmongocluster.CustomerManagedKeyEncryptionProperties{
	// 					KeyEncryptionKeyIdentity: &armmongocluster.KeyEncryptionKeyIdentity{
	// 						IdentityType: to.Ptr(armmongocluster.KeyEncryptionKeyIdentityTypeUserAssignedIdentity),
	// 						UserAssignedIdentityResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
	// 					},
	// 					KeyEncryptionKeyURL: to.Ptr("https://myVault.vault.azure.net/keys/myKey"),
	// 				},
	// 			},
	// 		},
	// 		Identity: &armmongocluster.ManagedServiceIdentity{
	// 			Type: to.Ptr(armmongocluster.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armmongocluster.UserAssignedIdentity{
	// 				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": &armmongocluster.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
