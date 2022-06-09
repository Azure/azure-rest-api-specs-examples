```go
package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Routes_Update.json
func ExampleRoutesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewRoutesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"RG",
		"profile1",
		"endpoint1",
		"route1",
		armcdn.RouteUpdateParameters{
			Properties: &armcdn.RouteUpdatePropertiesParameters{
				CacheConfiguration: &armcdn.AfdRouteCacheConfiguration{
					CompressionSettings: &armcdn.CompressionSettings{
						ContentTypesToCompress: []*string{
							to.Ptr("text/html"),
							to.Ptr("application/octet-stream")},
						IsCompressionEnabled: to.Ptr(true),
					},
					QueryStringCachingBehavior: to.Ptr(armcdn.AfdQueryStringCachingBehaviorIgnoreQueryString),
				},
				CustomDomains: []*armcdn.ActivatedResourceReference{
					{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1"),
					}},
				EnabledState:        to.Ptr(armcdn.EnabledStateEnabled),
				ForwardingProtocol:  to.Ptr(armcdn.ForwardingProtocolMatchRequest),
				HTTPSRedirect:       to.Ptr(armcdn.HTTPSRedirectEnabled),
				LinkToDefaultDomain: to.Ptr(armcdn.LinkToDefaultDomainEnabled),
				OriginGroup: &armcdn.ResourceReference{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"),
				},
				PatternsToMatch: []*string{
					to.Ptr("/*")},
				RuleSets: []*armcdn.ResourceReference{
					{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1"),
					}},
				SupportedProtocols: []*armcdn.AFDEndpointProtocols{
					to.Ptr(armcdn.AFDEndpointProtocolsHTTPS),
					to.Ptr(armcdn.AFDEndpointProtocolsHTTP)},
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv1.0.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.
