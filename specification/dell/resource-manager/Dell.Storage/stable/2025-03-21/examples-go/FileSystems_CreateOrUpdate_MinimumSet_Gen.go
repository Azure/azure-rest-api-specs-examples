package armdellstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dell/armdellstorage"
)

// Generated from example definition: 2025-03-21/FileSystems_CreateOrUpdate_MinimumSet_Gen.json
func ExampleFileSystemsClient_BeginCreateOrUpdate_fileSystemsCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdellstorage.NewClientFactory("BF7E7352-2FE4-4163-9CF7-5FF8EC2E9B92", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSystemsClient().BeginCreateOrUpdate(ctx, "rgDell", "abcd", armdellstorage.FileSystemResource{
		Properties: &armdellstorage.FileSystemResourceProperties{
			Marketplace: &armdellstorage.MarketplaceDetails{
				PlanID:         to.Ptr("lgozf"),
				OfferID:        to.Ptr("pzhjvibxqgeqkndqnjlduwnxqbr"),
				PrivateOfferID: to.Ptr("privateOfferId"),
				PlanName:       to.Ptr("planeName"),
			},
			DelegatedSubnetID:   to.Ptr("yp"),
			DelegatedSubnetCidr: to.Ptr("10.0.0.1/24"),
			User: &armdellstorage.UserDetails{
				Email: to.Ptr("hoznewwtzmyjzctzosfuh"),
			},
			DellReferenceNumber: to.Ptr("fhewkj"),
			Encryption: &armdellstorage.EncryptionProperties{
				EncryptionType: to.Ptr(armdellstorage.ResourceEncryptionTypeMicrosoftManagedKeysMMK),
			},
		},
		Location: to.Ptr("tbcvhxzpgrijtdygkttnfswwtacs"),
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
	// res = armdellstorage.FileSystemsClientCreateOrUpdateResponse{
	// 	FileSystemResource: &armdellstorage.FileSystemResource{
	// 		Properties: &armdellstorage.FileSystemResourceProperties{
	// 			Capacity: &armdellstorage.Capacity{
	// 				Min: to.Ptr("1"),
	// 				Max: to.Ptr("10"),
	// 				Incremental: to.Ptr("1"),
	// 				Current: to.Ptr("5"),
	// 			},
	// 			Marketplace: &armdellstorage.MarketplaceDetails{
	// 				PlanID: to.Ptr("lgozf"),
	// 				OfferID: to.Ptr("pzhjvibxqgeqkndqnjlduwnxqbr"),
	// 				PrivateOfferID: to.Ptr("privateOfferId"),
	// 				PlanName: to.Ptr("planeName"),
	// 				EndDate: to.Ptr("2023-05-27T17:00:00-07:00"),
	// 				TermUnit: to.Ptr("P1Y"),
	// 			},
	// 			FileSystemID: to.Ptr("filesystemId"),
	// 			DelegatedSubnetID: to.Ptr("yp"),
	// 			DelegatedSubnetCidr: to.Ptr("10.0.0.1/24"),
	// 			User: &armdellstorage.UserDetails{
	// 			},
	// 			DellReferenceNumber: to.Ptr("fhewkj"),
	// 			Encryption: &armdellstorage.EncryptionProperties{
	// 				EncryptionType: to.Ptr(armdellstorage.ResourceEncryptionTypeMicrosoftManagedKeysMMK),
	// 			},
	// 		},
	// 		Location: to.Ptr("tbcvhxzpgrijtdygkttnfswwtacs"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 	},
	// }
}
