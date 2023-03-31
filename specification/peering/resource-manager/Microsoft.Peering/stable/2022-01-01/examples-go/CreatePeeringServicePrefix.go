package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/CreatePeeringServicePrefix.json
func ExamplePrefixesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrefixesClient().CreateOrUpdate(ctx, "rgName", "peeringServiceName", "peeringServicePrefixName", armpeering.ServicePrefix{
		Properties: &armpeering.ServicePrefixProperties{
			PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
			Prefix:                  to.Ptr("192.168.1.0/24"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServicePrefix = armpeering.ServicePrefix{
	// 	Name: to.Ptr("peeringServicePrefixName"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peeringServices/peeringServiceName/prefixes/peeringServicePrefixName"),
	// 	Properties: &armpeering.ServicePrefixProperties{
	// 		ErrorMessage: to.Ptr("Prefix is not announced by the service provider to Microsoft."),
	// 		LearnedType: to.Ptr(armpeering.LearnedTypeNone),
	// 		PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Prefix: to.Ptr("192.168.1.0/24"),
	// 		PrefixValidationState: to.Ptr(armpeering.PrefixValidationStateFailed),
	// 		ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 	},
	// }
}
