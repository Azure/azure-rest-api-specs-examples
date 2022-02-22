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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyQuerySignatureOverrides.json
func ExampleFirewallPolicyIdpsSignaturesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewFirewallPolicyIdpsSignaturesClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<firewall-policy-name>",
		armnetwork.IDPSQueryObject{
			Filters: []*armnetwork.FilterItems{
				{
					Field: to.StringPtr("<field>"),
					Values: []*string{
						to.StringPtr("Deny")},
				}},
			OrderBy: &armnetwork.OrderBy{
				Field: to.StringPtr("<field>"),
				Order: armnetwork.OrderByOrder("Ascending").ToPtr(),
			},
			ResultsPerPage: to.Int32Ptr(20),
			Search:         to.StringPtr("<search>"),
			Skip:           to.Int32Ptr(0),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FirewallPolicyIdpsSignaturesClientListResult)
}
```
