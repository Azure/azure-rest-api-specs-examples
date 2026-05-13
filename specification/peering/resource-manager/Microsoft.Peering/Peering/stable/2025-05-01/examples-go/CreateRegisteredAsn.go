package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering/v2"
)

// Generated from example definition: 2025-05-01/CreateRegisteredAsn.json
func ExampleRegisteredAsnsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegisteredAsnsClient().CreateOrUpdate(ctx, "rgName", "peeringName", "registeredAsnName", armpeering.RegisteredAsn{
		Properties: &armpeering.RegisteredAsnProperties{
			Asn: to.Ptr[int32](65000),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpeering.RegisteredAsnsClientCreateOrUpdateResponse{
	// 	RegisteredAsn: &armpeering.RegisteredAsn{
	// 		Name: to.Ptr("registeredAsnName"),
	// 		Type: to.Ptr("Microsoft.Peering/registeredAsns"),
	// 		ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName/registeredAsns/registeredAsnName"),
	// 		Properties: &armpeering.RegisteredAsnProperties{
	// 			Asn: to.Ptr[int32](65000),
	// 			PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
