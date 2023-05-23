package armqumulo_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/liftrqumulo/armqumulo"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_CreateOrUpdate_MinimumSet_Gen.json
func ExampleFileSystemsClient_BeginCreateOrUpdate_fileSystemsCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSystemsClient().BeginCreateOrUpdate(ctx, "rgopenapi", "aaaaaaaa", armqumulo.FileSystemResource{
		Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		Properties: &armqumulo.FileSystemResourceProperties{
			AdminPassword:     to.Ptr("ekceujoecaashtjlsgcymnrdozk"),
			DelegatedSubnetID: to.Ptr("aaaaaaaaaa"),
			InitialCapacity:   to.Ptr[int32](9),
			MarketplaceDetails: &armqumulo.MarketplaceDetails{
				MarketplaceSubscriptionID:     to.Ptr("aaaaaaaaaaaaa"),
				MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferID:                       to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
				PlanID:                        to.Ptr("aaaaaa"),
				PublisherID:                   to.Ptr("aa"),
			},
			ProvisioningState: to.Ptr(armqumulo.ProvisioningStateAccepted),
			StorageSKU:        to.Ptr(armqumulo.StorageSKUStandard),
			UserDetails: &armqumulo.UserDetails{
				Email: to.Ptr("viptslwulnpaupfljvnjeq"),
			},
		},
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
	// res.FileSystemResource = armqumulo.FileSystemResource{
	// 	ID: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 	Location: to.Ptr("aaaaaaaaa"),
	// 	Properties: &armqumulo.FileSystemResourceProperties{
	// 		DelegatedSubnetID: to.Ptr("aaaaaaaaaa"),
	// 		InitialCapacity: to.Ptr[int32](9),
	// 		MarketplaceDetails: &armqumulo.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 			OfferID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			PlanID: to.Ptr("aaaaaa"),
	// 			PublisherID: to.Ptr("aa"),
	// 		},
	// 		ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
	// 		StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
	// 		UserDetails: &armqumulo.UserDetails{
	// 		},
	// 	},
	// }
}
