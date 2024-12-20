package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/Operations_example.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armservicefabricmanagedclusters.OperationListResult{
		// 	Value: []*armservicefabricmanagedclusters.OperationResult{
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceFabric/managedClusters/read"),
		// 			Display: &armservicefabricmanagedclusters.AvailableOperationDisplay{
		// 				Description: to.Ptr("Read any Cluster"),
		// 				Operation: to.Ptr("Read Cluster"),
		// 				Provider: to.Ptr("Microsoft ServiceFabric"),
		// 				Resource: to.Ptr("managedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceFabric/managedClusters/write"),
		// 			Display: &armservicefabricmanagedclusters.AvailableOperationDisplay{
		// 				Description: to.Ptr("Create or Update any Cluster"),
		// 				Operation: to.Ptr("Create or Update Cluster"),
		// 				Provider: to.Ptr("Microsoft ServiceFabric"),
		// 				Resource: to.Ptr("managedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceFabric/managedClusters/delete"),
		// 			Display: &armservicefabricmanagedclusters.AvailableOperationDisplay{
		// 				Description: to.Ptr("Delete any Cluster"),
		// 				Operation: to.Ptr("Delete Cluster"),
		// 				Provider: to.Ptr("Microsoft ServiceFabric"),
		// 				Resource: to.Ptr("managedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceFabric/nodeTypes/read"),
		// 			Display: &armservicefabricmanagedclusters.AvailableOperationDisplay{
		// 				Description: to.Ptr("Read any Node type"),
		// 				Operation: to.Ptr("Read Node type"),
		// 				Provider: to.Ptr("Microsoft ServiceFabric"),
		// 				Resource: to.Ptr("nodeTypes"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceFabric/nodeTypes/write"),
		// 			Display: &armservicefabricmanagedclusters.AvailableOperationDisplay{
		// 				Description: to.Ptr("Create or Update any Node type"),
		// 				Operation: to.Ptr("Create or Update Node type"),
		// 				Provider: to.Ptr("Microsoft ServiceFabric"),
		// 				Resource: to.Ptr("nodeTypes"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceFabric/nodeTypes/delete"),
		// 			Display: &armservicefabricmanagedclusters.AvailableOperationDisplay{
		// 				Description: to.Ptr("Delete any Node type"),
		// 				Operation: to.Ptr("Delete Node type"),
		// 				Provider: to.Ptr("Microsoft ServiceFabric"),
		// 				Resource: to.Ptr("nodeTypes"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceFabric/nodeTypes/restart/action"),
		// 			Display: &armservicefabricmanagedclusters.AvailableOperationDisplay{
		// 				Description: to.Ptr("Restart nodes from Node type"),
		// 				Operation: to.Ptr("Restart nodes from Node type"),
		// 				Provider: to.Ptr("Microsoft ServiceFabric"),
		// 				Resource: to.Ptr("nodeTypes"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceFabric/nodeTypes/reimage/action"),
		// 			Display: &armservicefabricmanagedclusters.AvailableOperationDisplay{
		// 				Description: to.Ptr("Reimage nodes from Node type"),
		// 				Operation: to.Ptr("Reimage nodes from Node type"),
		// 				Provider: to.Ptr("Microsoft ServiceFabric"),
		// 				Resource: to.Ptr("nodeTypes"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceFabric/nodeTypes/deleteNode/action"),
		// 			Display: &armservicefabricmanagedclusters.AvailableOperationDisplay{
		// 				Description: to.Ptr("Delete nodes from Node type"),
		// 				Operation: to.Ptr("Delete nodes from Node type"),
		// 				Provider: to.Ptr("Microsoft ServiceFabric"),
		// 				Resource: to.Ptr("nodeTypes"),
		// 			},
		// 	}},
		// }
	}
}
