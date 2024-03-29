package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/NameAvailabilityTest_NameAvailable-POST-example-21.json
func ExampleProfilesClient_CheckTrafficManagerRelativeDNSNameAvailability_nameAvailabilityTestNameAvailablePost21() {
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
		Name: to.Ptr("azsmnet5403"),
		Type: to.Ptr("microsoft.network/trafficmanagerprofiles"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NameAvailability = armtrafficmanager.NameAvailability{
	// 	Name: to.Ptr("azsmnet5403"),
	// 	Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
	// 	NameAvailable: to.Ptr(true),
	// }
}
