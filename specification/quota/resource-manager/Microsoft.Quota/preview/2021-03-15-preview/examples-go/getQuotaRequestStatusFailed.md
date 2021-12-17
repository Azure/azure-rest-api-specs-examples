Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fquota%2Farmquota%2Fv0.1.0/sdk/resourcemanager/quota/armquota/README.md) on how to add the SDK to your project and authenticate.

```go
package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/getQuotaRequestStatusFailed.json
func ExampleQuotaRequestStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armquota.NewQuotaRequestStatusClient(cred, nil)
	res, err := client.Get(ctx,
		"<id>",
		"<scope>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("QuotaRequestDetails.ID: %s\n", *res.ID)
}
```
