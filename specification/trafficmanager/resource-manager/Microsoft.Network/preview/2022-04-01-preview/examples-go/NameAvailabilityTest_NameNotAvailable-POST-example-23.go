package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/NameAvailabilityTest_NameNotAvailable-POST-example-23.json
func ExampleProfilesClient_CheckTrafficManagerRelativeDNSNameAvailability_nameAvailabilityTestNameNotAvailablePost23() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewProfilesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckTrafficManagerRelativeDNSNameAvailability(ctx, armtrafficmanager.CheckTrafficManagerRelativeDNSNameAvailabilityParameters{
		Name: to.Ptr("azsmnet4696"),
		Type: to.Ptr("microsoft.network/trafficmanagerprofiles"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
