package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-03-01-preview/ApplicationTypeNamePutOperation_example.json
func ExampleApplicationTypesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationTypesClient().CreateOrUpdate(ctx, "resRg", "myCluster", "myAppType", armservicefabricmanagedclusters.ApplicationTypeResource{
		Location: to.Ptr("eastus"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabricmanagedclusters.ApplicationTypesClientCreateOrUpdateResponse{
	// 	ApplicationTypeResource: &armservicefabricmanagedclusters.ApplicationTypeResource{
	// 		Name: to.Ptr("myAppType"),
	// 		Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applicationTypes"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType"),
	// 		Properties: &armservicefabricmanagedclusters.ApplicationTypeResourceProperties{
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
