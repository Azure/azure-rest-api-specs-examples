package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/Endpoints_Update.json
func ExampleEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEndpointsClient().BeginUpdate(ctx, "RG", "profile1", "endpoint1", armcdn.EndpointUpdateParameters{
		Tags: map[string]*string{
			"additionalProperties": to.Ptr("Tag1"),
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
	// res = armcdn.EndpointsClientUpdateResponse{
	// 	Endpoint: armcdn.Endpoint{
	// 		Name: to.Ptr("endpoint1"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/endpoints"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1"),
	// 		Location: to.Ptr("WestCentralUs"),
	// 		Properties: &armcdn.EndpointProperties{
	// 			ContentTypesToCompress: []*string{
	// 			},
	// 			DefaultOriginGroup: &armcdn.ResourceReference{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/originGroups/originGroup1"),
	// 			},
	// 			GeoFilters: []*armcdn.GeoFilter{
	// 			},
	// 			HostName: to.Ptr("endpoint1.azureedge.net"),
	// 			IsCompressionEnabled: to.Ptr(false),
	// 			IsHTTPAllowed: to.Ptr(true),
	// 			IsHTTPSAllowed: to.Ptr(true),
	// 			OriginGroups: []*armcdn.DeepCreatedOriginGroup{
	// 				{
	// 					Name: to.Ptr("originGroup1"),
	// 					Properties: &armcdn.DeepCreatedOriginGroupProperties{
	// 						HealthProbeSettings: &armcdn.HealthProbeParameters{
	// 							ProbeIntervalInSeconds: to.Ptr[int32](120),
	// 							ProbePath: to.Ptr("/health.aspx"),
	// 							ProbeProtocol: to.Ptr(armcdn.ProbeProtocolHTTP),
	// 							ProbeRequestType: to.Ptr(armcdn.HealthProbeRequestTypeGET),
	// 						},
	// 						Origins: []*armcdn.ResourceReference{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/www-bing-com"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Origins: []*armcdn.DeepCreatedOrigin{
	// 				{
	// 					Name: to.Ptr("www-bing-com"),
	// 					Properties: &armcdn.DeepCreatedOriginProperties{
	// 						Enabled: to.Ptr(true),
	// 						HostName: to.Ptr("www.bing.com"),
	// 						HTTPPort: to.Ptr[int32](80),
	// 						HTTPSPort: to.Ptr[int32](443),
	// 						OriginHostHeader: to.Ptr("www.someDomain2.net"),
	// 						Priority: to.Ptr[int32](2),
	// 						Weight: to.Ptr[int32](50),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcdn.EndpointProvisioningStateCreating),
	// 			QueryStringCachingBehavior: to.Ptr(armcdn.QueryStringCachingBehaviorIgnoreQueryString),
	// 			ResourceState: to.Ptr(armcdn.EndpointResourceStateCreating),
	// 		},
	// 		Tags: map[string]*string{
	// 			"additionalProperties": to.Ptr("Tag1"),
	// 		},
	// 	},
	// }
}
