package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/CreatePeeringService.json
func ExampleServicesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewServicesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<peering-service-name>",
		armpeering.Service{
			Location: to.StringPtr("<location>"),
			Properties: &armpeering.ServiceProperties{
				PeeringServiceLocation:         to.StringPtr("<peering-service-location>"),
				PeeringServiceProvider:         to.StringPtr("<peering-service-provider>"),
				ProviderBackupPeeringLocation:  to.StringPtr("<provider-backup-peering-location>"),
				ProviderPrimaryPeeringLocation: to.StringPtr("<provider-primary-peering-location>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientCreateOrUpdateResult)
}
