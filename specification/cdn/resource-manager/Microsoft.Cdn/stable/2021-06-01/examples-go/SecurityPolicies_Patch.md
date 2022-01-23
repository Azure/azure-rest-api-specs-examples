Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.3.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

```go
package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/SecurityPolicies_Patch.json
func ExampleSecurityPoliciesClient_BeginPatch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewSecurityPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPatch(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<security-policy-name>",
		armcdn.SecurityPolicyUpdateParameters{
			Properties: &armcdn.SecurityPolicyUpdateProperties{
				Parameters: &armcdn.SecurityPolicyWebApplicationFirewallParameters{
					Type: armcdn.SecurityPolicyType("WebApplicationFirewall").ToPtr(),
					Associations: []*armcdn.SecurityPolicyWebApplicationFirewallAssociation{
						{
							Domains: []*armcdn.ActivatedResourceReference{
								{
									ID: to.StringPtr("<id>"),
								},
								{
									ID: to.StringPtr("<id>"),
								}},
							PatternsToMatch: []*string{
								to.StringPtr("/*")},
						}},
					WafPolicy: &armcdn.ResourceReference{
						ID: to.StringPtr("<id>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SecurityPoliciesClientPatchResult)
}
```
