package armprivatetrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatetrafficmanager/armprivatetrafficmanager"
)

// Generated from example definition: 2026-10-01-preview/Profiles_ListByResourceGroup_MaximumSet_Gen.json
func ExampleProfilesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatetrafficmanager.NewClientFactory("10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProfilesClient().NewListByResourceGroupPager("rgprivateTrafficManager", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armprivatetrafficmanager.ProfilesClientListByResourceGroupResponse{
		// 	ProfileListResult: armprivatetrafficmanager.ProfileListResult{
		// 		Value: []*armprivatetrafficmanager.Profile{
		// 			{
		// 				Properties: &armprivatetrafficmanager.ProfileProperties{
		// 					CustomTopologyMapMode: to.Ptr(armprivatetrafficmanager.CustomTopologyMapModeDisabled),
		// 					DNSConfig: &armprivatetrafficmanager.DNSConfig{
		// 						RecordType: to.Ptr(armprivatetrafficmanager.RecordTypeA),
		// 						TTL: to.Ptr[int64](10),
		// 					},
		// 					TopologyMapID: to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/topologyMaps/myTopologyMap"),
		// 					ProfileStatus: to.Ptr(armprivatetrafficmanager.ProfileStatusEnabled),
		// 					TrafficRoutingMethod: to.Ptr(armprivatetrafficmanager.TrafficRoutingMethodPriority),
		// 					Endpoints: []*armprivatetrafficmanager.ProfileEndpoint{
		// 						{
		// 							Target: to.Ptr("10.0.0.1"),
		// 							MonitoringTarget: to.Ptr("10.0.0.1"),
		// 							EndpointStatus: to.Ptr(armprivatetrafficmanager.AdministrativeStatusEnabled),
		// 							Kind: to.Ptr(armprivatetrafficmanager.EndpointsKindEndpoint),
		// 							Weight: to.Ptr[int64](100),
		// 							Priority: to.Ptr[int64](10),
		// 							AlwaysServe: to.Ptr(armprivatetrafficmanager.AlwaysServeEnabled),
		// 							HealthPolicyID: to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/privateTrafficManagerProfiles/myProfile/healthPolicies/myHealthPolicy"),
		// 							ProvisioningState: to.Ptr(armprivatetrafficmanager.ProvisioningStateSucceeded),
		// 							Name: to.Ptr("jihmlhsjltkcrdgogtpolvydeqa"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armprivatetrafficmanager.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 				},
		// 				Location: to.Ptr("global"),
		// 				ID: to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/privateTrafficManagerProfiles/myProfile"),
		// 				Name: to.Ptr("myProfile"),
		// 				Type: to.Ptr("Microsoft.Network/privateTrafficManagerProfiles"),
		// 				SystemData: &armprivatetrafficmanager.SystemData{
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armprivatetrafficmanager.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.February, 18, 10, 5, 58, 138000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armprivatetrafficmanager.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.February, 18, 10, 5, 58, 139000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/privateTrafficManagerProfiles?api-version=2026-10-01-preview&$skipToken=eyJuZXh0UGFnZSI6IjIifQ"),
		// 	},
		// }
	}
}
