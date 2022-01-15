Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Felastic%2Farmelastic%2Fv0.2.0/sdk/resourcemanager/elastic/armelastic/README.md) on how to add the SDK to your project and authenticate.

```go
package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2020-07-01-preview/examples/VMIngestion_Details.json
func ExampleVMIngestionClient_Details() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armelastic.NewVMIngestionClient("<subscription-id>", cred, nil)
	res, err := client.Details(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VMIngestionClientDetailsResult)
}
```
