package armmongocluster_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
)

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster/TempTypeSpecFiles/DocumentDB.MongoCluster.Management/examples/2024-07-01/MongoClusters_ListByResourceGroup.json
func ExampleMongoClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongocluster.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMongoClustersClient().NewListByResourceGroupPager("TestResourceGroup", nil)
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
		// page = armmongocluster.MongoClustersClientListByResourceGroupResponse{
		// 	ListResult: armmongocluster.ListResult{
		// 		Value: []*armmongocluster.MongoCluster{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myMongoCluster"),
		// 				Name: to.Ptr("myMongoCluster"),
		// 				Type: to.Ptr("/Microsoft.DocumentDB/mongoClusters"),
		// 				Tags: map[string]*string{
		// 					"additionalProp1": to.Ptr("string"),
		// 					"additionalProp2": to.Ptr("string"),
		// 					"additionalProp3": to.Ptr("string"),
		// 				},
		// 				SystemData: &armmongocluster.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armmongocluster.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armmongocluster.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				},
		// 				Properties: &armmongocluster.Properties{
		// 					ProvisioningState: to.Ptr(armmongocluster.ProvisioningStateSucceeded),
		// 					Administrator: &armmongocluster.AdministratorProperties{
		// 						UserName: to.Ptr("mongoAdmin"),
		// 					},
		// 					ServerVersion: to.Ptr("5.0"),
		// 					Storage: &armmongocluster.StorageProperties{
		// 						SizeGb: to.Ptr[int64](128),
		// 					},
		// 					Compute: &armmongocluster.ComputeProperties{
		// 						Tier: to.Ptr("M30"),
		// 					},
		// 					Sharding: &armmongocluster.ShardingProperties{
		// 						ShardCount: to.Ptr[int32](4),
		// 					},
		// 					HighAvailability: &armmongocluster.HighAvailabilityProperties{
		// 						TargetMode: to.Ptr(armmongocluster.HighAvailabilityModeSameZone),
		// 					},
		// 					Backup: &armmongocluster.BackupProperties{
		// 						EarliestRestoreTime: to.Ptr("2023-01-13T20:07:35Z"),
		// 					},
		// 					PreviewFeatures: []*armmongocluster.PreviewFeature{
		// 					},
		// 					InfrastructureVersion: to.Ptr("2.0"),
		// 					PublicNetworkAccess: to.Ptr(armmongocluster.PublicNetworkAccessEnabled),
		// 					ConnectionString: to.Ptr("mongodb+srv://<user>:<password>@myMongoCluster.mongocluster.cosmos.azure.com"),
		// 					Replica: &armmongocluster.ReplicationProperties{
		// 						ReplicationState: to.Ptr(armmongocluster.ReplicationStateActive),
		// 						Role: to.Ptr(armmongocluster.ReplicationRolePrimary),
		// 					},
		// 				},
		// 				Location: to.Ptr("westus2"),
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myMongoCluster2"),
		// 				Name: to.Ptr("myMongoCluster2"),
		// 				Type: to.Ptr("/Microsoft.DocumentDB/mongoClusters"),
		// 				Tags: map[string]*string{
		// 					"additionalProp1": to.Ptr("string"),
		// 				},
		// 				SystemData: &armmongocluster.SystemData{
		// 					CreatedBy: to.Ptr("user2"),
		// 					CreatedByType: to.Ptr(armmongocluster.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armmongocluster.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T17:18:19.1234567Z"); return t}()),
		// 				},
		// 				Properties: &armmongocluster.Properties{
		// 					ProvisioningState: to.Ptr(armmongocluster.ProvisioningStateSucceeded),
		// 					Administrator: &armmongocluster.AdministratorProperties{
		// 						UserName: to.Ptr("mongoAdmin"),
		// 					},
		// 					ServerVersion: to.Ptr("5.0"),
		// 					Storage: &armmongocluster.StorageProperties{
		// 						SizeGb: to.Ptr[int64](256),
		// 					},
		// 					Compute: &armmongocluster.ComputeProperties{
		// 						Tier: to.Ptr("M40"),
		// 					},
		// 					Sharding: &armmongocluster.ShardingProperties{
		// 						ShardCount: to.Ptr[int32](2),
		// 					},
		// 					HighAvailability: &armmongocluster.HighAvailabilityProperties{
		// 						TargetMode: to.Ptr(armmongocluster.HighAvailabilityModeSameZone),
		// 					},
		// 					Backup: &armmongocluster.BackupProperties{
		// 						EarliestRestoreTime: to.Ptr("2023-01-13T20:07:35Z"),
		// 					},
		// 					PreviewFeatures: []*armmongocluster.PreviewFeature{
		// 					},
		// 					InfrastructureVersion: to.Ptr("1.0"),
		// 					PublicNetworkAccess: to.Ptr(armmongocluster.PublicNetworkAccessEnabled),
		// 					ConnectionString: to.Ptr("mongodb+srv://<user>:<password>@myMongoCluster2.mongocluster.cosmos.azure.com"),
		// 					Replica: &armmongocluster.ReplicationProperties{
		// 						ReplicationState: to.Ptr(armmongocluster.ReplicationStateActive),
		// 						Role: to.Ptr(armmongocluster.ReplicationRolePrimary),
		// 					},
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
