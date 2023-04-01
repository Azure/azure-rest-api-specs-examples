package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/abd5d0016f12f6862cae88ef70f1333e84e20c07/specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/NameAvailabilityTest_NameNotAvailable-POST-example-23.json
func ExampleProfilesClient_CheckTrafficManagerRelativeDNSNameAvailability_nameAvailabilityTestNameNotAvailablePost23() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().CheckTrafficManagerRelativeDNSNameAvailability(ctx, armtrafficmanager.CheckTrafficManagerRelativeDNSNameAvailabilityParameters{
		Name: to.Ptr("azsmnet4696"),
		Type: to.Ptr("microsoft.network/trafficmanagerprofiles"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NameAvailability = armtrafficmanager.NameAvailability{
	// 	Name: to.Ptr("azsmnet4696"),
	// 	Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
	// 	Message: to.Ptr("Domain name azsmnet4696.tmpreview.watmtest.azure-test.net already exists. Please choose a different DNS prefix."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr("AlreadyExists"),
	// }
}
