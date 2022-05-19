Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaps%2Farmmaps%2Fv0.5.0/sdk/resourcemanager/maps/armmaps/README.md) on how to add the SDK to your project and authenticate.

```go
package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/AccountListSAS.json
func ExampleAccountsClient_ListSas() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmaps.NewAccountsClient("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListSas(ctx,
		"myResourceGroup",
		"myMapsAccount",
		armmaps.AccountSasParameters{
			Expiry:           to.Ptr("2017-05-24T11:42:03.1567373Z"),
			MaxRatePerSecond: to.Ptr[int32](500),
			PrincipalID:      to.Ptr("e917f87b-324d-4728-98ed-e31d311a7d65"),
			Regions: []*string{
				to.Ptr("eastus")},
			SigningKey: to.Ptr(armmaps.SigningKeyPrimaryKey),
			Start:      to.Ptr("2017-05-24T10:42:03.1567373Z"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
