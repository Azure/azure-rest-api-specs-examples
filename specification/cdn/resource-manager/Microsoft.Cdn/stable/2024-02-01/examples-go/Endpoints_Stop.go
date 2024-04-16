package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Endpoints_Stop.json
func ExampleEndpointsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEndpointsClient().BeginStop(ctx, "RG", "profile1", "endpoint1", nil)
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
	// res.Endpoint = armcdn.Endpoint{
	// 	Name: to.Ptr("endpoint4899"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/endpoints"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1"),
	// 	Location: to.Ptr("WestUs"),
	// 	Tags: map[string]*string{
	// 		"kay1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armcdn.EndpointProperties{
	// 		ContentTypesToCompress: []*string{
	// 		},
	// 		GeoFilters: []*armcdn.GeoFilter{
	// 		},
	// 		IsCompressionEnabled: to.Ptr(false),
	// 		IsHTTPAllowed: to.Ptr(true),
	// 		IsHTTPSAllowed: to.Ptr(true),
	// 		OptimizationType: to.Ptr(armcdn.OptimizationTypeDynamicSiteAcceleration),
	// 		OriginHostHeader: to.Ptr("www.bing.com"),
	// 		ProbePath: to.Ptr("/image"),
	// 		QueryStringCachingBehavior: to.Ptr(armcdn.QueryStringCachingBehaviorNotSet),
	// 		HostName: to.Ptr("endpoint1.azureedge.net"),
	// 		Origins: []*armcdn.DeepCreatedOrigin{
	// 			{
	// 				Name: to.Ptr("www-bing-com"),
	// 				Properties: &armcdn.DeepCreatedOriginProperties{
	// 					HostName: to.Ptr("www.bing.com"),
	// 					HTTPPort: to.Ptr[int32](80),
	// 					HTTPSPort: to.Ptr[int32](443),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armcdn.EndpointProvisioningStateSucceeded),
	// 		ResourceState: to.Ptr(armcdn.EndpointResourceStateStopping),
	// 	},
	// }
}
