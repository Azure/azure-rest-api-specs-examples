package armprivatetrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatetrafficmanager/armprivatetrafficmanager"
)

// Generated from example definition: 2026-10-01-preview/Endpoints_CreateOrUpdate_MaximumSet_Gen.json
func ExampleEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatetrafficmanager.NewClientFactory("10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEndpointsClient().BeginCreateOrUpdate(ctx, "rgprivateTrafficManager", "myProfile", "myEndpoint", armprivatetrafficmanager.Endpoint{
		Properties: &armprivatetrafficmanager.EndpointProperties{
			Target:           to.Ptr("10.0.0.1"),
			MonitoringTarget: to.Ptr("10.0.0.1"),
			EndpointStatus:   to.Ptr(armprivatetrafficmanager.AdministrativeStatusEnabled),
			Kind:             to.Ptr(armprivatetrafficmanager.EndpointsKindEndpoint),
			Weight:           to.Ptr[int64](100),
			Priority:         to.Ptr[int64](10),
			AlwaysServe:      to.Ptr(armprivatetrafficmanager.AlwaysServeEnabled),
			HealthPolicyID:   to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/privateTrafficManagerProfiles/myProfile/healthPolicies/myHealthPolicy"),
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
	// res = armprivatetrafficmanager.EndpointsClientCreateOrUpdateResponse{
	// 	Endpoint: armprivatetrafficmanager.Endpoint{
	// 		Properties: &armprivatetrafficmanager.EndpointProperties{
	// 			Target: to.Ptr("10.0.0.1"),
	// 			MonitoringTarget: to.Ptr("10.0.0.1"),
	// 			EndpointStatus: to.Ptr(armprivatetrafficmanager.AdministrativeStatusEnabled),
	// 			Kind: to.Ptr(armprivatetrafficmanager.EndpointsKindEndpoint),
	// 			Weight: to.Ptr[int64](100),
	// 			Priority: to.Ptr[int64](10),
	// 			AlwaysServe: to.Ptr(armprivatetrafficmanager.AlwaysServeEnabled),
	// 			HealthPolicyID: to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/privateTrafficManagerProfiles/myProfile/healthPolicies/myHealthPolicy"),
	// 			ProvisioningState: to.Ptr(armprivatetrafficmanager.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/privateTrafficManagerProfiles/myProfile/endpoints/myEndpoint"),
	// 		Name: to.Ptr("myEndpoint"),
	// 		Type: to.Ptr("Microsoft.Network/privateTrafficManagerProfiles/endpoints"),
	// 		SystemData: &armprivatetrafficmanager.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armprivatetrafficmanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.February, 18, 10, 5, 58, 138000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armprivatetrafficmanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.February, 18, 10, 5, 58, 139000000, time.UTC)),
	// 		},
	// 	},
	// }
}
