package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/NameAvailabilityTest_NameAvailable-POST-example-21.json
func ExampleProfilesClient_CheckTrafficManagerRelativeDNSNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtrafficmanager.NewProfilesClient("<subscription-id>", cred, nil)
	_, err = client.CheckTrafficManagerRelativeDNSNameAvailability(ctx,
		armtrafficmanager.CheckTrafficManagerRelativeDNSNameAvailabilityParameters{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
