package armdellstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dell/armdellstorage"
)

// Generated from example definition: 2025-03-21/FileSystems_ListByResourceGroup_MinimumSet_Gen.json
func ExampleFileSystemsClient_NewListByResourceGroupPager_fileSystemsListByResourceGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdellstorage.NewClientFactory("BF7E7352-2FE4-4163-9CF7-5FF8EC2E9B92", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListByResourceGroupPager("rgDell", nil)
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
		// page = armdellstorage.FileSystemsClientListByResourceGroupResponse{
		// 	FileSystemResourceListResult: armdellstorage.FileSystemResourceListResult{
		// 		Value: []*armdellstorage.FileSystemResource{
		// 			{
		// 				Properties: &armdellstorage.FileSystemResourceProperties{
		// 					Capacity: &armdellstorage.Capacity{
		// 						Min: to.Ptr("1"),
		// 						Max: to.Ptr("10"),
		// 						Incremental: to.Ptr("1"),
		// 						Current: to.Ptr("5"),
		// 					},
		// 					Marketplace: &armdellstorage.MarketplaceDetails{
		// 						PlanID: to.Ptr("lgozf"),
		// 						OfferID: to.Ptr("pzhjvibxqgeqkndqnjlduwnxqbr"),
		// 						PrivateOfferID: to.Ptr("privateOfferId"),
		// 						PlanName: to.Ptr("planeName"),
		// 						EndDate: to.Ptr("2023-05-27T17:00:00-07:00"),
		// 					},
		// 					DelegatedSubnetID: to.Ptr("yp"),
		// 					DelegatedSubnetCidr: to.Ptr("10.0.0.1/24"),
		// 					User: &armdellstorage.UserDetails{
		// 					},
		// 					FileSystemID: to.Ptr("filesystemId"),
		// 					DellReferenceNumber: to.Ptr("fhewkj"),
		// 					Encryption: &armdellstorage.EncryptionProperties{
		// 						EncryptionType: to.Ptr(armdellstorage.ResourceEncryptionTypeMicrosoftManagedKeysMMK),
		// 					},
		// 				},
		// 				Location: to.Ptr("tbcvhxzpgrijtdygkttnfswwtacs"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
