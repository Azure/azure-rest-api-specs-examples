package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/OriginGroups_ListByEndpoint.json
func ExampleOriginGroupsClient_NewListByEndpointPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOriginGroupsClient().NewListByEndpointPager("RG", "profile1", "endpoint1", nil)
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
		// page = armcdn.OriginGroupsClientListByEndpointResponse{
		// 	OriginGroupListResult: armcdn.OriginGroupListResult{
		// 		Value: []*armcdn.OriginGroup{
		// 			{
		// 				Name: to.Ptr("origingroup1"),
		// 				Type: to.Ptr("Microsoft.Cdn/profiles/endpoints/origingroups"),
		// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/originGroups/originGroup1"),
		// 				Properties: &armcdn.OriginGroupProperties{
		// 					HealthProbeSettings: &armcdn.HealthProbeParameters{
		// 						ProbeIntervalInSeconds: to.Ptr[int32](120),
		// 						ProbePath: to.Ptr("/health.aspx"),
		// 						ProbeProtocol: to.Ptr(armcdn.ProbeProtocolHTTP),
		// 						ProbeRequestType: to.Ptr(armcdn.HealthProbeRequestTypeGET),
		// 					},
		// 					Origins: []*armcdn.ResourceReference{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armcdn.OriginGroupProvisioningStateSucceeded),
		// 					ResourceState: to.Ptr(armcdn.OriginGroupResourceStateActive),
		// 					ResponseBasedOriginErrorDetectionSettings: &armcdn.ResponseBasedOriginErrorDetectionParameters{
		// 						ResponseBasedDetectedErrorTypes: to.Ptr(armcdn.ResponseBasedDetectedErrorTypesTCPErrorsOnly),
		// 						ResponseBasedFailoverThresholdPercentage: to.Ptr[int32](10),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
