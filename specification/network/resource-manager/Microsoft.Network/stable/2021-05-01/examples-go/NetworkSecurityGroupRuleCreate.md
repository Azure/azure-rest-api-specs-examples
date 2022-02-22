Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkSecurityGroupRuleCreate.json
func ExampleSecurityRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewSecurityRulesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<network-security-group-name>",
		"<security-rule-name>",
		armnetwork.SecurityRule{
			Properties: &armnetwork.SecurityRulePropertiesFormat{
				Access:                   armnetwork.SecurityRuleAccess("Deny").ToPtr(),
				DestinationAddressPrefix: to.StringPtr("<destination-address-prefix>"),
				DestinationPortRange:     to.StringPtr("<destination-port-range>"),
				Direction:                armnetwork.SecurityRuleDirection("Outbound").ToPtr(),
				Priority:                 to.Int32Ptr(100),
				SourceAddressPrefix:      to.StringPtr("<source-address-prefix>"),
				SourcePortRange:          to.StringPtr("<source-port-range>"),
				Protocol:                 armnetwork.SecurityRuleProtocol("*").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.SecurityRulesClientCreateOrUpdateResult)
}
```
