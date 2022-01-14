Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappconfiguration%2Farmappconfiguration%2Fv0.2.0/sdk/resourcemanager/appconfiguration/armappconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
)

// x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-03-01-preview/examples/ConfigurationStoresCreateKeyValue.json
func ExampleKeyValuesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappconfiguration.NewKeyValuesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<config-store-name>",
		"<key-value-name>",
		&armappconfiguration.KeyValuesClientCreateOrUpdateOptions{KeyValueParameters: &armappconfiguration.KeyValue{
			Properties: &armappconfiguration.KeyValueProperties{
				Tags: map[string]*string{
					"tag1": to.StringPtr("tagValue1"),
					"tag2": to.StringPtr("tagValue2"),
				},
				Value: to.StringPtr("<value>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KeyValuesClientCreateOrUpdateResult)
}
```
