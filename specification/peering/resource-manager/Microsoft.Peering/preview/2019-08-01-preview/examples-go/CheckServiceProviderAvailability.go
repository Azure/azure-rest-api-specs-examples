package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/CheckServiceProviderAvailability.json
func ExampleManagementClient_CheckServiceProviderAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewManagementClient("<subscription-id>", cred, nil)
	res, err := client.CheckServiceProviderAvailability(ctx,
		armpeering.CheckServiceProviderAvailabilityInput{
			PeeringServiceLocation: to.StringPtr("<peering-service-location>"),
			PeeringServiceProvider: to.StringPtr("<peering-service-provider>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ManagementClientCheckServiceProviderAvailabilityResult)
}
