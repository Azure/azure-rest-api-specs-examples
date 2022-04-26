Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.6.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/NfsV3AccountCreate.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.AccountCreateParameters{
			Kind:     to.Ptr(armstorage.KindBlockBlobStorage),
			Location: to.Ptr("<location>"),
			Properties: &armstorage.AccountPropertiesCreateParameters{
				IsHnsEnabled: to.Ptr(true),
				EnableNfsV3:  to.Ptr(true),
				NetworkRuleSet: &armstorage.NetworkRuleSet{
					Bypass:        to.Ptr(armstorage.BypassAzureServices),
					DefaultAction: to.Ptr(armstorage.DefaultActionAllow),
					IPRules:       []*armstorage.IPRule{},
					VirtualNetworkRules: []*armstorage.VirtualNetworkRule{
						{
							VirtualNetworkResourceID: to.Ptr("<virtual-network-resource-id>"),
						}},
				},
				EnableHTTPSTrafficOnly: to.Ptr(false),
			},
			SKU: &armstorage.SKU{
				Name: to.Ptr(armstorage.SKUNamePremiumLRS),
			},
		},
		&armstorage.AccountsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
