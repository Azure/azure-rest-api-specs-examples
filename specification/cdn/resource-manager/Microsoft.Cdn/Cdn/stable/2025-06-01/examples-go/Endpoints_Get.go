package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/Endpoints_Get.json
func ExampleEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().Get(ctx, "RG", "profile1", "endpoint1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcdn.EndpointsClientGetResponse{
	// 	Endpoint: armcdn.Endpoint{
	// 		Name: to.Ptr("endpoint1"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/endpoints"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1"),
	// 		Location: to.Ptr("CentralUs"),
	// 		Properties: &armcdn.EndpointProperties{
	// 			ContentTypesToCompress: []*string{
	// 			},
	// 			CustomDomains: []*armcdn.DeepCreatedCustomDomain{
	// 				{
	// 					Name: to.Ptr("www-someDomain-net"),
	// 					Properties: &armcdn.DeepCreatedCustomDomainProperties{
	// 						HostName: to.Ptr("www.someDomain.Net"),
	// 					},
	// 				},
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
	// 			OptimizationType: to.Ptr(armcdn.OptimizationTypeDynamicSiteAcceleration),
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
	// 			OriginHostHeader: to.Ptr("www.bing.com"),
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
	// 			ProbePath: to.Ptr("/image"),
	// 			ProvisioningState: to.Ptr(armcdn.EndpointProvisioningStateSucceeded),
	// 			QueryStringCachingBehavior: to.Ptr(armcdn.QueryStringCachingBehaviorNotSet),
	// 			ResourceState: to.Ptr(armcdn.EndpointResourceStateRunning),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
