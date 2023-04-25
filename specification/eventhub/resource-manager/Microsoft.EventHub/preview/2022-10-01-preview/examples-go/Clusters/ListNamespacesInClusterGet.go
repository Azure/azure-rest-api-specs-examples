package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/Clusters/ListNamespacesInClusterGet.json
func ExampleClustersClient_ListNamespaces() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().ListNamespaces(ctx, "myResourceGroup", "testCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EHNamespaceIDListResult = armeventhub.EHNamespaceIDListResult{
	// 	Value: []*armeventhub.EHNamespaceIDContainer{
	// 		{
	// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-EventHub-SouthCentralUS/providers/Microsoft.EventHub/namespaces/rrama-int7-ns1"),
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-EventHub-SouthCentralUS/providers/Microsoft.EventHub/namespaces/rrama-ehns2-int7"),
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.EventHub/namespaces/db3-rrama-foo1"),
	// 	}},
	// }
}
