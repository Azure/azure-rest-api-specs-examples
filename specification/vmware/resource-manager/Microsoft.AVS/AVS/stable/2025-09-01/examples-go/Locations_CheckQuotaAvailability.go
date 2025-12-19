package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/Locations_CheckQuotaAvailability.json
func ExampleLocationsClient_CheckQuotaAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().CheckQuotaAvailability(ctx, "eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.LocationsClientCheckQuotaAvailabilityResponse{
	// 	Quota: &armavs.Quota{
	// 		HostsRemaining: map[string]*int32{
	// 			"AV20": to.Ptr[int32](0),
	// 			"AV36": to.Ptr[int32](999),
	// 		},
	// 		QuotaEnabled: to.Ptr(armavs.QuotaEnabledEnabled),
	// 	},
	// }
}
