Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredis%2Farmredis%2Fv0.3.1/sdk/resourcemanager/redis/armredis/README.md) on how to add the SDK to your project and authenticate.

```go
package armredis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
)

// x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheImport.json
func ExampleClient_BeginImportData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armredis.NewClient("<subscription-id>", cred, nil)
	poller, err := client.BeginImportData(ctx,
		"<resource-group-name>",
		"<name>",
		armredis.ImportRDBParameters{
			Format: to.StringPtr("<format>"),
			Files: []*string{
				to.StringPtr("http://fileuris.contoso.com/pathtofile1")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
