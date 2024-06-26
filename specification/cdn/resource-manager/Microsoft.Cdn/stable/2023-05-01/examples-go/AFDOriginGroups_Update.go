package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDOriginGroups_Update.json
func ExampleAFDOriginGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAFDOriginGroupsClient().BeginUpdate(ctx, "RG", "profile1", "origingroup1", armcdn.AFDOriginGroupUpdateParameters{
		Properties: &armcdn.AFDOriginGroupUpdatePropertiesParameters{
			HealthProbeSettings: &armcdn.HealthProbeParameters{
				ProbeIntervalInSeconds: to.Ptr[int32](10),
				ProbePath:              to.Ptr("/path2"),
				ProbeProtocol:          to.Ptr(armcdn.ProbeProtocolNotSet),
				ProbeRequestType:       to.Ptr(armcdn.HealthProbeRequestTypeNotSet),
			},
			LoadBalancingSettings: &armcdn.LoadBalancingSettingsParameters{
				AdditionalLatencyInMilliseconds: to.Ptr[int32](1000),
				SampleSize:                      to.Ptr[int32](3),
				SuccessfulSamplesRequired:       to.Ptr[int32](3),
			},
			TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: to.Ptr[int32](5),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AFDOriginGroup = armcdn.AFDOriginGroup{
	// 	Name: to.Ptr("origingroup1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/origingroups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1"),
	// 	Properties: &armcdn.AFDOriginGroupProperties{
	// 		HealthProbeSettings: &armcdn.HealthProbeParameters{
	// 			ProbeIntervalInSeconds: to.Ptr[int32](10),
	// 			ProbePath: to.Ptr("/path1"),
	// 			ProbeProtocol: to.Ptr(armcdn.ProbeProtocolHTTP),
	// 			ProbeRequestType: to.Ptr(armcdn.HealthProbeRequestTypeHEAD),
	// 		},
	// 		LoadBalancingSettings: &armcdn.LoadBalancingSettingsParameters{
	// 			AdditionalLatencyInMilliseconds: to.Ptr[int32](1000),
	// 			SampleSize: to.Ptr[int32](3),
	// 			SuccessfulSamplesRequired: to.Ptr[int32](3),
	// 		},
	// 		TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: to.Ptr[int32](5),
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 	},
	// }
}
