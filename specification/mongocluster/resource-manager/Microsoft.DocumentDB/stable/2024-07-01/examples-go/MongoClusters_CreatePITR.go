package armmongocluster_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
)

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster/TempTypeSpecFiles/DocumentDB.MongoCluster.Management/examples/2024-07-01/MongoClusters_CreatePITR.json
func ExampleMongoClustersClient_BeginCreateOrUpdate_createsAMongoClusterResourceFromAPointInTimeRestore() {
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
		Location: to.Ptr("westus2"),
		Properties: &armmongocluster.Properties{
			CreateMode: to.Ptr(armmongocluster.CreateModePointInTimeRestore),
			RestoreParameters: &armmongocluster.RestoreParameters{
				PointInTimeUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-13T20:07:35Z"); return t }()),
				SourceResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myOtherMongoCluster"),
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
	// 		Type: to.Ptr("/Microsoft.DocumentDB/mongoClusters"),
	// 		Properties: &armmongocluster.Properties{
	// 			ProvisioningState: to.Ptr(armmongocluster.ProvisioningStateSucceeded),
	// 			Administrator: &armmongocluster.AdministratorProperties{
	// 				UserName: to.Ptr("mongoAdmin"),
	// 			},
	// 			ServerVersion: to.Ptr("5.0"),
	// 			Storage: &armmongocluster.StorageProperties{
	// 				SizeGb: to.Ptr[int64](128),
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
	// 				EarliestRestoreTime: to.Ptr("2023-01-13T20:07:35Z"),
	// 			},
	// 			PreviewFeatures: []*armmongocluster.PreviewFeature{
	// 			},
	// 			InfrastructureVersion: to.Ptr("2.0"),
	// 			PublicNetworkAccess: to.Ptr(armmongocluster.PublicNetworkAccessEnabled),
	// 			ConnectionString: to.Ptr("mongodb+srv://<user>:<password>@myMongoCluster.mongocluster.cosmos.azure.com"),
	// 			Replica: &armmongocluster.ReplicationProperties{
	// 				Role: to.Ptr(armmongocluster.ReplicationRolePrimary),
	// 			},
	// 		},
	// 		Location: to.Ptr("westus2"),
	// 	},
	// }
}
