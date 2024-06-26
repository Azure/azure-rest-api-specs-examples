package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoPrivateLinkResourcesGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "kustorptest", "kustoCluster", "cluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armkusto.PrivateLinkResource{
	// 	Name: to.Ptr("cluster"),
	// 	Type: to.Ptr("Microsoft.Kusto/clusters/PrivateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/privateLinkResources/cluster"),
	// 	Properties: &armkusto.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("cluster"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("Engine"),
	// 			to.Ptr("DataManagement"),
	// 			to.Ptr("blob-ne3kstrldkustoCluster00"),
	// 			to.Ptr("queue-ne3kstrldkustoCluster00"),
	// 			to.Ptr("table-ne3kstrldkustoCluster00"),
	// 			to.Ptr("blob-qmukstrldkustoCluster01"),
	// 			to.Ptr("queue-qmukstrldkustoCluster01"),
	// 			to.Ptr("table-qmukstrldkustoCluster01")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.westus2.kusto.windows.net"),
	// 				to.Ptr("privatelink.blob.core.windows.net"),
	// 				to.Ptr("privatelink.queue.core.windows.net"),
	// 				to.Ptr("privatelink.table.core.windows.net")},
	// 			},
	// 		}
}
