package armhybridkubernetes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ecee919199a39cc0d864410f540aa105bf7cdb64/specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/ListConnectedClusterOperationsExample.json
func ExampleOperationsClient_NewGetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridkubernetes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewGetPager(nil)
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
		// page.OperationList = armhybridkubernetes.OperationList{
		// 	Value: []*armhybridkubernetes.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Kubernetes/connectedClusters/write"),
		// 			Display: &armhybridkubernetes.OperationDisplay{
		// 				Description: to.Ptr("Create any ConnectedCluster"),
		// 				Operation: to.Ptr("Create ConnectedCluster"),
		// 				Provider: to.Ptr("Microsoft.connectedClusters"),
		// 				Resource: to.Ptr("connectedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Kubernetes/connectedClusters/read"),
		// 			Display: &armhybridkubernetes.OperationDisplay{
		// 				Description: to.Ptr("Get any ConnectedCluster"),
		// 				Operation: to.Ptr("Get ConnectedCluster"),
		// 				Provider: to.Ptr("Microsoft.Kubernetes"),
		// 				Resource: to.Ptr("connectedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Kubernetes/connectedClusters/delete"),
		// 			Display: &armhybridkubernetes.OperationDisplay{
		// 				Description: to.Ptr("Delete any ConnectedCluster"),
		// 				Operation: to.Ptr("Delete ConnectedCluster"),
		// 				Provider: to.Ptr("Microsoft.Kubernetes"),
		// 				Resource: to.Ptr("connectedClusters"),
		// 			},
		// 	}},
		// }
	}
}
