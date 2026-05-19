package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/AFDOriginGroups_ListByProfile.json
func ExampleAFDOriginGroupsClient_NewListByProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armcdn.AFDOriginGroupsClientListByProfileResponse{
		// 	AFDOriginGroupListResult: armcdn.AFDOriginGroupListResult{
		// 		Value: []*armcdn.AFDOriginGroup{
		// 			{
		// 				Name: to.Ptr("origingroup1"),
		// 				Type: to.Ptr("Microsoft.Cdn/profiles/origingroups"),
		// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1"),
		// 				Properties: &armcdn.AFDOriginGroupProperties{
		// 					Authentication: &armcdn.OriginAuthenticationProperties{
		// 						Type: to.Ptr(armcdn.OriginAuthenticationTypeUserAssignedIdentity),
		// 						Scope: to.Ptr("https://www.contoso.com/.default"),
		// 						UserAssignedIdentity: &armcdn.ResourceReference{
		// 							ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-assigned-id-1"),
		// 						},
		// 					},
		// 					DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 					HealthProbeSettings: &armcdn.HealthProbeParameters{
		// 						ProbeIntervalInSeconds: to.Ptr[int32](10),
		// 						ProbePath: to.Ptr("/path1"),
		// 						ProbeProtocol: to.Ptr(armcdn.ProbeProtocolHTTP),
		// 						ProbeRequestType: to.Ptr(armcdn.HealthProbeRequestTypeHEAD),
		// 					},
		// 					LoadBalancingSettings: &armcdn.LoadBalancingSettingsParameters{
		// 						AdditionalLatencyInMilliseconds: to.Ptr[int32](1000),
		// 						SampleSize: to.Ptr[int32](3),
		// 						SuccessfulSamplesRequired: to.Ptr[int32](3),
		// 					},
		// 					ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
		// 					TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: to.Ptr[int32](5),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
