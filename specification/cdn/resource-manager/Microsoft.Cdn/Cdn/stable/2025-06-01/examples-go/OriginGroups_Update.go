package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/OriginGroups_Update.json
func ExampleOriginGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOriginGroupsClient().BeginUpdate(ctx, "RG", "profile1", "endpoint1", "originGroup1", armcdn.OriginGroupUpdateParameters{
		Properties: &armcdn.OriginGroupUpdatePropertiesParameters{
			HealthProbeSettings: &armcdn.HealthProbeParameters{
				ProbeIntervalInSeconds: to.Ptr[int32](120),
				ProbePath:              to.Ptr("/health.aspx"),
				ProbeProtocol:          to.Ptr(armcdn.ProbeProtocolHTTP),
				ProbeRequestType:       to.Ptr(armcdn.HealthProbeRequestTypeGET),
			},
			Origins: []*armcdn.ResourceReference{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin2"),
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
	// res = armcdn.OriginGroupsClientUpdateResponse{
	// 	OriginGroup: armcdn.OriginGroup{
	// 		Name: to.Ptr("www-someDomain-net"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/endpoints/origins"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/www-someDomain-net"),
	// 		Properties: &armcdn.OriginGroupProperties{
	// 			HealthProbeSettings: &armcdn.HealthProbeParameters{
	// 				ProbeIntervalInSeconds: to.Ptr[int32](120),
	// 				ProbePath: to.Ptr("/health.aspx"),
	// 				ProbeProtocol: to.Ptr(armcdn.ProbeProtocolHTTP),
	// 				ProbeRequestType: to.Ptr(armcdn.HealthProbeRequestTypeGET),
	// 			},
	// 			Origins: []*armcdn.ResourceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin2"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcdn.OriginGroupProvisioningStateSucceeded),
	// 			ResourceState: to.Ptr(armcdn.OriginGroupResourceStateActive),
	// 		},
	// 	},
	// }
}
