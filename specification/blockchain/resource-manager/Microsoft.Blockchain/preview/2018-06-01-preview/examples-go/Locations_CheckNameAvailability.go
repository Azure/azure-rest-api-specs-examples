package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/Locations_CheckNameAvailability.json
func ExampleLocationsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().CheckNameAvailability(ctx, "southeastasia", &armblockchain.LocationsClientCheckNameAvailabilityOptions{NameAvailabilityRequest: &armblockchain.NameAvailabilityRequest{
		Name: to.Ptr("contosemember1"),
		Type: to.Ptr("Microsoft.Blockchain/blockchainMembers"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NameAvailability = armblockchain.NameAvailability{
	// 	Message: to.Ptr("A blockchain member named 'contosemember1' is already in use."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armblockchain.NameAvailabilityReasonAlreadyExists),
	// }
}
