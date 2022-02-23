Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicySignatureOverridesPatch.json
func ExampleFirewallPolicyIdpsSignaturesOverridesClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewFirewallPolicyIdpsSignaturesOverridesClient("<subscription-id>", cred, nil)
	res, err := client.Patch(ctx,
		"<resource-group-name>",
		"<firewall-policy-name>",
		armnetwork.SignaturesOverrides{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
			ID:   to.StringPtr("<id>"),
			Properties: &armnetwork.SignaturesOverridesProperties{
				Signatures: map[string]*string{
					"2000105": to.StringPtr("Off"),
					"2000106": to.StringPtr("Deny"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FirewallPolicyIdpsSignaturesOverridesClientPatchResult)
}
```
