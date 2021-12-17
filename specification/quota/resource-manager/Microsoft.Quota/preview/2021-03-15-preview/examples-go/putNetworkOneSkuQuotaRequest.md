Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fquota%2Farmquota%2Fv0.1.0/sdk/resourcemanager/quota/armquota/README.md) on how to add the SDK to your project and authenticate.

```go
package armquota_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/putNetworkOneSkuQuotaRequest.json
func ExampleQuotaClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armquota.NewQuotaClient(cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-name>",
		"<scope>",
		armquota.CurrentQuotaLimitBase{
			Properties: &armquota.QuotaProperties{
				Name: &armquota.ResourceName{
					Value: to.StringPtr("<value>"),
				},
				Limit: &armquota.LimitObject{
					LimitJSONObject: armquota.LimitJSONObject{
						LimitObjectType: armquota.LimitTypeLimitValue.ToPtr(),
					},
					Value: to.Int32Ptr(10),
				},
				ResourceType: to.StringPtr("<resource-type>"),
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
	log.Printf("CurrentQuotaLimitBase.ID: %s\n", *res.ID)
}
```
