package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-11-15-preview/examples/CosmosDBMongoDBCollectionThroughputUpdate.json
func ExampleMongoDBResourcesClient_BeginUpdateMongoDBCollectionThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMongoDBResourcesClient().BeginUpdateMongoDBCollectionThroughput(ctx, "rg1", "ddb1", "databaseName", "collectionName", armcosmos.ThroughputSettingsUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ThroughputSettingsUpdateProperties{
			Resource: &armcosmos.ThroughputSettingsResource{
				Throughput: to.Ptr[int32](400),
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
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/mongodbCollections/throughputSettings"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/mongodbDatabases/databaseName/mongodbCollections/collectionName/throughputSettings/default"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}
