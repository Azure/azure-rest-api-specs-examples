package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/VirtualClusterUpdate.json
func ExampleVirtualClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
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
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.VirtualClustersClientUpdateResponse{
	// 	VirtualCluster: armsql.VirtualCluster{
	// 		Name: to.Ptr("vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2"),
	// 		Type: to.Ptr("Microsoft.Sql/virtualClusters"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2"),
	// 		Location: to.Ptr("japaneast"),
	// 		Properties: &armsql.VirtualClusterProperties{
	// 		},
	// 		Tags: map[string]*string{
	// 			"tkey": to.Ptr("tvalue3"),
	// 		},
	// 	},
	// }
}
