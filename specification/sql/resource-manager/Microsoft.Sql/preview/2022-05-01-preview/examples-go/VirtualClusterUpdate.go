package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/VirtualClusterUpdate.json
func ExampleVirtualClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualClustersClient().BeginUpdate(ctx, "testrg", "vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2", armsql.VirtualClusterUpdate{
		Tags: map[string]*string{
			"tkey": to.Ptr("tvalue1"),
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
	// res.VirtualCluster = armsql.VirtualCluster{
	// 	Name: to.Ptr("vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2"),
	// 	Type: to.Ptr("Microsoft.Sql/virtualClusters"),
	// 	ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 		"tkey": to.Ptr("tvalue3"),
	// 	},
	// 	Properties: &armsql.VirtualClusterProperties{
	// 	},
	// }
}
