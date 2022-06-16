package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Update.json
func ExampleAFDOriginGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewAFDOriginGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"RG",
		"profile1",
		"origingroup1",
		armcdn.AFDOriginGroupUpdateParameters{
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
