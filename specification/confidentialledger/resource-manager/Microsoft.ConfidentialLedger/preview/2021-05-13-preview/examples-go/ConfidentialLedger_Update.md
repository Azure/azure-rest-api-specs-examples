Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fconfidentialledger%2Farmconfidentialledger%2Fv0.1.0/sdk/resourcemanager/confidentialledger/armconfidentialledger/README.md) on how to add the SDK to your project and authenticate.

```go
package armconfidentialledger_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger"
)

// x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2021-05-13-preview/examples/ConfidentialLedger_Update.json
func ExampleLedgerClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfidentialledger.NewLedgerClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<ledger-name>",
		armconfidentialledger.ConfidentialLedger{
			Location: armconfidentialledger.Location{
				Location: to.StringPtr("<location>"),
			},
			Tags: armconfidentialledger.Tags{
				Tags: map[string]*string{
					"additionProps2":   to.StringPtr("additional property value"),
					"additionalProps1": to.StringPtr("additional properties"),
				},
			},
			Properties: &armconfidentialledger.LedgerProperties{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfidentialLedger.ID: %s\n", *res.ID)
}
```
