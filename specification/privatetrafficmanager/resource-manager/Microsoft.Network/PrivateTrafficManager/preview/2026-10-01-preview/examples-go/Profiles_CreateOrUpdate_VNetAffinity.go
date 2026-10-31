package armprivatetrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatetrafficmanager/armprivatetrafficmanager"
)

// Generated from example definition: 2026-10-01-preview/Profiles_CreateOrUpdate_VNetAffinity.json
func ExampleProfilesClient_BeginCreateOrUpdate_profilesCreateOrUpdateVnetAffinity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatetrafficmanager.NewClientFactory("10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginCreateOrUpdate(ctx, "rgprivateTrafficManager", "vnet-affinity-profile", armprivatetrafficmanager.Profile{
		Location: to.Ptr("global"),
		Properties: &armprivatetrafficmanager.ProfileProperties{
			CustomTopologyMapMode: to.Ptr(armprivatetrafficmanager.CustomTopologyMapModeDisabled),
			ProfileStatus:         to.Ptr(armprivatetrafficmanager.ProfileStatusEnabled),
			TrafficRoutingMethod:  to.Ptr(armprivatetrafficmanager.TrafficRoutingMethodVNet),
			Endpoints: []*armprivatetrafficmanager.ProfileEndpoint{
				{
					Name:           to.Ptr("matched-endpoint"),
					Target:         to.Ptr("10.0.0.10"),
					EndpointStatus: to.Ptr(armprivatetrafficmanager.AdministrativeStatusEnabled),
					VirtualNetworkIDs: []*string{
						to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/virtualNetworks/vnet-east"),
					},
				},
				{
					Name:           to.Ptr("fallback-endpoint"),
					Target:         to.Ptr("10.0.0.20"),
					EndpointStatus: to.Ptr(armprivatetrafficmanager.AdministrativeStatusEnabled),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armprivatetrafficmanager.ProfilesClientCreateOrUpdateResponse{
	// 	Profile: armprivatetrafficmanager.Profile{
	// 		ID: to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/privateTrafficManagerProfiles/vnet-affinity-profile"),
	// 		Name: to.Ptr("vnet-affinity-profile"),
	// 		Type: to.Ptr("Microsoft.Network/privateTrafficManagerProfiles"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armprivatetrafficmanager.ProfileProperties{
	// 			CustomTopologyMapMode: to.Ptr(armprivatetrafficmanager.CustomTopologyMapModeDisabled),
	// 			ProfileStatus: to.Ptr(armprivatetrafficmanager.ProfileStatusEnabled),
	// 			TrafficRoutingMethod: to.Ptr(armprivatetrafficmanager.TrafficRoutingMethodVNet),
	// 			Endpoints: []*armprivatetrafficmanager.ProfileEndpoint{
	// 				{
	// 					Name: to.Ptr("matched-endpoint"),
	// 					Target: to.Ptr("10.0.0.10"),
	// 					EndpointStatus: to.Ptr(armprivatetrafficmanager.AdministrativeStatusEnabled),
	// 					VirtualNetworkIDs: []*string{
	// 						to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/virtualNetworks/vnet-east"),
	// 					},
	// 					ProvisioningState: to.Ptr(armprivatetrafficmanager.ProvisioningStateSucceeded),
	// 				},
	// 				{
	// 					Name: to.Ptr("fallback-endpoint"),
	// 					Target: to.Ptr("10.0.0.20"),
	// 					EndpointStatus: to.Ptr(armprivatetrafficmanager.AdministrativeStatusEnabled),
	// 					ProvisioningState: to.Ptr(armprivatetrafficmanager.ProvisioningStateSucceeded),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armprivatetrafficmanager.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
