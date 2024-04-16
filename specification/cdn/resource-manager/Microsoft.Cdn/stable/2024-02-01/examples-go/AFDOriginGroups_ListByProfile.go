package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOriginGroups_ListByProfile.json
func ExampleAFDOriginGroupsClient_NewListByProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAFDOriginGroupsClient().NewListByProfilePager("RG", "profile1", nil)
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
		// page.AFDOriginGroupListResult = armcdn.AFDOriginGroupListResult{
		// 	Value: []*armcdn.AFDOriginGroup{
		// 		{
		// 			Name: to.Ptr("origingroup1"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles/origingroups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1"),
		// 			Properties: &armcdn.AFDOriginGroupProperties{
		// 				HealthProbeSettings: &armcdn.HealthProbeParameters{
		// 					ProbeIntervalInSeconds: to.Ptr[int32](10),
		// 					ProbePath: to.Ptr("/path1"),
		// 					ProbeProtocol: to.Ptr(armcdn.ProbeProtocolHTTP),
		// 					ProbeRequestType: to.Ptr(armcdn.HealthProbeRequestTypeHEAD),
		// 				},
		// 				LoadBalancingSettings: &armcdn.LoadBalancingSettingsParameters{
		// 					AdditionalLatencyInMilliseconds: to.Ptr[int32](1000),
		// 					SampleSize: to.Ptr[int32](3),
		// 					SuccessfulSamplesRequired: to.Ptr[int32](3),
		// 				},
		// 				TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: to.Ptr[int32](5),
		// 				DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 				ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
