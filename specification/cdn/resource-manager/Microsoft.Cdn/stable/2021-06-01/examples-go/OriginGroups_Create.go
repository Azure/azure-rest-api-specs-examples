package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/OriginGroups_Create.json
func ExampleOriginGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewOriginGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"RG",
		"profile1",
		"endpoint1",
		"origingroup1",
		armcdn.OriginGroup{
			Properties: &armcdn.OriginGroupProperties{
				HealthProbeSettings: &armcdn.HealthProbeParameters{
					ProbeIntervalInSeconds: to.Ptr[int32](120),
					ProbePath:              to.Ptr("/health.aspx"),
					ProbeProtocol:          to.Ptr(armcdn.ProbeProtocolHTTP),
					ProbeRequestType:       to.Ptr(armcdn.HealthProbeRequestTypeGET),
				},
				Origins: []*armcdn.ResourceReference{
					{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1"),
					}},
				ResponseBasedOriginErrorDetectionSettings: &armcdn.ResponseBasedOriginErrorDetectionParameters{
					ResponseBasedDetectedErrorTypes:          to.Ptr(armcdn.ResponseBasedDetectedErrorTypesTCPErrorsOnly),
					ResponseBasedFailoverThresholdPercentage: to.Ptr[int32](10),
				},
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
