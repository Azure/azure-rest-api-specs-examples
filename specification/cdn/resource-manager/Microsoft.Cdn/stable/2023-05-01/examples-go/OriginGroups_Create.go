package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/OriginGroups_Create.json
func ExampleOriginGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOriginGroupsClient().BeginCreate(ctx, "RG", "profile1", "endpoint1", "origingroup1", armcdn.OriginGroup{
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
	// res.OriginGroup = armcdn.OriginGroup{
	// 	Name: to.Ptr("origingroup1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/endpoints/origingroups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/originGroups/originGroup1"),
	// 	Properties: &armcdn.OriginGroupProperties{
	// 		HealthProbeSettings: &armcdn.HealthProbeParameters{
	// 			ProbeIntervalInSeconds: to.Ptr[int32](120),
	// 			ProbePath: to.Ptr("/health.aspx"),
	// 			ProbeProtocol: to.Ptr(armcdn.ProbeProtocolHTTP),
	// 			ProbeRequestType: to.Ptr(armcdn.HealthProbeRequestTypeGET),
	// 		},
	// 		Origins: []*armcdn.ResourceReference{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1"),
	// 		}},
	// 		ResponseBasedOriginErrorDetectionSettings: &armcdn.ResponseBasedOriginErrorDetectionParameters{
	// 			ResponseBasedDetectedErrorTypes: to.Ptr(armcdn.ResponseBasedDetectedErrorTypesTCPErrorsOnly),
	// 			ResponseBasedFailoverThresholdPercentage: to.Ptr[int32](10),
	// 		},
	// 		ProvisioningState: to.Ptr(armcdn.OriginGroupProvisioningStateSucceeded),
	// 		ResourceState: to.Ptr(armcdn.OriginGroupResourceStateActive),
	// 	},
	// }
}
