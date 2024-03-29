package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ListOperations.json
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
		// page.ResourceProviderOperationList = armhybridcontainerservice.ResourceProviderOperationList{
		// 	Value: []*armhybridcontainerservice.ResourceProviderOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridContainerService/provisionedClusters/Read"),
		// 			Display: &armhybridcontainerservice.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Read provisionedClusters"),
		// 				Operation: to.Ptr("Gets/List provisionedClusters resources"),
		// 				Provider: to.Ptr("Microsoft.HybridContainerService"),
		// 				Resource: to.Ptr("provisionedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridContainerService/provisionedClusters/Write"),
		// 			Display: &armhybridcontainerservice.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Writes provisionedClusters"),
		// 				Operation: to.Ptr("Create/update provisionedClusters resources"),
		// 				Provider: to.Ptr("Microsoft.HybridContainerService"),
		// 				Resource: to.Ptr("provisionedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridContainerService/provisionedClusters/Delete"),
		// 			Display: &armhybridcontainerservice.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes provisionedClusters"),
		// 				Operation: to.Ptr("Deletes provisionedClusters resource"),
		// 				Provider: to.Ptr("Microsoft.HybridContainerService"),
		// 				Resource: to.Ptr("provisionedClusters"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridContainerService/provisionedClusters/upgradeProfiles/Read"),
		// 			Display: &armhybridcontainerservice.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets the upgrade profile of the cluster"),
		// 				Operation: to.Ptr("Get UpgradeProfile"),
		// 				Provider: to.Ptr("Microsoft.HybridContainerService"),
		// 				Resource: to.Ptr("UpgradeProfile"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridContainerService/provisionedClusters/upgradeNodeImageVersionForEntireCluster/write"),
		// 			Display: &armhybridcontainerservice.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Upgrade the node image version of the entire cluster"),
		// 				Operation: to.Ptr("Upgrade entire cluster node image version"),
		// 				Provider: to.Ptr("Microsoft.HybridContainerService"),
		// 				Resource: to.Ptr("provisionedClusters"),
		// 			},
		// 	}},
		// }
	}
}
