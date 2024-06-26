package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armhybridcontainerservice.OperationListResult{
		// 	Value: []*armhybridcontainerservice.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridContainerService/provisionedClusterInstances/Read"),
		// 			Display: &armhybridcontainerservice.OperationDisplay{
		// 				Description: to.Ptr("Read provisionedClusters"),
		// 				Operation: to.Ptr("Gets/List provisionedClusters resources"),
		// 				Provider: to.Ptr("Microsoft.HybridContainerService"),
		// 				Resource: to.Ptr("provisionedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridContainerService/provisionedClusterInstances/Write"),
		// 			Display: &armhybridcontainerservice.OperationDisplay{
		// 				Description: to.Ptr("Writes provisionedClusters"),
		// 				Operation: to.Ptr("Create/update provisionedClusters resources"),
		// 				Provider: to.Ptr("Microsoft.HybridContainerService"),
		// 				Resource: to.Ptr("provisionedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridContainerService/provisionedClusterInstances/Delete"),
		// 			Display: &armhybridcontainerservice.OperationDisplay{
		// 				Description: to.Ptr("Deletes provisionedClusters"),
		// 				Operation: to.Ptr("Deletes provisionedClusters resource"),
		// 				Provider: to.Ptr("Microsoft.HybridContainerService"),
		// 				Resource: to.Ptr("provisionedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridContainerService/provisionedClusterInstances/upgradeProfiles/Read"),
		// 			Display: &armhybridcontainerservice.OperationDisplay{
		// 				Description: to.Ptr("Gets the upgrade profile of the cluster"),
		// 				Operation: to.Ptr("Get UpgradeProfile"),
		// 				Provider: to.Ptr("Microsoft.HybridContainerService"),
		// 				Resource: to.Ptr("UpgradeProfile"),
		// 			},
		// 	}},
		// }
	}
}
