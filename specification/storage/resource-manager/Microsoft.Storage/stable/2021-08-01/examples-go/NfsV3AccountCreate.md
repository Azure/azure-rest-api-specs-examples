Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.3.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/NfsV3AccountCreate.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.AccountCreateParameters{
			Kind:     armstorage.Kind("BlockBlobStorage").ToPtr(),
			Location: to.StringPtr("<location>"),
			Properties: &armstorage.AccountPropertiesCreateParameters{
				IsHnsEnabled: to.BoolPtr(true),
				EnableNfsV3:  to.BoolPtr(true),
				NetworkRuleSet: &armstorage.NetworkRuleSet{
					Bypass:        armstorage.Bypass("AzureServices").ToPtr(),
					DefaultAction: armstorage.DefaultActionAllow.ToPtr(),
					IPRules:       []*armstorage.IPRule{},
					VirtualNetworkRules: []*armstorage.VirtualNetworkRule{
						{
							VirtualNetworkResourceID: to.StringPtr("<virtual-network-resource-id>"),
						}},
				},
				EnableHTTPSTrafficOnly: to.BoolPtr(false),
			},
			SKU: &armstorage.SKU{
				Name: armstorage.SKUName("Premium_LRS").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.AccountsClientCreateResult)
}
```
