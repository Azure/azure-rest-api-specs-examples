package armqumulo_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/liftrqumulo/armqumulo/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72de08114673a547de8a017c85ed89a2017a86f7/specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/FileSystems_Update_MinimumSet_Gen.json
func ExampleFileSystemsClient_Update_fileSystemsUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSystemsClient().Update(ctx, "rgQumulo", "aaaaaaaaaaaaaaaaa", armqumulo.FileSystemResourceUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileSystemResource = armqumulo.FileSystemResource{
	// 	Name: to.Ptr("aaaaa"),
	// 	Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	Properties: &armqumulo.FileSystemResourceProperties{
	// 		DelegatedSubnetID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		MarketplaceDetails: &armqumulo.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 			OfferID: to.Ptr("aaaaaaaaa"),
	// 			PlanID: to.Ptr("aaaaaaa"),
	// 		},
	// 		ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
	// 		StorageSKU: to.Ptr("Standard"),
	// 		UserDetails: &armqumulo.UserDetails{
	// 		},
	// 	},
	// }
}
