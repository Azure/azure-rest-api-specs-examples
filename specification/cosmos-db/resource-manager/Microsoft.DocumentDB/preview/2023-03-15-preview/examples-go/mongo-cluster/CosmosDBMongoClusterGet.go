package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/mongo-cluster/CosmosDBMongoClusterGet.json
func ExampleMongoClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMongoClustersClient().Get(ctx, "TestResourceGroup", "myMongoCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MongoCluster = armcosmos.MongoCluster{
	// 	Name: to.Ptr("myMongoCluster"),
	// 	Type: to.Ptr("/Microsoft.DocumentDB/mongoClusters"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myMongoCluster"),
	// 	SystemData: &armcosmos.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armcosmos.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armcosmos.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"additionalProp1": to.Ptr("string"),
	// 		"additionalProp2": to.Ptr("string"),
	// 		"additionalProp3": to.Ptr("string"),
	// 	},
	// 	Properties: &armcosmos.MongoClusterProperties{
	// 		AdministratorLogin: to.Ptr("mongoAdmin"),
	// 		ConnectionString: to.Ptr("mongodb+srv://<user>:<password>@myMongoCluster.mongocluster.cosmos.azure.com"),
	// 		EarliestRestoreTime: to.Ptr("2023-01-13T20:07:35Z"),
	// 		NodeGroupSpecs: []*armcosmos.NodeGroupSpec{
	// 			{
	// 				DiskSizeGB: to.Ptr[int64](128),
	// 				EnableHa: to.Ptr(true),
	// 				SKU: to.Ptr("M30"),
	// 				Kind: to.Ptr(armcosmos.NodeKindShard),
	// 				NodeCount: to.Ptr[int32](4),
	// 		}},
	// 		ProvisioningState: to.Ptr(armcosmos.ProvisioningStateSucceeded),
	// 		ServerVersion: to.Ptr("5.0"),
	// 	},
	// }
}
