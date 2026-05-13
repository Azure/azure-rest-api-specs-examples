package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering/v2"
)

// Generated from example definition: 2025-05-01/GetPeeringServicePrefix.json
func ExamplePrefixesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrefixesClient().Get(ctx, "rgName", "peeringServiceName", "peeringServicePrefixName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpeering.PrefixesClientGetResponse{
	// 	ServicePrefix: &armpeering.ServicePrefix{
	// 		Name: to.Ptr("peeringServicePrefixName"),
	// 		ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peeringServices/peeringServiceName/prefixes/peeringServicePrefixName"),
	// 		Properties: &armpeering.ServicePrefixProperties{
	// 			LearnedType: to.Ptr(armpeering.LearnedTypeViaServiceProvider),
	// 			PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Prefix: to.Ptr("192.168.1.0/24"),
	// 			PrefixValidationState: to.Ptr(armpeering.PrefixValidationStateVerified),
	// 			ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
