package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Routes_Get.json
func ExampleRoutesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoutesClient().Get(ctx, "RG", "profile1", "endpoint1", "route1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Route = armcdn.Route{
	// 	Name: to.Ptr("route1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/afdendpoints/routes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/afdendpoints/endpoint1/routes/route1"),
	// 	Properties: &armcdn.RouteProperties{
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		CacheConfiguration: &armcdn.AfdRouteCacheConfiguration{
	// 			CompressionSettings: &armcdn.CompressionSettings{
	// 				ContentTypesToCompress: []*string{
	// 					to.Ptr("text/html"),
	// 					to.Ptr("application/octet-stream")},
	// 					IsCompressionEnabled: to.Ptr(true),
	// 				},
	// 				QueryStringCachingBehavior: to.Ptr(armcdn.AfdQueryStringCachingBehaviorIgnoreQueryString),
	// 			},
	// 			CustomDomains: []*armcdn.ActivatedResourceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1"),
	// 			}},
	// 			EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
	// 			ForwardingProtocol: to.Ptr(armcdn.ForwardingProtocolMatchRequest),
	// 			HTTPSRedirect: to.Ptr(armcdn.HTTPSRedirectEnabled),
	// 			LinkToDefaultDomain: to.Ptr(armcdn.LinkToDefaultDomainEnabled),
	// 			OriginGroup: &armcdn.ResourceReference{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"),
	// 			},
	// 			PatternsToMatch: []*string{
	// 				to.Ptr("/*")},
	// 				RuleSets: []*armcdn.ResourceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1"),
	// 				}},
	// 				SupportedProtocols: []*armcdn.AFDEndpointProtocols{
	// 					to.Ptr(armcdn.AFDEndpointProtocolsHTTPS),
	// 					to.Ptr(armcdn.AFDEndpointProtocolsHTTP)},
	// 				},
	// 			}
}
